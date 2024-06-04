package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/internal/api"
	"github.com/goto/siren/pkg/errors"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const notificationAPIScope = "notification_api"

func (s *GRPCServer) validatePostNotificationPayload(receiverSelectors []map[string]string, labels map[string]string) error {
	if len(receiverSelectors) > 0 && len(labels) > 0 {
		return errors.ErrInvalid.WithMsgf("receivers and labels cannot being used at the same time, should be used either one of them")
	}

	return nil
}

func (s *GRPCServer) PostNotification(ctx context.Context, req *sirenv1beta1.PostNotificationRequest) (*sirenv1beta1.PostNotificationResponse, error) {
	idempotencyScope := api.GetHeaderString(ctx, s.headers.IdempotencyScope)
	if idempotencyScope == "" {
		idempotencyScope = notificationAPIScope
	}

	idempotencyKey := api.GetHeaderString(ctx, s.headers.IdempotencyKey)
	if idempotencyKey != "" {
		if notificationID, err := s.notificationService.CheckIdempotency(ctx, idempotencyScope, idempotencyKey); notificationID != "" {
			return &sirenv1beta1.PostNotificationResponse{
				NotificationId: notificationID,
			}, nil
		} else if errors.Is(err, errors.ErrNotFound) {
			s.logger.Debug("no idempotency found with detail", "scope", idempotencyScope, "key", idempotencyKey)
		} else {
			return nil, s.generateRPCErr(fmt.Errorf("error when checking idempotency: %w", err))
		}
	}

	var receiverSelectors = []map[string]string{}
	for _, pbSelector := range req.GetReceivers() {
		var mss = make(map[string]string)
		for k, v := range pbSelector.AsMap() {
			vString, ok := v.(string)
			if !ok {
				err := errors.ErrInvalid.WithMsgf("invalid receiver selectors, value must be string but found %v", v)
				return nil, s.generateRPCErr(err)
			}
			mss[k] = vString
		}
		receiverSelectors = append(receiverSelectors, mss)
	}

	if err := s.validatePostNotificationPayload(receiverSelectors, req.GetLabels()); err != nil {
		return nil, s.generateRPCErr(err)
	}

	var notificationTemplate = template.ReservedName_SystemDefault
	if req.GetTemplate() != "" {
		notificationTemplate = req.GetTemplate()
	}

	notificationIDs, err := s.notificationService.Dispatch(ctx, []notification.Notification{
		{
			Type:              notification.TypeEvent,
			Data:              req.GetData().AsMap(),
			Labels:            req.GetLabels(),
			Template:          notificationTemplate,
			ReceiverSelectors: receiverSelectors,
		},
	}, notification.DispatchKindSingleNotification)
	if err != nil {
		return nil, s.generateRPCErr(err)
	}

	if len(notificationIDs) != 1 {
		return nil, s.generateRPCErr(errors.ErrInternal.WithMsgf("should send 1 notification only but got %d", len(notificationIDs)))
	}

	if idempotencyKey != "" {
		if err := s.notificationService.InsertIdempotency(ctx, idempotencyScope, idempotencyKey, notificationIDs[0]); err != nil {
			return nil, s.generateRPCErr(err)
		}
	}

	return &sirenv1beta1.PostNotificationResponse{
		NotificationId: notificationIDs[0],
	}, nil
}

