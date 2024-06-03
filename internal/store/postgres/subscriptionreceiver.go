package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/core/subscriptionreceiver"
	"github.com/goto/siren/internal/store/model"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/pgc"
	"github.com/goto/siren/pkg/structure"
	"github.com/lib/pq"
	"go.nhat.io/otelsql"
	"go.opentelemetry.io/otel/attribute"
)

const subscriptionReceiverBulkInsertQuery = `
INSERT INTO subscriptions_receivers (subscription_id, receiver_id, labels, created_at, updated_at)
VALUES (:subscription_id, :receiver_id, :labels, now(), now())
RETURNING *
`

const subscriptionReceiverBulkUpsertQuery = `
INSERT INTO subscriptions_receivers 
	(subscription_id, receiver_id, labels, created_at, updated_at)
    VALUES (:subscription_id, :receiver_id, :labels, now(), now())
ON CONFLICT
	(subscription_id, receiver_id)
DO UPDATE SET
	labels = EXCLUDED.labels, updated_at = now(), deleted_at = NULL
RETURNING *
`

const subscriptionReceiverUpdateQuery = `
UPDATE subscriptions_receivers SET labels=$3, updated_at=now()
WHERE subscription_id = $1 AND receiver_id = $2 AND deleted_at IS NULL
RETURNING *
`

var subscriptionReceiverBulkSoftDeleteQueryBuilder = sq.Update("subscriptions_receivers").
	Set("deleted_at", "now()")

var subscriptionReceiverListQueryBuilder = sq.Select(
	"id",
	"subscription_id",
	"receiver_id",
	"labels",
	"created_at",
	"updated_at",
	"deleted_at",
).From("subscriptions_receivers")

// SubscriptionReceiverRepository talks to the store to read or insert data
type SubscriptionReceiverRepository struct {
	client *pgc.Client
}

// NewSSubscriptionReceiverRepository returns SubscriptionReceiverRepository struct
func NewSubscriptionReceiverRepository(client *pgc.Client) *SubscriptionReceiverRepository {
	return &SubscriptionReceiverRepository{
		client: client,
	}
}

func (s *SubscriptionReceiverRepository) BulkCreate(ctx context.Context, subscriptionsReceivers []subscriptionreceiver.Relation) error {
	if len(subscriptionsReceivers) == 0 {
		return errors.ErrInvalid.WithMsgf("subscription receiver  cannot be empty")
	}

	srsModel := []model.SubscriptionReceiverRelation{}
	for _, sr := range subscriptionsReceivers {
		srModel := new(model.SubscriptionReceiverRelation)
		srModel.FromDomain(sr)

		srsModel = append(srsModel, *srModel)
	}

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "BulkCreate"),
			attribute.String("db.sql.table", "subscriptions_receivers"),
		}...,
	)

	res, err := s.client.NamedExecContext(ctx, subscriptionReceiverBulkInsertQuery, srsModel)
	if err != nil {
		err = pgc.CheckError(err)
		if errors.Is(err, sql.ErrNoRows) {
			return subscriptionreceiver.NotFoundError{}
		}
		if errors.Is(err, pgc.ErrDuplicateKey) {
			return subscriptionreceiver.ErrDuplicate
		}
		if errors.Is(err, pgc.ErrForeignKeyViolation) {
			return subscriptionreceiver.ErrRelation
		}
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected when inserting subscription receivers")
	}

	return nil
}

func (s *SubscriptionReceiverRepository) List(ctx context.Context, flt subscriptionreceiver.Filter) ([]subscriptionreceiver.Relation, error) {
	var queryBuilder = subscriptionReceiverListQueryBuilder

	if len(flt.SubscriptionIDs) != 0 {
		queryBuilder = queryBuilder.Where("subscription_id = any(?)", pq.Array(flt.SubscriptionIDs))
	}
	if flt.ReceiverID != 0 {
		queryBuilder = queryBuilder.Where("receiver_id = ?", flt.ReceiverID)
	}
	// given map of string from input [lf], look for rows that [lf] exist in labels column in DB
	if len(flt.Labels) != 0 {
		labelsJSON, err := json.Marshal(flt.Labels)
		if err != nil {
			return nil, errors.ErrInvalid.WithMsgf("problem marshalling labels json to string with err: %s", err.Error())
		}
		conditionedJSON := structure.ConditionJSONString(json.RawMessage(labelsJSON))
		queryBuilder = queryBuilder.Where(fmt.Sprintf("labels @> '%s'::jsonb", conditionedJSON))
	}

	if flt.Deleted {
		queryBuilder = queryBuilder.Where("deleted_at IS NOT NULL")
	} else {
		queryBuilder = queryBuilder.Where("deleted_at IS NULL")
	}

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "List"),
			attribute.String("db.sql.table", "subscriptions_receivers"),
		}...,
	)

	rows, err := s.client.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptionsReceiversDomain []subscriptionreceiver.Relation
	for rows.Next() {
		var srModel model.SubscriptionReceiverRelation
		if err := rows.StructScan(&srModel); err != nil {
			return nil, err
		}

		subscriptionsReceiversDomain = append(subscriptionsReceiversDomain, *srModel.ToDomain())
	}

	return subscriptionsReceiversDomain, nil
}

