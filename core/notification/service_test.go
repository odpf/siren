package notification_test

import (
	"context"
	"testing"

	saltlog "github.com/goto/salt/log"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/core/notification/mocks"
	"github.com/goto/siren/pkg/errors"
	"github.com/stretchr/testify/mock"
)

const (
	testType = "test"
)

func TestService_DispatchFailure(t *testing.T) {
	tests := []struct {
		name    string
		n       []notification.Notification
		setup   func([]notification.Notification, *mocks.Repository, *mocks.AlertRepository, *mocks.LogService, *mocks.AlertService, *mocks.Queuer, *mocks.Dispatcher)
		wantErr bool
	}{
		{
			name: "should return error if repository return error",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, _ *mocks.LogService, _ *mocks.AlertService, _ *mocks.Queuer, _ *mocks.Dispatcher) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if dispatcher service return error",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, _ *mocks.LogService, _ *mocks.AlertService, _ *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
				d.EXPECT().Dispatch(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return(nil, errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if dispatcher service return empty results",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, _ *mocks.LogService, _ *mocks.AlertService, _ *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*errors.errorString")).Return(nil)
				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
				d.EXPECT().Dispatch(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, nil)
			},
			wantErr: true,
		},
		{
			name: "should return error if log notifications return error",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, _ *mocks.AlertRepository, l *mocks.LogService, _ *mocks.AlertService, _ *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*fmt.wrapError")).Return(nil)
				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
				d.EXPECT().Dispatch(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return([]string{"123"}, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if update alerts silence status return error",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, ar *mocks.AlertRepository, l *mocks.LogService, a *mocks.AlertService, _ *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*fmt.wrapError")).Return(nil)
				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
				d.EXPECT().Dispatch(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]string{"123"}, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
				ar.EXPECT().BulkUpdateSilence(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]int64"), mock.AnythingOfType("string")).Return(errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "should return error if enqueue return error",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, ar *mocks.AlertRepository, l *mocks.LogService, a *mocks.AlertService, q *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*fmt.wrapError")).Return(nil)
				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(notification.Notification{}, nil)
				d.EXPECT().Dispatch(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return([]string{"123"}, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
				ar.EXPECT().BulkUpdateSilence(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]int64"), mock.AnythingOfType("string")).Return(errors.New("some error"))
				q.EXPECT().Enqueue(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(errors.New("some error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				mockQueuer          = new(mocks.Queuer)
				mockRepository      = new(mocks.Repository)
				mockAlertRepository = new(mocks.AlertRepository)
				mockLogService      = new(mocks.LogService)
				mockAlertService    = new(mocks.AlertService)
				mockDispatcher      = new(mocks.Dispatcher)
			)

			if tt.setup != nil {
				mockRepository.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
				tt.setup(tt.n, mockRepository, mockAlertRepository, mockLogService, mockAlertService, mockQueuer, mockDispatcher)
			}

			s := notification.NewService(
				notification.Deps{
					Cfg: notification.Config{
						EnableSilenceFeature: true,
					},
					Logger:          saltlog.NewNoop(),
					Repository:      mockRepository,
					Q:               mockQueuer,
					AlertRepository: mockAlertRepository,
					LogService:      mockLogService,
				},
				map[string]notification.Dispatcher{
					testType: mockDispatcher,
				},
			)
			if _, err := s.Dispatch(context.TODO(), tt.n, notification.DispatchKindSingleNotification); (err != nil) != tt.wantErr {
				t.Errorf("Service.DispatchFailure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_DispatchSuccess(t *testing.T) {
	tests := []struct {
		name    string
		n       []notification.Notification
		setup   func([]notification.Notification, *mocks.Repository, *mocks.AlertRepository, *mocks.LogService, *mocks.AlertService, *mocks.Queuer, *mocks.Dispatcher)
		wantErr bool
	}{
		{
			name: "should return no error if enqueue success",
			n: []notification.Notification{
				{
					Type: notification.TypeAlert,
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			setup: func(n []notification.Notification, r *mocks.Repository, ar *mocks.AlertRepository, l *mocks.LogService, a *mocks.AlertService, q *mocks.Queuer, d *mocks.Dispatcher) {
				r.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(n[0], nil)
				d.EXPECT().Dispatch(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]string{"123"}, nil)
				l.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
				ar.EXPECT().BulkUpdateSilence(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]int64"), mock.AnythingOfType("string")).Return(nil)
				q.EXPECT().Enqueue(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				mockQueuer          = new(mocks.Queuer)
				mockRepository      = new(mocks.Repository)
				mockAlertRepository = new(mocks.AlertRepository)
				mockLogService      = new(mocks.LogService)
				mockAlertService    = new(mocks.AlertService)
				mockDispatcher      = new(mocks.Dispatcher)
			)

			if tt.setup != nil {
				mockRepository.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
				mockRepository.EXPECT().Commit(mock.AnythingOfType("context.todoCtx")).Return(nil)
				tt.setup(tt.n, mockRepository, mockAlertRepository, mockLogService, mockAlertService, mockQueuer, mockDispatcher)
			}

			s := notification.NewService(
				notification.Deps{
					Cfg: notification.Config{
						EnableSilenceFeature: true,
					},
					Logger:          saltlog.NewNoop(),
					Repository:      mockRepository,
					Q:               mockQueuer,
					AlertRepository: mockAlertRepository,
					LogService:      mockLogService,
				},
				map[string]notification.Dispatcher{
					testType: mockDispatcher,
				},
			)
			if _, err := s.Dispatch(context.TODO(), tt.n, testType); (err != nil) != tt.wantErr {
				t.Errorf("Service.Dispatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
