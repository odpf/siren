package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/internal/store/model"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/pgc"
	"github.com/goto/siren/pkg/structure"
	"github.com/lib/pq"
	"go.nhat.io/otelsql"
	"go.opentelemetry.io/otel/attribute"
)

const subscriptionInsertQuery = `
INSERT INTO subscriptions (namespace_id, urn, receiver, match, metadata, created_by, updated_by, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, now(), now())
RETURNING *
`

const subscriptionUpdateQuery = `
UPDATE subscriptions SET namespace_id=$2, urn=$3, receiver=$4, match=$5, metadata=$6, updated_by=$7, updated_at=now()
WHERE id = $1
RETURNING *
`

const subscriptionDeleteQuery = `
DELETE from subscriptions where id=$1
`

var subscriptionListQueryBuilder = sq.Select(
	"s.id as id",
	"s.namespace_id as namespace_id",
	"s.urn as urn",
	"s.receiver as receiver",
	"s.match as match",
	"s.metadata as metadata",
	"s.created_by as created_by",
	"s.updated_by as updated_by",
	"s.created_at as created_at",
	"s.updated_at as updated_at",
).Distinct().From("subscriptions s")

var subscriptionMatchLabelsFetchReceiversQueryBuilder = sq.Select(
	"r.id as id",
	"r.name as name",
	"r.type as type",
	"r.labels as labels",
	"r.parent_id as parent_id",
	"r.created_at as created_at",
	"r.updated_at as updated_at",
	"s.id as subscription_id",
	"s.match as match",
).Column(
	sq.Expr("r.configurations || COALESCE(pr.configurations, '{}'::jsonb) AS configurations"),
).
	From("subscriptions s").
	InnerJoin("subscriptions_receivers sr ON s.id = sr.subscription_id").
	InnerJoin("receivers r ON r.id = sr.receiver_id").
	LeftJoin("receivers pr ON pr.id = r.parent_id")

// SubscriptionRepository talks to the store to read or insert data
type SubscriptionRepository struct {
	client *pgc.Client
}

// NewSubscriptionRepository returns SubscriptionRepository struct
func NewSubscriptionRepository(client *pgc.Client) *SubscriptionRepository {
	return &SubscriptionRepository{
		client: client,
	}
}

func (r *SubscriptionRepository) List(ctx context.Context, flt subscription.Filter) ([]subscription.Subscription, error) {
	var queryBuilder = subscriptionListQueryBuilder

	if len(flt.IDs) != 0 {
		queryBuilder = queryBuilder.Where("s.id = any(?)", pq.Array(flt.IDs))
	}

	if flt.NamespaceID != 0 {
		queryBuilder = queryBuilder.Where("s.namespace_id = ?", flt.NamespaceID)
	}

	// given map of metadata from input [mf], look for rows that [mf] exist in metadata column in DB
	if len(flt.Metadata) != 0 {
		metadataJSON, err := json.Marshal(flt.Metadata)
		if err != nil {
			return nil, errors.ErrInvalid.WithMsgf("problem marshalling subscription metadata json to string with err: %s", err.Error())
		}
		conditionedJSON := structure.ConditionJSONString(json.RawMessage(metadataJSON))
		queryBuilder = queryBuilder.Where(fmt.Sprintf("s.metadata @> '%s'::jsonb", conditionedJSON))
	}

	// given map of string from input [mf], look for rows that [mf] exist in match column in DB
	if len(flt.Match) != 0 {
		labelsJSON, err := json.Marshal(flt.Match)
		if err != nil {
			return nil, errors.ErrInvalid.WithMsgf("problem marshalling match json to string with err: %s", err.Error())
		}
		conditionedJSON := structure.ConditionJSONString(json.RawMessage(labelsJSON))
		queryBuilder = queryBuilder.Where(fmt.Sprintf("s.match @> '%s'::jsonb", conditionedJSON))
	}

	// given map of string from input [mf], look for rows that has match column in DB that are subset of [mf]
	if len(flt.NotificationMatch) != 0 {
		labelsJSON, err := json.Marshal(flt.NotificationMatch)
		if err != nil {
			return nil, errors.ErrInvalid.WithMsgf("problem marshalling notification labels json to string with err: %s", err.Error())
		}
		conditionedJSON := structure.ConditionJSONString(json.RawMessage(labelsJSON))
		queryBuilder = queryBuilder.Where(fmt.Sprintf("s.match <@ '%s'::jsonb", conditionedJSON))
	}

	// should lookup subscription_receiver table if we want to query based on receiver
	if flt.ReceiverID != 0 || len(flt.SubscriptionReceiverLabels) != 0 {
		queryBuilder = queryBuilder.InnerJoin("subscriptions_receivers sr ON s.id = sr.subscription_id")
		if flt.ReceiverID != 0 {
			queryBuilder = queryBuilder.Where("sr.receiver_id = ?", flt.ReceiverID)
		}
		if len(flt.SubscriptionReceiverLabels) != 0 {
			subrReceiverLabelsJSON, err := json.Marshal(flt.SubscriptionReceiverLabels)
			if err != nil {
				return nil, errors.ErrInvalid.WithMsgf("problem marshalling subscription_receiver_labels json to string with err: %s", err.Error())
			}
			conditionedJSON := structure.ConditionJSONString(json.RawMessage(subrReceiverLabelsJSON))
			queryBuilder = queryBuilder.Where(fmt.Sprintf("sr.labels @> '%s'::jsonb", conditionedJSON))
		}
	}

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "List"),
			attribute.String("db.sql.table", "subscriptions"),
		}...,
	)

	rows, err := r.client.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptionsDomain []subscription.Subscription
	for rows.Next() {
		var subscriptionModel model.Subscription
		if err := rows.StructScan(&subscriptionModel); err != nil {
			return nil, err
		}

		subscriptionsDomain = append(subscriptionsDomain, *subscriptionModel.ToDomain())
	}

	return subscriptionsDomain, nil
}