func (s *GRPCServer) PostBulkNotifications(ctx context.Context, req *sirenv1beta1.PostBulkNotificationsRequest) (*sirenv1beta1.PostBulkNotificationsResponse, error) {
	if len(req.GetNotifications()) == 0 {
		return nil, s.generateRPCErr(errors.ErrInvalid.WithMsgf("no bulk notifications found"))
	}
	idempotencyScope := api.GetHeaderString(ctx, s.headers.IdempotencyScope)
	if idempotencyScope == "" {
		idempotencyScope = notificationAPIScope
	}

	idempotencyKey := api.GetHeaderString(ctx, s.headers.IdempotencyKey)
	if idempotencyKey != "" {
		if notificationIDs, err := s.notificationService.CheckIdempotency(ctx, idempotencyScope, idempotencyKey); notificationIDs != "" {
			splittedNotificationIDs := strings.Split(notificationIDs, ",")
			return &sirenv1beta1.PostBulkNotificationsResponse{
				NotificationIds: splittedNotificationIDs,
			}, nil
		} else if errors.Is(err, errors.ErrNotFound) {
			s.logger.Debug("no idempotency found with detail", "scope", idempotencyScope, "key", idempotencyKey)
		} else {
			return nil, s.generateRPCErr(fmt.Errorf("error when checking idempotency: %w", err))
		}
	}

	notifications := []notification.Notification{}

	for _, nProto := range req.GetNotifications() {
		var notificationTemplate = template.ReservedName_SystemDefault
		if nProto.GetTemplate() != "" {
			notificationTemplate = nProto.GetTemplate()
		}
		notifications = append(notifications, notification.Notification{
			NamespaceID:   nProto.GetNamespaceId(),
			Type:          notification.TypeEvent,
			Data:          nProto.GetData().AsMap(),
			Labels:        nProto.GetLabels(),
			ValidDuration: nProto.GetValidDuration().AsDuration(),
			Template:      notificationTemplate,
		})
	}

	notificationIDs, err := s.notificationService.Dispatch(ctx, notifications, notification.DispatchKindBulkNotification)
	if err != nil {
		return nil, s.generateRPCErr(err)
	}

	if len(notificationIDs) == 0 {
		return &sirenv1beta1.PostBulkNotificationsResponse{
			NotificationIds: []string{},
		}, nil
	}

	if idempotencyKey != "" {
		if err := s.notificationService.InsertIdempotency(ctx, idempotencyScope, idempotencyKey, strings.Join(notificationIDs, ",")); err != nil {
			return nil, s.generateRPCErr(err)
		}
	}

	return &sirenv1beta1.PostBulkNotificationsResponse{
		NotificationIds: notificationIDs,
	}, nil
}

func (s *GRPCServer) ListNotificationMessages(ctx context.Context, req *sirenv1beta1.ListNotificationMessagesRequest) (*sirenv1beta1.ListNotificationMessagesResponse, error) {
	resp, err := s.notificationService.ListNotificationMessages(ctx, req.GetNotificationId())
	if err != nil {
		return nil, s.generateRPCErr(err)
	}

	items := []*sirenv1beta1.NotificationMessage{}
	for _, msg := range resp {
		item, err := msg.ToV1beta1Proto()
		if err != nil {
			return nil, s.generateRPCErr(err)
		}

		items = append(items, item)
	}
	return &sirenv1beta1.ListNotificationMessagesResponse{
		Messages: items,
	}, nil
}

func (s *GRPCServer) ListNotifications(ctx context.Context, req *sirenv1beta1.ListNotificationsRequest) (*sirenv1beta1.ListNotificationsResponse, error) {
	notifications, err := s.notificationService.List(ctx, notification.Filter{
		Type:             req.GetType(),
		Template:         req.GetTemplate(),
		Labels:           req.GetLabels(),
		ReceiverSelector: req.GetReceiverSelectors(),
	})
	if err != nil {
		return nil, s.generateRPCErr(err)
	}

	items := []*sirenv1beta1.Notification{}

	for _, notification := range notifications {
		data, err := structpb.NewStruct(notification.Data)
		if err != nil {
			return nil, s.generateRPCErr(err)
		}

		item := &sirenv1beta1.Notification{
			Id:            notification.ID,
			NamespaceId:   notification.NamespaceID,
			Type:          notification.Type,
			Data:          data,
			Labels:        notification.Labels,
			ValidDuration: durationpb.New(notification.ValidDuration),
			Template:      notification.Template,
			CreateAt:      timestamppb.New(notification.CreatedAt),
			UniqueKey:     notification.UniqueKey,
		}
		items = append(items, item)
	}
	return &sirenv1beta1.ListNotificationsResponse{
		Notifications: items,
	}, nil
}
