package notification_test

import (
	"context"
	"errors"
	"testing"

	"github.com/goto/salt/log"
	"github.com/stretchr/testify/mock"

	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/core/notification/mocks"
)

const testReceiverType = "test"

func TestHandler_SingleMessageHandler(t *testing.T) {
	testCases := []struct {
		name       string
		messages   []notification.Message
		setup      func(*mocks.Queuer, *mocks.Notifier)
		wantErrStr string
	}{
		{
			name: "return error if plugin type is not supported",
			messages: []notification.Message{
				{
					ReceiverType: "random",
				},
			},
			setup: func(q *mocks.Queuer, _ *mocks.Notifier) {
				q.EXPECT().ErrorCallback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(nil)
				q.EXPECT().ErrorCallback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(nil)
			},
			wantErrStr: "unsupported receiver type: \"random\" on handler ",
		},
		{
			name: "return error if post hook transform config is failing and error callback success",
			messages: []notification.Message{
				{
					ReceiverType: testType,
				},
			},
			setup: func(q *mocks.Queuer, n *mocks.Notifier) {
				n.EXPECT().PostHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("map[string]interface {}")).Return(nil, errors.New("some error"))
				q.EXPECT().ErrorCallback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(nil)
			},
			wantErrStr: "some error",
		},
		{
			name: "return error if post hook transform config is failing and error callback is failing",
			messages: []notification.Message{
				{
					ReceiverType: testType,
				},
			},
			setup: func(q *mocks.Queuer, n *mocks.Notifier) {
				n.EXPECT().PostHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("map[string]interface {}")).Return(nil, errors.New("some error"))
				q.EXPECT().ErrorCallback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(errors.New("some error"))
			},
			wantErrStr: "failed to execute error callback with receiver type test and error some error",
		},
		{
			name: "return error if send message return error and error handler queue return error",
			messages: []notification.Message{
				{
					ReceiverType: testType,
				},
			},
			setup: func(q *mocks.Queuer, n *mocks.Notifier) {
				n.EXPECT().PostHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("map[string]interface {}")).Return(map[string]any{}, nil)
				n.EXPECT().Send(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(false, errors.New("some error"))
				q.EXPECT().ErrorCallback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(errors.New("some error"))
			},
			wantErrStr: "failed to execute error callback with receiver type test and error some error",
		},
		{
			name: "return error if send message return error and error handler queue return no error",
			messages: []notification.Message{
				{
					ReceiverType: testType,
				},
			},
			setup: func(q *mocks.Queuer, n *mocks.Notifier) {
				n.EXPECT().PostHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("map[string]interface {}")).Return(map[string]any{}, nil)
				n.EXPECT().Send(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(false, errors.New("some error"))
				q.EXPECT().ErrorCallback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(nil)
			},
			wantErrStr: "some error",
		},
		{
			name: "return error if send message success and success handler queue return error",
			messages: []notification.Message{
				{
					ReceiverType: testType,
				},
			},
			setup: func(q *mocks.Queuer, n *mocks.Notifier) {
				n.EXPECT().PostHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("map[string]interface {}")).Return(map[string]any{}, nil)
				n.EXPECT().Send(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(false, nil)
				q.EXPECT().SuccessCallback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(errors.New("some error"))
			},
			wantErrStr: "some error",
		},
		{
			name: "return no error if send message success and success handler queue return no error",
			messages: []notification.Message{
				{
					ReceiverType: testType,
				},
			},
			setup: func(q *mocks.Queuer, n *mocks.Notifier) {
				n.EXPECT().PostHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("map[string]interface {}")).Return(map[string]any{}, nil)
				n.EXPECT().Send(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(false, nil)
				q.EXPECT().SuccessCallback(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Message")).Return(nil)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				mockQueue    = new(mocks.Queuer)
				mockNotifier = new(mocks.Notifier)
			)

			if tc.setup != nil {
				tc.setup(mockQueue, mockNotifier)
			}

			h := notification.NewHandler(notification.HandlerConfig{}, log.NewNoop(), mockQueue, map[string]notification.Notifier{
				testReceiverType: mockNotifier,
			})
			if err := h.SingleMessageHandler(context.TODO(), &tc.messages[0]); err != nil {
				if err.Error() != tc.wantErrStr {
					t.Errorf("Handler.messageHandler() error = %s, wantErr = %s", err.Error(), tc.wantErrStr)
				}
			}

			mockQueue.AssertExpectations(t)
			mockNotifier.AssertExpectations(t)
		})
	}
}
