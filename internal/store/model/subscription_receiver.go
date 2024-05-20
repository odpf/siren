package model

import (
	"database/sql"
	"time"

	"github.com/goto/siren/core/subscriptionreceiver"
	"github.com/goto/siren/pkg/pgc"
)

type SubscriptionReceiverRelation struct {
	ID             uint64              `db:"id"`
	SubscriptionID uint64              `db:"subscription_id"`
	ReceiverID     uint64              `db:"receiver_id"`
	Labels         pgc.StringStringMap `db:"labels"`
	CreatedAt      time.Time           `db:"created_at"`
	UpdatedAt      time.Time           `db:"updated_at"`
	DeletedAt      sql.NullTime        `db:"deleted_at"`
}

func (s *SubscriptionReceiverRelation) FromDomain(rel subscriptionreceiver.Relation) {
	s.ID = rel.ID
	s.SubscriptionID = rel.SubscriptionID
	s.ReceiverID = rel.ReceiverID
	s.Labels = rel.Labels
	s.CreatedAt = rel.CreatedAt
	s.UpdatedAt = rel.UpdatedAt

	if rel.DeletedAt.IsZero() {
		s.DeletedAt = sql.NullTime{Valid: false}
	} else {
		s.DeletedAt = sql.NullTime{Time: rel.DeletedAt, Valid: true}
	}
}

func (s *SubscriptionReceiverRelation) ToDomain() *subscriptionreceiver.Relation {
	return &subscriptionreceiver.Relation{
		ID:             s.ID,
		SubscriptionID: s.SubscriptionID,
		ReceiverID:     s.ReceiverID,
		Labels:         s.Labels,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
		DeletedAt:      s.DeletedAt.Time,
	}
}