func (r *SubscriptionReceiverRepository) BulkUpsert(ctx context.Context, subscriptionsReceivers []subscriptionreceiver.Relation) error {
	if len(subscriptionsReceivers) == 0 {
		return errors.ErrInvalid.WithMsgf("subscription receiver cannot be empty")
	}

	srsModel := []model.SubscriptionReceiverRelation{}
	for _, sr := range subscriptionsReceivers {
		srModel := new(model.SubscriptionReceiverRelation)
		srModel.FromDomain(sr)

		srsModel = append(srsModel, *srModel)
	}

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "BulkUpsert"),
			attribute.String("db.sql.table", "subscriptions_receivers"),
		}...,
	)

	res, err := r.client.NamedExecContext(ctx, subscriptionReceiverBulkUpsertQuery, srsModel)
	if err != nil {
		err = pgc.CheckError(err)
		if errors.Is(err, sql.ErrNoRows) {
			return subscriptionreceiver.NotFoundError{}
		}
		if errors.Is(err, pgc.ErrDuplicateKey) {
			return subscriptionreceiver.ErrDuplicate
		}
		if errors.Is(err, pgc.ErrForeignKeyViolation) {
			return subscriptionreceiver.ErrRelation
		}
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected when upserting subscription receivers")
	}

	return nil
}

func (r *SubscriptionReceiverRepository) Update(ctx context.Context, rel *subscriptionreceiver.Relation) error {
	if rel == nil {
		return errors.New("subscription receiver relation is nil")
	}

	srModel := new(model.SubscriptionReceiverRelation)
	srModel.FromDomain(*rel)

	var newModel model.SubscriptionReceiverRelation

	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "Update"),
			attribute.String("db.sql.table", "subscriptions_receivers"),
		}...,
	)

	if err := r.client.QueryRowxContext(ctx, subscriptionReceiverUpdateQuery,
		srModel.SubscriptionID,
		srModel.ReceiverID,
		srModel.Labels,
	).StructScan(&newModel); err != nil {
		err = pgc.CheckError(err)
		if errors.Is(err, sql.ErrNoRows) {
			return subscriptionreceiver.NotFoundError{
				SubscriptionID: srModel.SubscriptionID,
				ReceiverID:     srModel.ReceiverID,
			}
		}
		if errors.Is(err, pgc.ErrDuplicateKey) {
			return subscription.ErrDuplicate
		}
		if errors.Is(err, pgc.ErrForeignKeyViolation) {
			return subscriptionreceiver.NotFoundError{
				SubscriptionID: srModel.SubscriptionID,
				ReceiverID:     srModel.ReceiverID,
			}
		}
		return err
	}

	*rel = *newModel.ToDomain()

	return nil
}

func (r *SubscriptionReceiverRepository) BulkSoftDelete(ctx context.Context, flt subscriptionreceiver.DeleteFilter) error {
	if len(flt.Pair) > 0 && flt.SubscriptionID != 0 {
		return errors.New("use either pairs of subscription id and receiver id or a single subscription id")
	}
	var queryBuilder = subscriptionReceiverBulkSoftDeleteQueryBuilder
	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "BulkSoftDelete"),
			attribute.String("db.sql.table", "subscriptions_receivers"),
		}...,
	)

	if len(flt.Pair) > 0 {
		queryBuilder = queryBuilder.Where(fmt.Sprintf("(subscription_id, receiver_id) IN ( VALUES %s)", r.buildSoftDeleteByPairParams(flt.Pair)))
	}
	if flt.SubscriptionID != 0 {
		queryBuilder = queryBuilder.Where("subscription_id = ?", flt.SubscriptionID)
	}

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	res, err := r.client.ExecContext(ctx, query, args...)
	if err != nil {
		err = pgc.CheckError(err)
		if errors.Is(err, sql.ErrNoRows) {
			return subscriptionreceiver.NotFoundError{}
		}
		if errors.Is(err, pgc.ErrDuplicateKey) {
			return subscriptionreceiver.ErrDuplicate
		}
		if errors.Is(err, pgc.ErrForeignKeyViolation) {
			return subscriptionreceiver.ErrRelation
		}
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return subscriptionreceiver.NotFoundError{
			ErrStr: "no data found",
		}
	}

	return nil
}

func (r *SubscriptionReceiverRepository) buildSoftDeleteByPairParams(subscriptionReceivers []subscriptionreceiver.Relation) string {
	// (1,23),(1,50),(1,44)
	params := ""
	for _, sr := range subscriptionReceivers {
		params += fmt.Sprintf("(%d,%d),", sr.SubscriptionID, sr.ReceiverID)
	}
	return strings.TrimSuffix(params, ",")
}

func (r *SubscriptionReceiverRepository) BulkDelete(ctx context.Context, flt subscriptionreceiver.DeleteFilter) error {
	if len(flt.Pair) > 0 && flt.SubscriptionID != 0 {
		return errors.New("use either pairs of subscription id and receiver id or a single subscription id")
	}
	var queryBuilder = sq.Delete("subscriptions_receivers")
	ctx = otelsql.WithCustomAttributes(
		ctx,
		[]attribute.KeyValue{
			attribute.String("db.repository.method", "BulkDelete"),
			attribute.String("db.sql.table", "subscriptions_receivers"),
		}...,
	)

	if len(flt.Pair) > 0 {
		queryBuilder = queryBuilder.Where(fmt.Sprintf("(subscription_id, receiver_id) IN ( VALUES %s)", r.buildSoftDeleteByPairParams(flt.Pair)))
	}
	if flt.SubscriptionID != 0 {
		queryBuilder = queryBuilder.Where("subscription_id = ?", flt.SubscriptionID)
	}

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	res, err := r.client.ExecContext(ctx, query, args...)
	if err != nil {
		err = pgc.CheckError(err)
		if errors.Is(err, sql.ErrNoRows) {
			return subscriptionreceiver.NotFoundError{}
		}
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected when deleting subscription receivers")
	}

	return nil
}
