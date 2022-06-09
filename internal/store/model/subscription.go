package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/odpf/siren/core/subscription"
)

type ReceiverMetadata struct {
	Id            uint64            `json:"id"`
	Configuration map[string]string `json:"configuration"`
}

type ReceiverMetadataList []ReceiverMetadata

func (list *ReceiverMetadataList) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), &list)
}

func (list ReceiverMetadataList) Value() (driver.Value, error) {
	val, err := json.Marshal(list)
	return string(val), err
}

type Subscription struct {
	Id          uint64 `gorm:"primarykey"`
	Namespace   *Namespace
	NamespaceId uint64
	Urn         string               `gorm:"unique"`
	Receiver    ReceiverMetadataList `gorm:"type:jsonb" sql:"type:jsonb" `
	Match       StringStringMap      `gorm:"type:jsonb" sql:"type:jsonb" `
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s *Subscription) FromDomain(sub *subscription.Subscription) {
	if s == nil {
		return
	}
	s.Id = sub.Id
	s.Urn = sub.Urn
	s.NamespaceId = sub.Namespace
	s.Match = sub.Match
	s.Receiver = make([]ReceiverMetadata, 0)
	for _, item := range sub.Receivers {
		receiver := ReceiverMetadata{
			Id:            item.Id,
			Configuration: item.Configuration,
		}
		s.Receiver = append(s.Receiver, receiver)
	}
	s.CreatedAt = sub.CreatedAt
	s.UpdatedAt = sub.UpdatedAt
}

func (s *Subscription) ToDomain() *subscription.Subscription {
	if s == nil {
		return nil
	}
	receivers := make([]subscription.ReceiverMetadata, 0)
	for _, item := range s.Receiver {
		receiver := subscription.ReceiverMetadata{
			Id:            item.Id,
			Configuration: item.Configuration,
		}
		receivers = append(receivers, receiver)
	}
	subscription := &subscription.Subscription{
		Id:        s.Id,
		Urn:       s.Urn,
		Match:     s.Match,
		Namespace: s.NamespaceId,
		Receivers: receivers,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
	return subscription
}
