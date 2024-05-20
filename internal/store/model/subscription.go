package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/pkg/pgc"
)

type SubscriptionReceiver struct {
	ID            uint64         `json:"id"`
	Configuration map[string]any `json:"configuration"`
}

type SubscriptionReceivers []SubscriptionReceiver

func (list *SubscriptionReceivers) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), &list)
}

func (list SubscriptionReceivers) Value() (driver.Value, error) {
	val, err := json.Marshal(list)
	return string(val), err
}

type Subscription struct {
	ID          uint64                `db:"id"`
	NamespaceID uint64                `db:"namespace_id"`
	URN         string                `db:"urn"`
	Receiver    SubscriptionReceivers `db:"receiver"`
	Match       pgc.StringStringMap   `db:"match"`
	Metadata    pgc.StringAnyMap      `db:"metadata"`
	CreatedBy   sql.NullString        `db:"created_by"`
	UpdatedBy   sql.NullString        `db:"updated_by"`
	CreatedAt   time.Time             `db:"created_at"`
	UpdatedAt   time.Time             `db:"updated_at"`
}

func (s *Subscription) FromDomain(sub subscription.Subscription) {
	s.ID = sub.ID
	s.URN = sub.URN
	s.NamespaceID = sub.Namespace
	s.Match = sub.Match
	s.Receiver = make([]SubscriptionReceiver, 0)
	for _, item := range sub.Receivers {
		receiver := SubscriptionReceiver{
			ID:            item.ID,
			Configuration: item.Configuration,
		}
		s.Receiver = append(s.Receiver, receiver)
	}

	s.Metadata = pgc.StringAnyMap(sub.Metadata)
	if sub.CreatedBy == "" {
		s.CreatedBy = sql.NullString{Valid: false}
	} else {
		s.CreatedBy = sql.NullString{String: sub.CreatedBy, Valid: true}
	}
	if sub.UpdatedBy == "" {
		s.UpdatedBy = sql.NullString{Valid: false}
	} else {
		s.UpdatedBy = sql.NullString{String: sub.UpdatedBy, Valid: true}
	}
	s.CreatedAt = sub.CreatedAt
	s.UpdatedAt = sub.UpdatedAt
}

func (s *Subscription) ToDomain() *subscription.Subscription {
	receivers := make([]subscription.Receiver, 0)
	for _, item := range s.Receiver {
		receiver := subscription.Receiver{
			ID:            item.ID,
			Configuration: item.Configuration,
		}
		receivers = append(receivers, receiver)
	}

	return &subscription.Subscription{
		ID:        s.ID,
		URN:       s.URN,
		Match:     s.Match,
		Namespace: s.NamespaceID,
		Receivers: receivers,
		Metadata:  s.Metadata,
		CreatedBy: s.CreatedBy.String,
		UpdatedBy: s.UpdatedBy.String,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

type ReceiverView struct {
	ID             uint64              `db:"id"`
	Name           string              `db:"name"`
	Type           string              `db:"type"`
	Labels         pgc.StringStringMap `db:"labels"`
	Configurations pgc.StringAnyMap    `db:"configurations"`
	ParentID       sql.NullInt64       `db:"parent_id"`
	CreatedAt      time.Time           `db:"created_at"`
	UpdatedAt      time.Time           `db:"updated_at"`

	// This is optional and used as backreference for match by labels case
	SubscriptionID uint64              `db:"subscription_id"`
	Match          pgc.StringStringMap `db:"match"`
}

func (rcv *ReceiverView) FromDomain(t subscription.ReceiverView) {
	rcv.ID = t.ID
	rcv.Name = t.Name
	rcv.Type = t.Type
	rcv.Labels = t.Labels
	rcv.Configurations = pgc.StringAnyMap(t.Configurations)
	rcv.ParentID = sql.NullInt64{
		Valid: true,
		// since postgres does not support unsigned integer and ids in siren is autogenerated by postgres and never be < 0 (bigserial)
		// this operation would be safe and no overflow is expected
		Int64: int64(t.ParentID),
	}
	rcv.CreatedAt = t.CreatedAt
	rcv.UpdatedAt = t.UpdatedAt
	rcv.SubscriptionID = t.SubscriptionID
	rcv.Match = t.Match
}

func (rcv *ReceiverView) ToDomain() *subscription.ReceiverView {
	return &subscription.ReceiverView{
		ID:             rcv.ID,
		Name:           rcv.Name,
		Type:           rcv.Type,
		Labels:         rcv.Labels,
		Configurations: rcv.Configurations,

		// since postgres does not support unsigned integer and ids in siren is autogenerated by postgres and never be < 0 (bigserial)
		// this operation would be safe and no overflow is expected
		ParentID:       uint64(rcv.ParentID.Int64),
		CreatedAt:      rcv.CreatedAt,
		UpdatedAt:      rcv.UpdatedAt,
		SubscriptionID: rcv.SubscriptionID,
		Match:          rcv.Match,
	}
}
