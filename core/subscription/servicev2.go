package subscription

import (
	"context"

	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/subscriptionreceiver"
	"github.com/goto/siren/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type SubscriptionReceiverService interface {
	List(context.Context, subscriptionreceiver.Filter) ([]subscriptionreceiver.Relation, error)
	BulkCreate(context.Context, []subscriptionreceiver.Relation) error
	BulkUpsert(context.Context, []subscriptionreceiver.Relation) error
	BulkSoftDelete(context.Context, subscriptionreceiver.DeleteFilter) error
	BulkDelete(context.Context, subscriptionreceiver.DeleteFilter) error
}

func (s *Service) rollback(ctx context.Context, err error) error {
	if rbErr := s.repository.Rollback(ctx, err); rbErr != nil {
		return rbErr
	} else {
		return err
	}
}

func (s *Service) ListV2(ctx context.Context, flt Filter) ([]Subscription, error) {
	if flt.SilenceID != "" {
		subscriptionIDs, err := s.logService.ListSubscriptionIDsBySilenceID(ctx, flt.SilenceID)
		if err != nil {
			return nil, err
		}
		flt.IDs = subscriptionIDs
	}

	subscriptions, err := s.repository.List(ctx, flt)
	if err != nil {
		return nil, err
	}

	var subscriptionIDs []uint64
	for _, sub := range subscriptions {
		subscriptionIDs = append(subscriptionIDs, sub.ID)
	}

	subscriptionsReceivers, err := s.subscriptionReceiverService.List(ctx, subscriptionreceiver.Filter{
		SubscriptionIDs: subscriptionIDs,
	})
	if err != nil {
		return nil, err
	}

	var receiversMap = map[uint64][]Receiver{}
	for _, subRcv := range subscriptionsReceivers {
		if len(receiversMap[subRcv.SubscriptionID]) == 0 {
			receiversMap[subRcv.SubscriptionID] = []Receiver{
				{
					ID: subRcv.ReceiverID,
				},
			}
		} else {
			receiversMap[subRcv.SubscriptionID] = append(receiversMap[subRcv.SubscriptionID], Receiver{
				ID: subRcv.ReceiverID,
			})
		}
	}

	// enrich subscription
	if len(subscriptions) > 0 {
		for i := 0; i < len(subscriptions); i++ {
			subscriptions[i].Receivers = receiversMap[subscriptions[i].ID]
		}
	}

	return subscriptions, nil
}

