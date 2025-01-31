package events

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/openshift/assisted-service/internal/common"
	commonevents "github.com/openshift/assisted-service/internal/common/events"
	eventsapi "github.com/openshift/assisted-service/internal/events/api"
	"github.com/openshift/assisted-service/internal/stream"
	"github.com/openshift/assisted-service/models"
	"github.com/openshift/assisted-service/pkg/auth"
	logutil "github.com/openshift/assisted-service/pkg/log"
	"github.com/openshift/assisted-service/pkg/requestid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DefaultEventCategories = []string{
	models.EventCategoryUser,
}

var defaultEventLimit int64 = 5000

type Events struct {
	db     *gorm.DB
	log    logrus.FieldLogger
	authz  auth.Authorizer
	stream stream.Notifier
}

func New(db *gorm.DB, authz auth.Authorizer, stream stream.Notifier, log logrus.FieldLogger) eventsapi.Handler {
	return &Events{
		db:     db,
		log:    log,
		authz:  authz,
		stream: stream,
	}
}

func (e *Events) v2SaveEvent(ctx context.Context, clusterID *strfmt.UUID, hostID *strfmt.UUID, infraEnvID *strfmt.UUID, name string, category string, severity string, message string, t time.Time, requestID string, props ...interface{}) {
	log := logutil.FromContext(ctx, e.log)
	tt := strfmt.DateTime(t)
	rid := strfmt.UUID(requestID)
	errMsg := make([]string, 0)
	additionalProps, err := toProps(props...)
	if err != nil {
		log.WithError(err).Error("failed to parse event's properties field")
	}
	event := common.Event{
		Event: models.Event{
			EventTime: &tt,
			Name:      name,
			Severity:  &severity,
			Category:  category,
			Message:   &message,
			RequestID: rid,
			Props:     additionalProps,
		},
	}
	if clusterID != nil {
		event.ClusterID = clusterID
		errMsg = append(errMsg, fmt.Sprintf("cluster_id = %s", clusterID.String()))
	}

	if hostID != nil {
		event.HostID = hostID
		errMsg = append(errMsg, fmt.Sprintf("host_id = %s", hostID.String()))
	}

	if infraEnvID != nil {
		event.InfraEnvID = infraEnvID
		errMsg = append(errMsg, fmt.Sprintf("infra_env_id = %s", infraEnvID.String()))
	}

	//each event is saved in its own embedded transaction
	var dberr error
	tx := e.db.Begin()
	defer func() {
		if dberr != nil {
			log.WithError(err).Errorf("failed to add event. Rolling back transaction on event=%s resources: %s",
				message, strings.Join(errMsg, " "))
			tx.Rollback()
		} else {
			tx.Commit()
			err = e.stream.Notify(ctx, &event)
			if err != nil {
				log.WithError(err).Warning("failed to notify event")
			}
		}
	}()

	// Check and if the event exceeds the limits:
	limitExceeded, limitReason, dberr := e.exceedsLimits(ctx, tx, &event)
	if dberr != nil {
		return
	}
	if limitExceeded {
		e.reportDiscarded(ctx, &event, limitReason)
		return
	}

	// Create the new event:
	dberr = tx.Create(&event).Error
}

// exceedsLimit checks if there are already events that are too close to the given one. It returns
// a boolean flag with the result and a set of log fields that explain why the limit was exceeded.
func (e *Events) exceedsLimits(ctx context.Context, tx *gorm.DB, event *common.Event) (result bool,
	reason logrus.Fields, err error) {
	// Do nothing if there is no configured limit:
	limit, ok := eventLimits[event.Name]
	if !ok {
		return
	}

	// Prepare the query to find the events whose distance to this one is less than the limit:
	query := tx.Table("events").
		Select("count(*)").
		Where("name = ?", event.Name).
		Where("event_time > ?", time.Now().Add(-limit))
	if event.ClusterID != nil {
		query = query.Where("cluster_id = ?", event.ClusterID.String())
	}
	if event.HostID != nil {
		query = query.Where("host_id = ?", event.HostID.String())
	}
	if event.InfraEnvID != nil {
		query = query.Where("infra_env_id = ?", event.InfraEnvID.String())
	}

	// Run the query:
	var count int
	err = query.Scan(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		result = true
		reason = logrus.Fields{
			"limit": limit,
			"count": count,
		}
	}
	return
}

