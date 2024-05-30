package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/internal/store/model"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/pgc"
	"go.nhat.io/otelsql"
	"go.opentelemetry.io/otel/attribute"
)

const notificationInsertQuery = `
INSERT INTO notifications (namespace_id, type, data, labels, valid_duration, template, unique_key, receiver_selectors, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now())
RETURNING *
`

const notificationInsertNamedQuery = `
INSERT INTO notifications
	(namespace_id, type, data, labels, valid_duration, template, unique_key, receiver_selectors, created_at)
    VALUES (:namespace_id, :type, :data, :labels, :valid_duration, :template, :unique_key, :receiver_selectors, now())
RETURNING *
`

var notificationListQueryBuilder = sq.Select(
	"id",
	"namespace_id",
	"data",
	"type",
	"labels",
	"valid_duration",
	"template",
	"created_at",
	"unique_key",
	"receiver_selectors",
).From("notifications")

// NotificationRepository talks to the store to read or insert data
type NotificationRepository struct {
	client *pgc.Client
}

// NewNotificationRepository returns NotificationRepository struct
func NewNotificationRepository(client *pgc.Client) *NotificationRepository {
	return &NotificationRepository{
		client: client,
	}
}

func (r *NotificationRepository) Create(ctx context.Context, n notification.Notification) (notification.Notification, error) {
	nModel := new(model.Notification)
	nModel.FromDomain(n)

	var newNModel model.Notification

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "Create"),
			attribute.String("db.sql.table", "notifications"),
		}...,
	)

	if err := r.client.QueryRowxContext(ctx, notificationInsertQuery,
		nModel.NamespaceID,
		nModel.Type,
		nModel.Data,
		nModel.Labels,
		nModel.ValidDuration,
		nModel.Template,
		nModel.UniqueKey,
		nModel.ReceiverSelectors,
	).StructScan(&newNModel); err != nil {
		return notification.Notification{}, err
	}

	return *newNModel.ToDomain(), nil
}

func (r *NotificationRepository) BulkCreate(ctx context.Context, ns []notification.Notification) ([]notification.Notification, error) {
	var notificationsModel = []model.Notification{}

	for _, n := range ns {
		nModel := new(model.Notification)
		nModel.FromDomain(n)
		notificationsModel = append(notificationsModel, *nModel)
	}

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "BulkCreate"),
			attribute.String("db.sql.table", "notifications"),
		}...,
	)

	rows, err := r.client.NamedQueryContext(ctx, notificationInsertNamedQuery, notificationsModel)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notificationsDomain := []notification.Notification{}
	for rows.Next() {
		var notificationModel model.Notification
		if err := rows.StructScan(&notificationModel); err != nil {
			return nil, err
		}
		notificationsDomain = append(notificationsDomain, *notificationModel.ToDomain())
	}
	return notificationsDomain, nil
}

func (r *NotificationRepository) List(ctx context.Context, flt notification.Filter) ([]notification.Notification, error) {
	var queryBuilder = notificationListQueryBuilder

	if flt.Type != "" {
		queryBuilder = queryBuilder.Where("type = ?", flt.Type)
	}

	if flt.Template != "" {
		queryBuilder = queryBuilder.Where("template = ?", flt.Template)
	}

	if flt.Labels != nil {
		labelsJSON, err := json.Marshal(flt.Labels)
		if err != nil {
			return nil, errors.ErrInvalid.WithMsgf("problem marshalling label %v json to string with err: %s", flt.Labels, err.Error())
		}

		matchLabelsExpression := sq.Expr(fmt.Sprintf("labels @> '%s'::jsonb", string(json.RawMessage(labelsJSON))))

		queryBuilder = queryBuilder.Where(matchLabelsExpression)
	}

	if flt.ReceiverSelector != nil {
		rs, err := json.Marshal(flt.ReceiverSelector)
		if err != nil {
			return nil, err
		}

		recieverSelectors := fmt.Sprintf("[" + string(rs) + "]")
		matchReceiverSelectorExpression := sq.Expr("receiver_selectors @> ?", recieverSelectors)
		queryBuilder = queryBuilder.Where(matchReceiverSelectorExpression)
	}

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "List"),
			attribute.String("db.sql.table", "notitfications"),
		}...,
	)

	rows, err := r.client.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notificationsDomain := []notification.Notification{}
	for rows.Next() {
		var notificationModel model.Notification
		if err := rows.StructScan(&notificationModel); err != nil {
			return nil, err
		}
		notificationsDomain = append(notificationsDomain, *notificationModel.ToDomain())
	}
	return notificationsDomain, nil
}

func (r *NotificationRepository) WithTransaction(ctx context.Context) context.Context {
	return r.client.WithTransaction(ctx, nil)
}

func (r *NotificationRepository) Rollback(ctx context.Context, err error) error {
	if txErr := r.client.Rollback(ctx); txErr != nil {
		return fmt.Errorf("rollback error %s with error: %w", txErr.Error(), err)
	}
	return nil
}

func (r *NotificationRepository) Commit(ctx context.Context) error {
	return r.client.Commit(ctx)
}
