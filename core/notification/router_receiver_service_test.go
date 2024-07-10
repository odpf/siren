package notification_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// 	"github.com/google/go-cmp/cmp/cmpopts"
// 	saltlog "github.com/goto/salt/log"
// 	"github.com/goto/siren/core/log"
// 	"github.com/goto/siren/core/notification"
// 	"github.com/goto/siren/core/notification/mocks"
// 	"github.com/goto/siren/core/receiver"
// 	"github.com/goto/siren/core/template"
// 	"github.com/goto/siren/pkg/errors"
// 	"github.com/stretchr/testify/mock"
// )

// func TestRouterReceiverService_PrepareMessage(t *testing.T) {
// 	var notificationID = "1234-5678"
// 	tests := []struct {
// 		name    string
// 		setup   func(*mocks.ReceiverService, *mocks.Notifier)
// 		n       notification.Notification
// 		want    []notification.Message
// 		want1   []log.Notification
// 		want2   bool
// 		wantErr bool
// 	}{
// 		{
// 			name: "should return error if receiver service return error",
// 			n:    notification.Notification{},
// 			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
// 				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return(nil, errors.New("some error"))
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "should return error if receiver type is unknown",
// 			n: notification.Notification{
// 				ID: notificationID,
// 				ReceiverSelectors: []map[string]string{
// 					{
// 						"id": "11",
// 					},
// 				},
// 			},
// 			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
// 				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{}, nil)
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "should return error if init message return error",
// 			n: notification.Notification{
// 				ID: notificationID,
// 				ReceiverSelectors: []map[string]string{
// 					{
// 						"id": "11",
// 					},
// 				},
// 			},
// 			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
// 				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
// 					{
// 						ID:   11,
// 						Type: testType},
// 				}, nil)
// 				n.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("map[string]interface {}")).Return(nil, errors.New("some error"))
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "should return error if template is not passed",
// 			n: notification.Notification{
// 				ID: notificationID,
// 				ReceiverSelectors: []map[string]string{
// 					{
// 						"id": "11",
// 					},
// 				},
// 			},
// 			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
// 				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
// 					{
// 						ID:   11,
// 						Type: testType,
// 					},
// 				}, nil)
// 				n.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("map[string]interface {}")).Return(map[string]any{}, nil)
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "should return no error if all flow passed",
// 			n: notification.Notification{
// 				ID: notificationID,
// 				ReceiverSelectors: []map[string]string{
// 					{
// 						"id": "11",
// 					},
// 				},
// 				Template: template.ReservedName_SystemDefault,
// 			},
// 			setup: func(rs *mocks.ReceiverService, n *mocks.Notifier) {
// 				rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
// 					{
// 						ID:   11,
// 						Type: testType,
// 					},
// 				}, nil)
// 				n.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("map[string]interface {}")).Return(map[string]any{}, nil)
// 				n.EXPECT().GetSystemDefaultTemplate().Return("")
// 			},
// 			want: []notification.Message{
// 				{
// 					Status:          notification.MessageStatusEnqueued,
// 					NotificationIDs: []string{notificationID},
// 					ReceiverType:    testType,
// 					Configs:         map[string]any{},
// 					Details:         map[string]any{"notification_type": string("")},
// 					MaxTries:        3,
// 				},
// 			},
// 			want1: []log.Notification{{
// 				NotificationID: notificationID,
// 				ReceiverID:     11,
// 			}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var (
// 				mockReceiverService = new(mocks.ReceiverService)
// 				mockNotifier        = new(mocks.Notifier)
// 				mockTemplateService = new(mocks.TemplateService)
// 			)
// 			s := notification.NewRouterReceiverService(
// 				notification.Deps{
// 					Cfg: notification.Config{
// 						MaxMessagesReceiverFlow: 10,
// 						MaxNumReceiverSelectors: 10,
// 					},
// 					Logger:          saltlog.NewNoop(),
// 					ReceiverService: mockReceiverService,
// 					TemplateService: mockTemplateService,
// 				},
// 				map[string]notification.Notifier{
// 					testType: mockNotifier,
// 				},
// 			)
// 			if tt.setup != nil {
// 				tt.setup(mockReceiverService, mockNotifier)
// 			}
// 			got, got1, got2, err := s.PrepareMessage(context.TODO(), tt.n)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RouterReceiverService.PrepareMessage() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if diff := cmp.Diff(got, tt.want,
// 				cmpopts.IgnoreFields(notification.Message{}, "ID", "CreatedAt", "UpdatedAt"),
// 				cmpopts.IgnoreUnexported(notification.Message{})); diff != "" {
// 				t.Errorf("RouterReceiverService.PrepareMessage() diff = %v", diff)
// 			}
// 			if diff := cmp.Diff(got1, tt.want1); diff != "" {
// 				t.Errorf("RouterReceiverService.PrepareMessage() diff = %v", diff)
// 			}
// 			if got2 != tt.want2 {
// 				t.Errorf("RouterReceiverService.PrepareMessage() got2 = %v, want %v", got2, tt.want2)
// 			}
// 		})
// 	}
// }