// reportDiscarded writes to the log a message indicating that the given event has been discarded.
// The log message will include the details of the event and the reason.
func (e *Events) reportDiscarded(ctx context.Context, event *common.Event,
	reason logrus.Fields) {
	log := logutil.FromContext(ctx, e.log)
	fields := logrus.Fields{
		"name":       event.Name,
		"category":   event.Category,
		"request_id": event.RequestID.String(),
		"props":      event.Props,
	}
	if event.EventTime != nil {
		fields["time"] = event.EventTime.String()
	}
	if event.ClusterID != nil {
		fields["cluster_id"] = event.ClusterID.String()
	}
	if event.HostID != nil {
		fields["host_id"] = event.HostID.String()
	}
	if event.InfraEnvID != nil {
		fields["infra_env_id"] = event.InfraEnvID.String()
	}
	if event.Severity != nil {
		fields["severity"] = *event.Severity
	}
	if event.Message != nil {
		fields["message"] = *event.Message
	}
	for name, value := range reason {
		fields[name] = value
	}
	log.WithFields(fields).Warn("Event will be discarded")
}

// eventLimits contains the minimum distance in time between events. The key of the map is the
// event name and the value is the distance.
var eventLimits = map[string]time.Duration{
	commonevents.UpgradeAgentFailedEventName:   time.Hour,
	commonevents.UpgradeAgentFinishedEventName: time.Hour,
	commonevents.UpgradeAgentStartedEventName:  time.Hour,
}

func (e *Events) SendClusterEvent(ctx context.Context, event eventsapi.ClusterEvent) {
	e.SendClusterEventAtTime(ctx, event, time.Now())
}

func (e *Events) SendClusterEventAtTime(ctx context.Context, event eventsapi.ClusterEvent, eventTime time.Time) {
	cID := event.GetClusterId()
	e.V2AddEvent(ctx, &cID, nil, nil, event.GetName(), event.GetSeverity(), event.FormatMessage(), eventTime)
}

func (e *Events) SendHostEvent(ctx context.Context, event eventsapi.HostEvent) {
	e.SendHostEventAtTime(ctx, event, time.Now())
}

func (e *Events) SendHostEventAtTime(ctx context.Context, event eventsapi.HostEvent, eventTime time.Time) {
	hostID := event.GetHostId()
	infraEnvID := event.GetInfraEnvId()
	e.V2AddEvent(ctx, event.GetClusterId(), &hostID, &infraEnvID, event.GetName(), event.GetSeverity(), event.FormatMessage(), eventTime)
}

func (e *Events) SendInfraEnvEvent(ctx context.Context, event eventsapi.InfraEnvEvent) {
	e.SendInfraEnvEventAtTime(ctx, event, time.Now())
}

func (e *Events) SendInfraEnvEventAtTime(ctx context.Context, event eventsapi.InfraEnvEvent, eventTime time.Time) {
	infraEnvID := event.GetInfraEnvId()
	e.V2AddEvent(ctx, event.GetClusterId(), nil, &infraEnvID, event.GetName(), event.GetSeverity(), event.FormatMessage(), eventTime)
}

func (e *Events) V2AddEvent(ctx context.Context, clusterID *strfmt.UUID, hostID *strfmt.UUID, infraEnvID *strfmt.UUID, name string, severity string, msg string, eventTime time.Time, props ...interface{}) {
	requestID := requestid.FromContext(ctx)
	e.v2SaveEvent(ctx, clusterID, hostID, infraEnvID, name, models.EventCategoryUser, severity, msg, eventTime, requestID, props...)
}

func (e *Events) NotifyInternalEvent(ctx context.Context, clusterID *strfmt.UUID, hostID *strfmt.UUID, infraEnvID *strfmt.UUID, msg string) {
	log := logutil.FromContext(ctx, e.log)
	log.Debugf("Notifying internal event %s, nothing to do", msg)
}

func (e *Events) V2AddMetricsEvent(ctx context.Context, clusterID *strfmt.UUID, hostID *strfmt.UUID, infraEnvID *strfmt.UUID, name string, severity string, msg string, eventTime time.Time, props ...interface{}) {
	requestID := requestid.FromContext(ctx)
	e.v2SaveEvent(ctx, clusterID, hostID, infraEnvID, name, models.EventCategoryMetrics, severity, msg, eventTime, requestID, props...)
}