func (s *Service) CreateV2(ctx context.Context, sub *Subscription) error {
	ctx = s.repository.WithTransaction(ctx)
	if err := s.repository.Create(ctx, sub); err != nil {
		var outErr = err
		if errors.Is(err, ErrDuplicate) {
			outErr = errors.ErrConflict.WithMsgf(err.Error())
		}
		if errors.Is(err, ErrRelation) {
			outErr = errors.ErrNotFound.WithMsgf(err.Error())
		}
		return s.rollback(ctx, outErr)
	}

	//for the future, setting up receiver should be done via its own API
	// POST subscriptions/{subscription_id}/receivers
	// consider only creating subscription if no receiver passed
	if len(sub.Receivers) == 0 {
		if err := s.repository.Commit(ctx); err != nil {
			return err
		}
		return nil
	}

	var newReceiverIDs []uint64
	for _, rcv := range sub.Receivers {
		newReceiverIDs = append(newReceiverIDs, rcv.ID)
	}

	receivers, err := s.receiverService.List(ctx, receiver.Filter{
		ReceiverIDs: newReceiverIDs,
	})
	if err != nil {
		return s.rollback(ctx, err)
	}

	// errors out if one of the receiver not found
	if receiversNotFound := s.checkNonExistentReceiver(receivers, newReceiverIDs); len(receiversNotFound) > 0 {
		outErr := errors.ErrInvalid.WithMsgf("cannot found the receivers: %v", receiversNotFound)
		return s.rollback(ctx, outErr)
	}

	var newSubscriptionsReceivers []subscriptionreceiver.Relation
	for _, rcv := range receivers {
		newSubscriptionsReceivers = append(newSubscriptionsReceivers, rcv.ToSubscriptionReceiverRelation(sub.ID))
	}

	if err := s.subscriptionReceiverService.BulkCreate(ctx, newSubscriptionsReceivers); err != nil {
		var outErr = err
		if errors.Is(err, ErrDuplicate) {
			outErr = errors.ErrConflict.WithMsgf(err.Error())
		}
		if errors.Is(err, ErrRelation) {
			outErr = errors.ErrNotFound.WithMsgf(err.Error())
		}
		return s.rollback(ctx, outErr)
	}

	if err := s.repository.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetV2(ctx context.Context, id uint64) (*Subscription, error) {
	sub, err := s.repository.Get(ctx, id)
	if err != nil {
		if errors.As(err, new(NotFoundError)) {
			return nil, errors.ErrNotFound.WithMsgf(err.Error())
		}
		return nil, err
	}

	subscriptionsReceivers, err := s.subscriptionReceiverService.List(ctx, subscriptionreceiver.Filter{
		SubscriptionIDs: []uint64{sub.ID},
	})
	if err != nil {
		return nil, err
	}

	var newReceivers = []Receiver{}
	for _, subRcv := range subscriptionsReceivers {
		newReceivers = append(newReceivers, Receiver{
			ID: subRcv.ReceiverID,
		})
	}

	sub.Receivers = newReceivers

	return sub, nil
}

func (s *Service) checkNonExistentReceiver(sourceOfTruthReceivers []receiver.Receiver, newReceiverIDs []uint64) []uint64 {
	var (
		mapSourceOfTruth       = map[uint64]bool{}
		nonExistentReceiverIDs = []uint64{}
	)
	for _, sot := range sourceOfTruthReceivers {
		mapSourceOfTruth[sot.ID] = true
	}
	for _, rcvID := range newReceiverIDs {
		if _, ok := mapSourceOfTruth[rcvID]; !ok {
			nonExistentReceiverIDs = append(nonExistentReceiverIDs, rcvID)
		}
	}
	return nonExistentReceiverIDs
}

func (s *Service) UpdateV2(ctx context.Context, sub *Subscription) error {
	ctx = s.repository.WithTransaction(ctx)
	if err := s.repository.Update(ctx, sub); err != nil {
		var outErr = err
		if errors.Is(err, ErrDuplicate) {
			outErr = errors.ErrConflict.WithMsgf(err.Error())
		}
		if errors.Is(err, ErrRelation) {
			outErr = errors.ErrNotFound.WithMsgf(err.Error())
		}
		if errors.As(err, new(NotFoundError)) {
			outErr = errors.ErrNotFound.WithMsgf(err.Error())
		}
		return s.rollback(ctx, outErr)
	}

	//for the future, setting up receiver should be done via its own API
	// POST subscriptions/{subscription_id}/receivers
	// consider only creating subscription if no receiver passed
	if len(sub.Receivers) == 0 {
		if err := s.repository.Commit(ctx); err != nil {
			return err
		}
		return nil
	}

	var newReceiverIDs []uint64
	for _, rcv := range sub.Receivers {
		newReceiverIDs = append(newReceiverIDs, rcv.ID)
	}

	receivers, err := s.receiverService.List(ctx, receiver.Filter{
		ReceiverIDs: newReceiverIDs,
	})
	if err != nil {
		return s.rollback(ctx, err)
	}

	// errors out if one of the receiver not found
	if receiversNotFound := s.checkNonExistentReceiver(receivers, newReceiverIDs); len(receiversNotFound) > 0 {
		outErr := errors.ErrInvalid.WithMsgf("cannot found the receivers: %v", receiversNotFound)
		return s.rollback(ctx, outErr)
	}

	var newSubscriptionsReceivers []subscriptionreceiver.Relation
	for _, rcv := range receivers {
		newSubscriptionsReceivers = append(newSubscriptionsReceivers, rcv.ToSubscriptionReceiverRelation(sub.ID))
	}

	existingSubcriptionReceivers, err := s.subscriptionReceiverService.List(ctx, subscriptionreceiver.Filter{
		SubscriptionIDs: []uint64{sub.ID},
	})
	if err != nil {
		return s.rollback(ctx, err)
	}

	toUpsert, toDelete := ClassifyReceivers(newSubscriptionsReceivers, existingSubcriptionReceivers)

	g, ctx := errgroup.WithContext(ctx)

	if len(toUpsert) > 0 {
		g.Go(func() error {
			return s.subscriptionReceiverService.BulkUpsert(ctx, toUpsert)
		})
	}
	if len(toDelete) > 0 {
		g.Go(func() error {
			return s.subscriptionReceiverService.BulkSoftDelete(ctx, subscriptionreceiver.DeleteFilter{
				Pair: toDelete,
			})
		})
	}

	if err := g.Wait(); err != nil {
		return s.rollback(ctx, err)
	}

	if err := s.repository.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteV2(ctx context.Context, id uint64) error {
	ctx = s.repository.WithTransaction(ctx)
	if err := s.subscriptionReceiverService.BulkDelete(ctx, subscriptionreceiver.DeleteFilter{
		SubscriptionID: id,
	}); err != nil {
		return s.rollback(ctx, err)
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return s.rollback(ctx, err)
	}

	if err := s.repository.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Service) MatchByLabelsV2(ctx context.Context, namespaceID uint64, notificationLabels map[string]string) ([]ReceiverView, error) {
	// fetch all subscriptions by matching labels.
	receivers, err := s.repository.MatchLabelsFetchReceivers(ctx, Filter{
		NamespaceID: namespaceID,
		Match:       notificationLabels,
	})
	if err != nil {
		return nil, err
	}

	if len(receivers) == 0 {
		return nil, nil
	}

	for i := 0; i < len(receivers); i++ {
		transformedConfigs, err := s.receiverService.PostHookDBTransformConfigs(ctx, receivers[i].Type, receivers[i].Configurations)
		if err != nil {
			return nil, err
		}
		receivers[i].Configurations = transformedConfigs
	}

	return receivers, nil
}

// ClassifyReceivers compare existing and new receivers of a subscription
func ClassifyReceivers(newReceiver []subscriptionreceiver.Relation, existingReceiver []subscriptionreceiver.Relation) (toUpsert []subscriptionreceiver.Relation, toDelete []subscriptionreceiver.Relation) {
	var newReceiverMap = map[uint64]subscriptionreceiver.Relation{}
	for _, rcv := range newReceiver {
		newReceiverMap[rcv.ReceiverID] = rcv
	}

	var existingReceiverMap = map[uint64]subscriptionreceiver.Relation{}
	for _, rcv := range existingReceiver {
		existingReceiverMap[rcv.ReceiverID] = rcv
	}

	for _, newRcv := range newReceiverMap {
		toUpsert = append(toUpsert, newRcv)
	}

	for existing, existingRcv := range existingReceiverMap {
		if _, ok := newReceiverMap[existing]; !ok {
			toDelete = append(toDelete, existingRcv)
		}
	}
	return
}
