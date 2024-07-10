package notification_test

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	saltlog "github.com/goto/salt/log"
// 	"github.com/goto/siren/core/log"
// 	"github.com/goto/siren/core/notification"
// 	"github.com/goto/siren/core/notification/mocks"
// 	"github.com/stretchr/testify/mock"
// )

// func TestDispatchSingleNotificationServiceService_Dispatch(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		n             []notification.Notification
// 		setup         func(*mocks.Repository, *mocks.AlertRepository, *mocks.Router, *mocks.LogService, *mocks.Queuer)
// 		wantErrString string
// 	}{
// 		{
// 			name:          "should return error if notifications arg is not 1",
// 			n:             []notification.Notification{},
// 			wantErrString: "direct single notification should only accept 1 notification but found 0",
// 		},
// 		{
// 			name: "should return error if repository return error",
// 			n: []notification.Notification{
// 				{
// 					Type: notification.TypeAlert,
// 					Labels: map[string]string{
// 						"k1": "v1",
// 					},
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer) {
// 				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, errors.New("some error"))
// 			},
// 			wantErrString: "some error",
// 		},
// 		{
// 			name: "should return error if notification type is unknown",
// 			n: []notification.Notification{
// 				{
// 					Type:   "random",
// 					Labels: map[string]string{},
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer) {
// 				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
// 			},
// 			wantErrString: "unknown notification type random",
// 		},
// 		{
// 			name: "should return error if log notifications return error",
// 			n: []notification.Notification{
// 				{
// 					Type: notification.TypeAlert,
// 					Labels: map[string]string{
// 						"k1": "v1",
// 					},
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer) {
// 				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
// 				mr.EXPECT().PrepareMessageV2(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, []log.Notification{
// 					{
// 						NamespaceID:    3,
// 						SubscriptionID: 123,
// 						ReceiverID:     12,
// 					},
// 				}, false, nil)
// 				ls.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(errors.New("some error 3"))
// 			},
// 			wantErrString: "failed logging notifications: some error 3",
// 		},
// 		{
// 			name: "should return error if update alerts silence status return error",
// 			n: []notification.Notification{
// 				{
// 					Type: notification.TypeAlert,
// 					Labels: map[string]string{
// 						"k1": "v1",
// 					},
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer) {
// 				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
// 				mr.EXPECT().PrepareMessageV2(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, []log.Notification{
// 					{
// 						NamespaceID:    3,
// 						SubscriptionID: 123,
// 						ReceiverID:     12,
// 					},
// 				}, false, nil)
// 				ls.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
// 				ar.EXPECT().BulkUpdateSilence(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]int64"), mock.AnythingOfType("string")).Return(errors.New("some error 4"))
// 			},
// 			wantErrString: "failed updating silence status: some error 4",
// 		},
// 		{
// 			name: "should return no error if no messages to queue",
// 			n: []notification.Notification{
// 				{
// 					Type: notification.TypeAlert,
// 					Labels: map[string]string{
// 						"k1": "v1",
// 					},
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer) {
// 				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
// 				mr.EXPECT().PrepareMessageV2(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, []log.Notification{
// 					{
// 						NamespaceID:    3,
// 						SubscriptionID: 123,
// 						ReceiverID:     12,
// 					},
// 				}, false, nil)
// 				ls.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
// 				ar.EXPECT().BulkUpdateSilence(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]int64"), mock.AnythingOfType("string")).Return(nil)
// 			},
// 		},
// 		{
// 			name: "should return error if queueing error",
// 			n: []notification.Notification{
// 				{
// 					Type: notification.TypeAlert,
// 					Labels: map[string]string{
// 						"k1": "v1",
// 					},
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer) {
// 				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
// 				mr.EXPECT().PrepareMessageV2(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.Message{
// 					{
// 						ID:              "1234",
// 						NotificationIDs: []string{"n-1234"},
// 					},
// 				}, []log.Notification{
// 					{
// 						NamespaceID:    3,
// 						SubscriptionID: 123,
// 						ReceiverID:     12,
// 					},
// 				}, false, nil)
// 				ls.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
// 				ar.EXPECT().BulkUpdateSilence(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]int64"), mock.AnythingOfType("string")).Return(nil)
// 				q.EXPECT().Enqueue(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(errors.New("some error 5"))
// 			},
// 			wantErrString: "failed enqueuing messages: some error 5",
// 		},

// 		{
// 			name: "should return no error if queueing succeed",
// 			n: []notification.Notification{
// 				{
// 					Type: notification.TypeAlert,
// 					Labels: map[string]string{
// 						"k1": "v1",
// 					},
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer) {
// 				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
// 				mr.EXPECT().PrepareMessageV2(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return([]notification.Message{
// 					{
// 						ID:              "1234",
// 						NotificationIDs: []string{"n-1234"},
// 					},
// 				}, []log.Notification{
// 					{
// 						NamespaceID:    3,
// 						SubscriptionID: 123,
// 						ReceiverID:     12,
// 					},
// 				}, false, nil)
// 				ls.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
// 				ar.EXPECT().BulkUpdateSilence(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]int64"), mock.AnythingOfType("string")).Return(nil)
// 				q.EXPECT().Enqueue(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(nil)
// 			},
// 			wantErrString: "failed enqueuing messages: some error 5",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var (
// 				mockQueuer          = new(mocks.Queuer)
// 				mockRepository      = new(mocks.Repository)
// 				mockAlertRepository = new(mocks.AlertRepository)
// 				mockLogService      = new(mocks.LogService)
// 				mockNotifier        = new(mocks.Notifier)
// 				mockRouter          = new(mocks.Router)
// 			)

// 			if tt.setup != nil {
// 				tt.setup(mockRepository, mockAlertRepository, mockRouter, mockLogService, mockQueuer)
// 			}

// 			s := notification.NewDispatchSingleNotificationService(
// 				notification.Deps{
// 					Cfg: notification.Config{
// 						EnableSilenceFeature: true,
// 					},
// 					Logger:          saltlog.NewNoop(),
// 					Repository:      mockRepository,
// 					Q:               mockQueuer,
// 					AlertRepository: mockAlertRepository,
// 					LogService:      mockLogService,
// 				},
// 				map[string]notification.Notifier{
// 					testType: mockNotifier,
// 				},
// 				map[string]notification.Router{
// 					testType:                      mockRouter,
// 					notification.RouterSubscriber: mockRouter,
// 				},
// 			)
// 			if _, err := s.Dispatch(context.TODO(), tt.n); err != nil {
// 				if err.Error() != tt.wantErrString {
// 					t.Fatalf("Service.DispatchFailure() error = %v, wantErr %s", err, tt.wantErrString)
// 				}
// 			}
// 		})
// 	}
// }

// func TestDispatchSingleNotificationServiceService_DispatchFailureAlert(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		n             []notification.Notification
// 		setup         func(*mocks.Repository, *mocks.AlertRepository, *mocks.Router, *mocks.LogService, *mocks.Queuer)
// 		wantErrString string
// 	}{
// 		{
// 			name: "should return error if dispatchAlert.Validate return error",
// 			n: []notification.Notification{
// 				{
// 					Type:   notification.TypeAlert,
// 					Labels: map[string]string{},
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer) {
// 				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
// 			},
// 			wantErrString: "notification type subscriber should have labels: { 0 alert map[id:] map[] 0s   [] 0001-01-01 00:00:00 +0000 utc []}",
// 		},
// 		{
// 			name: "should return error if dispatchAlert no messages generated but not return subscription not found",
// 			n: []notification.Notification{
// 				{
// 					Type: notification.TypeAlert,
// 					Labels: map[string]string{
// 						"k1": "v1",
// 					},
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer) {
// 				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
// 				mr.EXPECT().PrepareMessageV2(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, nil, false, nil)
// 			},
// 			wantErrString: "something wrong and no messages will be sent with notification: { 0 alert map[id:] map[k1:v1] 0s   [] 0001-01-01 00:00:00 +0000 UTC []}",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var (
// 				mockQueuer          = new(mocks.Queuer)
// 				mockRepository      = new(mocks.Repository)
// 				mockAlertRepository = new(mocks.AlertRepository)
// 				mockLogService      = new(mocks.LogService)
// 				mockNotifier        = new(mocks.Notifier)
// 				mockRouter          = new(mocks.Router)
// 			)

// 			if tt.setup != nil {
// 				tt.setup(mockRepository, mockAlertRepository, mockRouter, mockLogService, mockQueuer)
// 			}

// 			s := notification.NewDispatchSingleNotificationService(
// 				notification.Deps{
// 					Cfg: notification.Config{
// 						EnableSilenceFeature: true,
// 					},
// 					Logger:          saltlog.NewNoop(),
// 					Repository:      mockRepository,
// 					Q:               mockQueuer,
// 					AlertRepository: mockAlertRepository,
// 					LogService:      mockLogService,
// 				},
// 				map[string]notification.Notifier{
// 					testType: mockNotifier,
// 				},
// 				map[string]notification.Router{
// 					testType:                      mockRouter,
// 					notification.RouterSubscriber: mockRouter,
// 				},
// 			)
// 			if _, err := s.Dispatch(context.TODO(), tt.n); err != nil {
// 				if err.Error() != tt.wantErrString {
// 					t.Fatalf("Service.DispatchFailureAlert() error = %v, wantErr %s", err, tt.wantErrString)
// 				}
// 			}
// 		})
// 	}
// }

// func TestDispatchSingleNotificationServiceService_DispatchFailureEvent(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		n             []notification.Notification
// 		setup         func(*mocks.Repository, *mocks.AlertRepository, *mocks.Router, *mocks.LogService, *mocks.Queuer)
// 		wantErrString string
// 	}{
// 		{
// 			name: "should return error if dispatchEvent found no receiver",
// 			n: []notification.Notification{
// 				{
// 					Type: notification.TypeEvent,
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer) {
// 				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
// 			},
// 			wantErrString: "no receivers found",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var (
// 				mockQueuer          = new(mocks.Queuer)
// 				mockRepository      = new(mocks.Repository)
// 				mockAlertRepository = new(mocks.AlertRepository)
// 				mockLogService      = new(mocks.LogService)
// 				mockNotifier        = new(mocks.Notifier)
// 				mockRouter          = new(mocks.Router)
// 			)

// 			if tt.setup != nil {
// 				tt.setup(mockRepository, mockAlertRepository, mockRouter, mockLogService, mockQueuer)
// 			}

// 			s := notification.NewDispatchSingleNotificationService(
// 				notification.Deps{
// 					Cfg: notification.Config{
// 						EnableSilenceFeature: true,
// 					},
// 					Logger:          saltlog.NewNoop(),
// 					Repository:      mockRepository,
// 					Q:               mockQueuer,
// 					AlertRepository: mockAlertRepository,
// 					LogService:      mockLogService,
// 				},
// 				map[string]notification.Notifier{
// 					testType: mockNotifier,
// 				},
// 				map[string]notification.Router{
// 					testType:                      mockRouter,
// 					notification.RouterSubscriber: mockRouter,
// 				},
// 			)
// 			if _, err := s.Dispatch(context.TODO(), tt.n); err != nil {
// 				if err.Error() != tt.wantErrString {
// 					t.Fatalf("Service.DispatchFailureEvent() error = %v, wantErr %s", err, tt.wantErrString)
// 				}
// 			}
// 		})
// 	}
// }