func filterEvents(db *gorm.DB, clusterID *strfmt.UUID, hostIds []strfmt.UUID, infraEnvID *strfmt.UUID, severity []string, message *string, deletedHosts, clusterLevel *bool) *gorm.DB {
	if clusterID != nil {
		db = db.Where("events.cluster_id = ?", clusterID.String())

		// filter by event severity
		if severity != nil {
			db = db.Where("events.severity IN (?)", severity)
		}

		// filter by event message
		if message != nil {
			db = db.Where("events.message LIKE ?", fmt.Sprintf("%%%s%%", *message))
		}

		// cluster level and specific hosts events
		if swag.BoolValue(clusterLevel) && !swag.BoolValue(deletedHosts) {
			db = db.Where("events.host_id IS NULL")
		}

		// deleted hosts and specific hosts events
		if !swag.BoolValue(clusterLevel) && swag.BoolValue(deletedHosts) {
			db = db.Where("hosts.deleted_at IS NOT NULL")
		}

		// deleted hosts, cluster level and specific hosts events
		if swag.BoolValue(clusterLevel) && swag.BoolValue(deletedHosts) {
			db = db.Where("hosts.deleted_at IS NOT NULL").
				Or("events.host_id IS NULL")
		}

		// In case none of the latter occurred and we are going to hit the next condition,
		// the query should be prefixed with 'Where' before 'Or'
		if !swag.BoolValue(clusterLevel) && !swag.BoolValue(deletedHosts) && hostIds != nil {
			db = db.Where("FALSE")
		}

		if hostIds != nil {
			db = db.Or("events.host_id IN (?)", hostsUUIDsToStrings(hostIds))
		}

		return db
	}

	if hostIds != nil {
		return db.Where("events.host_id IN (?)", hostsUUIDsToStrings(hostIds))
	}

	if infraEnvID != nil {
		return db.Where("events.infra_env_id = ?", infraEnvID.String())
	}

	return db
}

func hostsUUIDsToStrings(hostIDs []strfmt.UUID) []string {
	result := []string{}
	for _, hostID := range hostIDs {
		result = append(result, hostID.String())
	}
	return result
}

func countEventsBySeverity(db *gorm.DB, clusterID *strfmt.UUID) (*common.EventSeverityCount, error) {
	var (
		total    int
		severity string
		rows     *sql.Rows
		err      error
	)

	if clusterID == nil {
		return &common.EventSeverityCount{}, nil
	}

	rows, err = db.Table("events").Where("events.cluster_id = ?", clusterID.String()).Select("COUNT(events.severity), events.severity").Group("events.severity").Rows()
	if err != nil {
		return nil, err
	}

	eventSeverityCount := common.EventSeverityCount{}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&total, &severity)
		if err != nil {
			return nil, err
		}
		eventSeverityCount[severity] = int64(total)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &eventSeverityCount, nil
}

func (e Events) prepareEventsTable(ctx context.Context, tx *gorm.DB, clusterID *strfmt.UUID, hostIds []strfmt.UUID, infraEnvID *strfmt.UUID, severity []string, message *string, deletedHosts *bool) *gorm.DB {
	allEvents := func() bool {
		return clusterID == nil && infraEnvID == nil && hostIds == nil
	}

	clusterBoundEvents := func() bool {
		return clusterID != nil
	}

	nonBoundEvents := func() bool {
		return clusterID == nil && infraEnvID != nil
	}

	hostOnlyEvents := func() bool {
		return clusterID == nil && infraEnvID == nil && hostIds != nil
	}

	//retrieveing all events can be done only by admins. This is done to restrict data
	//intensive queries by common users
	if allEvents() && e.authz.IsAdmin(ctx) {
		return tx
	}

	//for bound events that are searched with cluster id (whether on clusters, bound infra-env ,
	//host bound to a cluster or registered to a bound infra-env) check the access permission
	//relative to the cluster ownership
	if clusterBoundEvents() {
		tx = tx.Model(&common.Cluster{}).
			Select("events.*, clusters.user_name, clusters.org_id").
			Joins("INNER JOIN events ON clusters.id = events.cluster_id")

		// if deleted hosts flag is true, we need to add 'deleted_at' to know whether events are related to a deleted host
		if swag.BoolValue(deletedHosts) {
			tx = tx.Select("events.*, clusters.user_name, clusters.org_id, hosts.deleted_at").
				Joins("LEFT JOIN hosts ON events.host_id = hosts.id")
		}
		return tx
	}

	//for unbound events that are searched with infra-env id (whether events on hosts or the
	//infra-env level itself) check the access permission relative to the infra-env ownership
	if nonBoundEvents() {
		return tx.Model(&common.InfraEnv{}).
			Select("events.*, infra_envs.user_name, infra_envs.org_id").
			Joins("INNER JOIN events ON infra_envs.id = events.infra_env_id")
	}

	//for query made on the host only check the permission relative to it's infra-env. since
	//host table does not contain an org_id we can not perform a join on that table and has to go
	//through the infra-env table which is good because authorization is done on the infra-level
	if hostOnlyEvents() {
		return tx.Model(&common.Host{}).Select("events.*, infra_envs.user_name, infra_envs.org_id").
			Joins("INNER JOIN infra_envs ON hosts.infra_env_id = infra_envs.id").Joins("INNER JOIN events ON events.host_id = hosts.id")
	}

	//non supported option
	return nil
}