func (r *SubscriptionRepository) Create(ctx context.Context, sub *subscription.Subscription) error {
	if sub == nil {
		return errors.New("subscription domain is nil")
	}

	subscriptionModel := new(model.Subscription)
	subscriptionModel.FromDomain(*sub)

	var newSubscriptionModel model.Subscription

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "Create"),
			attribute.String("db.sql.table", "subscriptions"),
		}...,
	)

	if err := r.client.QueryRowxContext(ctx, subscriptionInsertQuery,
		subscriptionModel.NamespaceID,
		subscriptionModel.URN,
		subscriptionModel.Receiver,
		subscriptionModel.Match,
		subscriptionModel.Metadata,
		subscriptionModel.CreatedBy,
		subscriptionModel.UpdatedBy,
	).StructScan(&newSubscriptionModel); err != nil {
		err = pgc.CheckError(err)
		if errors.Is(err, pgc.ErrDuplicateKey) {
			return subscription.ErrDuplicate
		}
		if errors.Is(err, pgc.ErrForeignKeyViolation) {
			return subscription.ErrRelation
		}
		return err
	}

	*sub = *newSubscriptionModel.ToDomain()

	return nil
}

func (r *SubscriptionRepository) Get(ctx context.Context, id uint64) (*subscription.Subscription, error) {
	query, args, err := subscriptionListQueryBuilder.Where("id = ?", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var subscriptionModel model.Subscription

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "Get"),
			attribute.String("db.sql.table", "subscriptions"),
		}...,
	)

	if err := r.client.QueryRowxContext(ctx, query, args...).StructScan(&subscriptionModel); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, subscription.NotFoundError{ID: id}
		}
		return nil, err
	}

	return subscriptionModel.ToDomain(), nil
}

func (r *SubscriptionRepository) Update(ctx context.Context, sub *subscription.Subscription) error {
	if sub == nil {
		return errors.New("subscription domain is nil")
	}

	subscriptionModel := new(model.Subscription)
	subscriptionModel.FromDomain(*sub)

	var newSubscriptionModel model.Subscription

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "Update"),
			attribute.String("db.sql.table", "subscriptions"),
		}...,
	)

	if err := r.client.QueryRowxContext(ctx, subscriptionUpdateQuery,
		subscriptionModel.ID,
		subscriptionModel.NamespaceID,
		subscriptionModel.URN,
		subscriptionModel.Receiver,
		subscriptionModel.Match,
		subscriptionModel.Metadata,
		subscriptionModel.UpdatedBy,
	).StructScan(&newSubscriptionModel); err != nil {
		err = pgc.CheckError(err)
		if errors.Is(err, sql.ErrNoRows) {
			return subscription.NotFoundError{ID: subscriptionModel.ID}
		}
		if errors.Is(err, pgc.ErrDuplicateKey) {
			return subscription.ErrDuplicate
		}
		if errors.Is(err, pgc.ErrForeignKeyViolation) {
			return subscription.ErrRelation
		}
		return err
	}

	*sub = *newSubscriptionModel.ToDomain()

	return nil
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id uint64) error {

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "Delete"),
			attribute.String("db.sql.table", "subscriptions"),
		}...,
	)

	if _, err := r.client.ExecContext(ctx, subscriptionDeleteQuery, id); err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepository) MatchLabelsFetchReceivers(ctx context.Context, flt subscription.Filter) ([]subscription.ReceiverView, error) {
	if len(flt.Match) == 0 {
		return nil, errors.ErrInvalid.WithMsgf("matcher cannot be empty")
	}

	var queryBuilder = subscriptionMatchLabelsFetchReceiversQueryBuilder.Where("sr.deleted_at IS NULL")
	// given map of string from input [mf], look for rows that [mf] exist in match column in DB
	matchJSON, err := json.Marshal(flt.Match)
	if err != nil {
		return nil, errors.ErrInvalid.WithMsgf("problem marshalling notification labels json to string with err: %s", err.Error())
	}
	queryBuilder = queryBuilder.Where(fmt.Sprintf("s.match <@ '%s'::jsonb", string(json.RawMessage(matchJSON))))

	if flt.NamespaceID != 0 {
		queryBuilder = queryBuilder.Where(fmt.Sprintf("s.namespace_id = %d", flt.NamespaceID))
	}

	query, _, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "MatchLabelsFetchReceivers"),
			attribute.String("db.sql.table", "subscriptions"),
		}...,
	)

	rows, err := r.client.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var receivers []subscription.ReceiverView
	for rows.Next() {
		var receiverModel model.ReceiverView
		if err := rows.StructScan(&receiverModel); err != nil {
			return nil, err
		}

		receivers = append(receivers, *receiverModel.ToDomain())
	}

	return receivers, nil
}

func (r *SubscriptionRepository) WithTransaction(ctx context.Context) context.Context {
	return r.client.WithTransaction(ctx, nil)
}

func (r *SubscriptionRepository) Rollback(ctx context.Context, err error) error {
	if txErr := r.client.Rollback(ctx); txErr != nil {
		return fmt.Errorf("rollback error %s with error: %w", txErr.Error(), err)
	}
	return nil
}

func (r *SubscriptionRepository) Commit(ctx context.Context) error {
	return r.client.Commit(ctx)
}
