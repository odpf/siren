package v1beta1

import (
	"context"

	"github.com/goto/siren/core/subscription"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *GRPCServer) ListSubscriptions(ctx context.Context, req *sirenv1beta1.ListSubscriptionsRequest) (*sirenv1beta1.ListSubscriptionsResponse, error) {
	// NOTE: only string value that is queryable with this approach
	var metadataQuery = map[string]any{}
	if len(req.GetMetadata()) == 0 {
		metadataQuery = nil
	} else {
		for k, v := range req.GetMetadata() {
			metadataQuery[k] = v
		}
	}

	var (
		subscriptions []subscription.Subscription
		err           error
	)

	if s.cfg.subscriptionV2Enabled {
		subscriptions, err = s.subscriptionService.ListV2(ctx, subscription.Filter{
			NamespaceID:       req.GetNamespaceId(),
			SilenceID:         req.GetSilenceId(),
			Metadata:          metadataQuery,
			Match:             req.GetMatch(),
			NotificationMatch: req.GetNotificationMatch(),
		})
	} else {
		subscriptions, err = s.subscriptionService.List(ctx, subscription.Filter{
			NamespaceID:       req.GetNamespaceId(),
			SilenceID:         req.GetSilenceId(),
			Metadata:          metadataQuery,
			Match:             req.GetMatch(),
			NotificationMatch: req.GetNotificationMatch(),
		})
	}
	if err != nil {
		return nil, s.generateRPCErr(err)
	}

	items := []*sirenv1beta1.Subscription{}

	for _, sub := range subscriptions {

		receiverMetadatasPB := make([]*sirenv1beta1.ReceiverMetadata, 0)
		for _, item := range sub.Receivers {
			configMapPB, err := structpb.NewStruct(item.Configuration)
			if err != nil {
				return nil, err
			}

			receiverMetadatasPB = append(receiverMetadatasPB, &sirenv1beta1.ReceiverMetadata{
				Id:            item.ID,
				Configuration: configMapPB,
			})
		}

		metadata, err := structpb.NewStruct(sub.Metadata)
		if err != nil {
			return nil, s.generateRPCErr(err)
		}

		item := &sirenv1beta1.Subscription{
			Id:        sub.ID,
			Urn:       sub.URN,
			Namespace: sub.Namespace,
			Match:     sub.Match,
			Receivers: receiverMetadatasPB,
			Metadata:  metadata,
			CreatedBy: sub.CreatedBy,
			UpdatedBy: sub.UpdatedBy,
			CreatedAt: timestamppb.New(sub.CreatedAt),
			UpdatedAt: timestamppb.New(sub.UpdatedAt),
		}
		items = append(items, item)
	}
	return &sirenv1beta1.ListSubscriptionsResponse{
		Subscriptions: items,
	}, nil
}

func (s *GRPCServer) CreateSubscription(ctx context.Context, req *sirenv1beta1.CreateSubscriptionRequest) (*sirenv1beta1.CreateSubscriptionResponse, error) {
	sub := &subscription.Subscription{
		Namespace: req.GetNamespace(),
		URN:       req.GetUrn(),
		Receivers: getReceiverMetadataListInDomainObject(req.GetReceivers()),
		Match:     req.GetMatch(),
		Metadata:  req.GetMetadata().AsMap(),
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.CreatedBy,
	}

	var err error

	if s.cfg.subscriptionV2Enabled {
		err = s.subscriptionService.CreateV2(ctx, sub)
	} else {
		err = s.subscriptionService.Create(ctx, sub)
	}
	if err != nil {
		return nil, s.generateRPCErr(err)
	}

	return &sirenv1beta1.CreateSubscriptionResponse{
		Id: sub.ID,
	}, nil
}

func (s *GRPCServer) GetSubscription(ctx context.Context, req *sirenv1beta1.GetSubscriptionRequest) (*sirenv1beta1.GetSubscriptionResponse, error) {
	var (
		sub *subscription.Subscription
		err error
	)

	if s.cfg.subscriptionV2Enabled {
		sub, err = s.subscriptionService.GetV2(ctx, req.GetId())
	} else {
		sub, err = s.subscriptionService.Get(ctx, req.GetId())
	}
	if err != nil {
		return nil, s.generateRPCErr(err)
	}

	receivers := make([]*sirenv1beta1.ReceiverMetadata, 0)
	for _, receiverMetadataItem := range sub.Receivers {
		configMapPB, err := structpb.NewStruct(receiverMetadataItem.Configuration)
		if err != nil {
			return nil, err
		}

		receivers = append(receivers, &sirenv1beta1.ReceiverMetadata{
			Id:            receiverMetadataItem.ID,
			Configuration: configMapPB,
		})
	}

	metadata, err := structpb.NewStruct(sub.Metadata)
	if err != nil {
		return nil, s.generateRPCErr(err)
	}

	return &sirenv1beta1.GetSubscriptionResponse{
		Subscription: &sirenv1beta1.Subscription{
			Id:        sub.ID,
			Urn:       sub.URN,
			Namespace: sub.Namespace,
			Match:     sub.Match,
			Receivers: receivers,
			Metadata:  metadata,
			CreatedBy: sub.CreatedBy,
			UpdatedBy: sub.UpdatedBy,
			CreatedAt: timestamppb.New(sub.CreatedAt),
			UpdatedAt: timestamppb.New(sub.UpdatedAt),
		},
	}, nil
}

func (s *GRPCServer) UpdateSubscription(ctx context.Context, req *sirenv1beta1.UpdateSubscriptionRequest) (*sirenv1beta1.UpdateSubscriptionResponse, error) {
	sub := &subscription.Subscription{
		ID:        req.GetId(),
		Namespace: req.GetNamespace(),
		URN:       req.GetUrn(),
		Receivers: getReceiverMetadataListInDomainObject(req.GetReceivers()),
		Match:     req.GetMatch(),
		Metadata:  req.Metadata.AsMap(),
		UpdatedBy: req.UpdatedBy,
	}

	var err error

	if s.cfg.subscriptionV2Enabled {
		err = s.subscriptionService.UpdateV2(ctx, sub)
	} else {
		err = s.subscriptionService.Update(ctx, sub)
	}
	if err != nil {
		return nil, s.generateRPCErr(err)
	}

	return &sirenv1beta1.UpdateSubscriptionResponse{
		Id: sub.ID,
	}, nil
}

func (s *GRPCServer) DeleteSubscription(ctx context.Context, req *sirenv1beta1.DeleteSubscriptionRequest) (*sirenv1beta1.DeleteSubscriptionResponse, error) {
	var err error

	if s.cfg.subscriptionV2Enabled {
		err = s.subscriptionService.DeleteV2(ctx, req.GetId())
	} else {
		err = s.subscriptionService.Delete(ctx, req.GetId())
	}
	if err != nil {
		return nil, s.generateRPCErr(err)
	}
	return &sirenv1beta1.DeleteSubscriptionResponse{}, nil
}

func getReceiverMetadataListInDomainObject(domainReceivers []*sirenv1beta1.ReceiverMetadata) []subscription.Receiver {
	receivers := make([]subscription.Receiver, 0)
	for _, item := range domainReceivers {
		receivers = append(receivers, subscription.Receiver{
			ID:            item.Id,
			Configuration: item.Configuration.AsMap(),
		})
	}
	return receivers
}
