package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/goto/siren/core/silence"
	"github.com/goto/siren/internal/store/model"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/pgc"
	"go.nhat.io/otelsql"
	"go.opentelemetry.io/otel/attribute"
)

const silenceInsertQuery = `
INSERT INTO silences (namespace_id, type, target_id, target_expression, creator, comment, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, now())
RETURNING *
`

var silenceListQueryBuilder = sq.Select(
	"id",
	"namespace_id",
	"type",
	"target_id",
	"target_expression",
	"creator",
	"comment",
	"created_at",
	"deleted_at",
).From("silences")

const silenceSoftDeleteQuery = `
UPDATE silences SET deleted_at=now()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *
`

// SilenceRepository talks to the store to read or insert data
type SilenceRepository struct {
	client *pgc.Client
}

// NewSilenceRepository returns repository struct
func NewSilenceRepository(client *pgc.Client) *SilenceRepository {
	return &SilenceRepository{client}
}

func (r *SilenceRepository) Create(ctx context.Context, s silence.Silence) (string, error) {
	sModel := new(model.Silence)
	sModel.FromDomain(s)

	var newSModel model.Silence

	// Instrumentation attributes
	attrs := []attribute.KeyValue{
		attribute.String("db.method", "Insert"),
		attribute.String("db.sql.table", "silences"),
	}

	if err := r.client.QueryRowxContext(otelsql.AddMeterLabels(ctx, attrs...), silenceInsertQuery,
		sModel.NamespaceID,
		sModel.Type,
		sModel.TargetID,
		sModel.TargetExpression,
		sModel.Creator,
		sModel.Comment,
	).StructScan(&newSModel); err != nil {
		err = pgc.CheckError(err)
		if errors.Is(err, pgc.ErrForeignKeyViolation) {
			return "", errors.ErrInvalid.WithMsgf(err.Error())
		}
		return "", err
	}

	return newSModel.ID, nil
}

func (r *SilenceRepository) List(ctx context.Context, flt silence.Filter) ([]silence.Silence, error) {
	var queryBuilder = silenceListQueryBuilder

	queryBuilder = queryBuilder.Where("deleted_at IS NULL")

	if flt.NamespaceID != 0 {
		queryBuilder = queryBuilder.Where("namespace_id = ?", flt.NamespaceID)
	}

	if flt.SubscriptionID != 0 {
		queryBuilder = queryBuilder.Where("target_id = ?", flt.SubscriptionID)
	}

	if len(flt.Match) != 0 {
		labelsJSON, err := json.Marshal(flt.Match)
		if err != nil {
			return nil, errors.ErrInvalid.WithCausef("problem marshalling json match to string with err: %s", err.Error())
		}
		queryBuilder = queryBuilder.Where(fmt.Sprintf("target_expression @> '%s'::jsonb", string(json.RawMessage(labelsJSON))))
	}

	if len(flt.SubscriptionMatch) != 0 {
		labelsJSON, err := json.Marshal(flt.SubscriptionMatch)
		if err != nil {
			return nil, errors.ErrInvalid.WithCausef("problem marshalling json subscription labels to string with err: %s", err.Error())
		}
		queryBuilder = queryBuilder.Where(fmt.Sprintf("target_expression <@ '%s'::jsonb", string(json.RawMessage(labelsJSON))))
	}

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	// Instrumentation attributes
	attrs := []attribute.KeyValue{
		attribute.String("db.method", "Select *"),
		attribute.String("db.sql.table", "silences"),
	}

	rows, err := r.client.QueryxContext(otelsql.AddMeterLabels(ctx, attrs...), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var silencesDomain []silence.Silence
	for rows.Next() {
		var silenceModel model.Silence
		if err := rows.StructScan(&silenceModel); err != nil {
			return nil, err
		}

		silencesDomain = append(silencesDomain, *silenceModel.ToDomain())
	}

	return silencesDomain, nil
}

func (r *SilenceRepository) Get(ctx context.Context, id string) (silence.Silence, error) {
	queryBuilder := silenceListQueryBuilder.Where("deleted_at IS NULL")

	query, args, err := queryBuilder.Where("id = ?", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return silence.Silence{}, err
	}

	var modelSilence model.Silence

	// Instrumentation attributes
	attrs := []attribute.KeyValue{
		attribute.String("db.method", "Select"),
		attribute.String("db.sql.table", "silences"),
	}
	
	if err := r.client.GetContext(otelsql.AddMeterLabels(ctx, attrs...), &modelSilence, query, args...); err != nil {
		return silence.Silence{}, err
	}

	return *modelSilence.ToDomain(), nil
}

func (r *SilenceRepository) SoftDelete(ctx context.Context, id string) error {
	if _, err := r.client.ExecContext(ctx, silenceSoftDeleteQuery, id); err != nil {
		return err
	}

	return nil
}