func preparePaginationParams(limit, offset *int64) (*int64, *int64) {
	if limit == nil {
		// If limit is not provided, we set it a default (currently 5000).
		limit = swag.Int64(defaultEventLimit)
	} else if *limit < -1 {
		// If limit not valid (smaller than -1), we set it -1 (no limit).
		limit = common.UnlimitedEvents
	}

	// if offset not specified or is negative, we return the first page.
	if offset == nil || *offset < 0 {
		offset = common.NoOffsetEvents
	}

	return limit, offset
}

func (e Events) queryEvents(ctx context.Context, selectedCategories []string, clusterID *strfmt.UUID, hostIds []strfmt.UUID, infraEnvID *strfmt.UUID, limit, offset *int64, severity []string, message *string, deletedHosts, clusterLevel *bool) ([]*common.Event, *common.EventSeverityCount, error) {

	tx := e.db.Where("category IN (?)", selectedCategories)

	eventSeverityCount, err := countEventsBySeverity(tx.Session(&gorm.Session{}), clusterID)
	if err != nil {
		return nil, nil, err
	}

	events := []*common.Event{}

	tx = tx.Order("event_time")

	// add authorization check to query
	if e.authz != nil {
		tx = e.authz.OwnedBy(ctx, tx)
	}

	tx = e.prepareEventsTable(ctx, tx, clusterID, hostIds, infraEnvID, severity, message, deletedHosts)
	if tx == nil {
		return make([]*common.Event, 0), &common.EventSeverityCount{}, nil
	}

	tx = filterEvents(tx, clusterID, hostIds, infraEnvID, severity, message, deletedHosts, clusterLevel)

	limit, offset = preparePaginationParams(limit, offset)
	if *limit == 0 {
		return make([]*common.Event, 0), eventSeverityCount, nil
	}

	err = tx.Offset(int(*offset)).Limit(int(*limit)).Find(&events).Error
	if err != nil {
		return nil, nil, err
	}

	return events, eventSeverityCount, nil
}

func (e Events) V2GetEvents(ctx context.Context, params *common.V2GetEventsParams) (*common.V2GetEventsResponse, error) {
	//initialize the selectedCategories either from the filter, if exists, or from the default values
	selectedCategories := make([]string, 0)
	if len(params.Categories) > 0 {
		selectedCategories = params.Categories[:]
	} else {
		selectedCategories = append(selectedCategories, DefaultEventCategories...)
	}

	events, eventSeverityCount, err := e.queryEvents(ctx, selectedCategories, params.ClusterID, params.HostIds, params.InfraEnvID, params.Limit, params.Offset, params.Severities, params.Message, params.DeletedHosts, params.ClusterLevel)
	if err != nil {
		return nil, err
	}
	return &common.V2GetEventsResponse{
		Events:             events,
		EventSeverityCount: eventSeverityCount,
	}, nil
}

func toProps(attrs ...interface{}) (result string, err error) {
	props := make(map[string]interface{})
	length := len(attrs)

	if length == 1 {
		if attr, ok := attrs[0].(map[string]interface{}); ok {
			props = attr
		}
	}

	if length > 1 && length%2 == 0 {
		for i := 0; i < length; i += 2 {
			props[attrs[i].(string)] = attrs[i+1]
		}
	}

	if len(props) > 0 {
		var b []byte
		if b, err = json.Marshal(props); err == nil {
			return string(b), nil
		}
	}

	return "", err
}
