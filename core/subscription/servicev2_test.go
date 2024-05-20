package subscription_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/core/subscription/mocks"
	"github.com/goto/siren/core/subscriptionreceiver"
	"github.com/goto/siren/pkg/errors"
	"github.com/stretchr/testify/mock"
)

func TestService_ListV2(t *testing.T) {
	type testCase struct {
		Description string
		Setup       func(*mocks.Repository, *mocks.SubscriptionReceiverService)
		ErrString   string
	}
	var (
		ctx       = context.TODO()
		testCases = []testCase{
			{
				Description: "should return error if subscription list repository return error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscription.Filter")).Return(nil, errors.New("some error"))
				},
				ErrString: "some error",
			},
			{
				Description: "should return nil error if all passed",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscription.Filter")).Return([]subscription.Subscription{}, nil)
					srs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.Filter")).Return([]subscriptionreceiver.Relation{}, nil)
				},
			},
			{
				Description: "should return error if subscription service list return error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscription.Filter")).Return([]subscription.Subscription{}, nil)
					srs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.Filter")).Return(nil, errors.New("some error 2"))
				},
				ErrString: "some error 2",
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock          = new(mocks.Repository)
				subscriptionServiceMock = new(mocks.SubscriptionReceiverService)
				logServiceMock          = new(mocks.LogService)
			)
			svc := subscription.NewService(repositoryMock, logServiceMock, nil, nil, subscriptionServiceMock)

			tc.Setup(repositoryMock, subscriptionServiceMock)

			_, err := svc.ListV2(ctx, subscription.Filter{})
			if tc.ErrString != "" {
				if tc.ErrString != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}

			repositoryMock.AssertExpectations(t)
			subscriptionServiceMock.AssertExpectations(t)
		})
	}
}

func TestService_GetV2(t *testing.T) {
	type testCase struct {
		Description string
		Setup       func(*mocks.Repository, *mocks.SubscriptionReceiverService)
		ErrString   string
	}
	var (
		ctx       = context.TODO()
		testCases = []testCase{
			{
				Description: "should return error if subscription get repository return error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(nil, errors.New("some error"))
				},
				ErrString: "some error",
			},
			{
				Description: "should return nil error if all passed",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&subscription.Subscription{}, nil)
					srs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.Filter")).Return([]subscriptionreceiver.Relation{}, nil)
				},
			},
			{
				Description: "should return error not found if subscription receiver service return error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().Get(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(&subscription.Subscription{}, nil)
					srs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.Filter")).Return(nil, errors.New("some error"))

				},
				ErrString: "some error",
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock          = new(mocks.Repository)
				subscriptionServiceMock = new(mocks.SubscriptionReceiverService)
				logServiceMock          = new(mocks.LogService)
			)
			svc := subscription.NewService(repositoryMock, logServiceMock, nil, nil, subscriptionServiceMock)

			tc.Setup(repositoryMock, subscriptionServiceMock)

			_, err := svc.GetV2(ctx, 100)
			if tc.ErrString != "" {
				if tc.ErrString != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}

			repositoryMock.AssertExpectations(t)
			subscriptionServiceMock.AssertExpectations(t)
		})
	}
}

func TestService_CreateV2(t *testing.T) {
	type testCase struct {
		Description  string
		Subscription *subscription.Subscription
		Setup        func(*mocks.Repository, *mocks.SubscriptionReceiverService, *mocks.ReceiverService)
		ErrString    string
	}

	var (
		ctx       = context.TODO()
		testCases = []testCase{
			{
				Description: "repository: should return error conflict if create subscription return error duplicate",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(subscription.ErrDuplicate)
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.ErrConflict.WithMsgf(subscription.ErrDuplicate.Error())).Return(nil)
				},
				ErrString: "urn already exist",
			},
			{
				Description: "repository: should return error not found if create subscription return error relation",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(subscription.ErrRelation)
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.ErrNotFound.WithMsgf(subscription.ErrRelation.Error())).Return(nil)
				},
				ErrString: "namespace id does not exist",
			},
			{
				Description: "repository: should return error if create subscription return some error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(nil)
				},
				ErrString: "some error",
			},
			{
				Description: "repository rollback error: subscription repository create return error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(errors.New("rollback error"))
				},
				ErrString: "rollback error",
			},
			{
				Description: "receiver service: return error if list receiver return error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return(nil, errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(nil)
				},
			},
			{
				Description: "receiver service rollback error: return error if error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return(nil, errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(errors.New("rollback error"))
				},
				ErrString: "rollback error",
			},
			{
				Description: "subscription receiver service: return error if a receiver does not exists",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
						{
							ID: 2,
						},
					}, nil)
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.ErrInvalid.WithMsgf("cannot found the receivers: [1]")).Return(nil)
				},
				ErrString: "cannot found the receivers: [1]",
			},
			{
				Description: "subscription receiver service: return error if bulk create return error duplicate",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
						{
							ID: 1,
						},
					}, nil)
					srs.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]subscriptionreceiver.Relation")).Return(subscriptionreceiver.ErrDuplicate)
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), subscriptionreceiver.ErrDuplicate).Return(nil)
				},
				ErrString: "subscription id and receiver id already registered",
			},
			{
				Description: "subscription receiver service: return error not found if bulk create return error relation",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
						{
							ID: 1,
						},
					}, nil)
					srs.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]subscriptionreceiver.Relation")).Return(subscriptionreceiver.ErrRelation)
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), subscriptionreceiver.ErrRelation).Return(nil)
				},
				ErrString: "subscription or receiver id does not exist",
			},
			{
				Description: "subscription receiver service: return error if bulk create return error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
						{
							ID: 1,
						},
					}, nil)
					srs.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]subscriptionreceiver.Relation")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(nil)
				},
				ErrString: "some error",
			},
			{
				Description: "should return no error if create subscription return no error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
						{
							ID: 1,
						},
					}, nil)
					srs.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]subscriptionreceiver.Relation")).Return(nil)
					sr.EXPECT().Commit(mock.AnythingOfType("context.todoCtx")).Return(nil)
				},
			},
			{
				Description:  "should return no error if create subscription without receiver return succeed",
				Subscription: &subscription.Subscription{},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					sr.EXPECT().Commit(mock.AnythingOfType("context.todoCtx")).Return(nil)
				},
			},
			{
				Description: "should return error if create subscription return no error but commit error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Create(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
						{
							ID: 1,
						},
					}, nil)
					srs.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]subscriptionreceiver.Relation")).Return(nil)
					sr.EXPECT().Commit(mock.AnythingOfType("context.todoCtx")).Return(errors.New("some error"))
				},
				ErrString: "some error",
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock          = new(mocks.Repository)
				subscriptionServiceMock = new(mocks.SubscriptionReceiverService)
				logServiceMock          = new(mocks.LogService)
				receiverServiceMock     = new(mocks.ReceiverService)
			)
			svc := subscription.NewService(
				repositoryMock,
				logServiceMock,
				nil,
				receiverServiceMock,
				subscriptionServiceMock,
			)
			tc.Setup(repositoryMock, subscriptionServiceMock, receiverServiceMock)

			err := svc.CreateV2(ctx, tc.Subscription)
			if tc.ErrString != "" {
				if tc.ErrString != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}

			repositoryMock.AssertExpectations(t)
			subscriptionServiceMock.AssertExpectations(t)
			receiverServiceMock.AssertExpectations(t)
		})
	}
}

func TestService_UpdateV2(t *testing.T) {
	type testCase struct {
		Description  string
		Subscription *subscription.Subscription
		Setup        func(*mocks.Repository, *mocks.SubscriptionReceiverService, *mocks.ReceiverService)
		ErrString    string
	}

	var (
		ctx       = context.TODO()
		testCases = []testCase{
			{
				Description: "repository: should return error conflict if update subscription return error duplicate",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(subscription.ErrDuplicate)
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.ErrConflict.WithMsgf(subscription.ErrDuplicate.Error())).Return(nil)
				},
				ErrString: "urn already exist",
			},
			{
				Description: "repository: should return error not found if update subscription return error relation",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(subscription.ErrRelation)
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.ErrNotFound.WithMsgf(subscription.ErrRelation.Error())).Return(nil)
				},
				ErrString: "namespace id does not exist",
			},
			{
				Description: "repository: should return error if update subscription return some error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(nil)
				},
				ErrString: "some error",
			},
			{
				Description: "repository rollback error: subscription repository update return error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(errors.New("rollback error"))
				},
				ErrString: "rollback error",
			},
			{
				Description: "receiver service: return error if list receiver return error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return(nil, errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(nil)
				},
			},
			{
				Description: "receiver service rollback error: return error if error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return(nil, errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(errors.New("rollback error"))
				},
				ErrString: "rollback error",
			},
			{
				Description: "subscription receiver service: return error if bulk upsert return error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
						{
							ID: 1,
						},
					}, nil)
					srs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.Filter")).Return([]subscriptionreceiver.Relation{}, nil)
					srs.EXPECT().BulkUpsert(mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("[]subscriptionreceiver.Relation")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("*context.cancelCtx"), errors.New("some error")).Return(nil)
				},
				ErrString: "some error",
			},
			{
				Description: "subscription receiver service: return error if receiver not found",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{}, nil)
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.ErrInvalid.WithMsgf("cannot found the receivers: [1]")).Return(nil)
				},
				ErrString: "cannot found the receivers: [1]",
			},
			{
				Description: "subscription receiver service: return error if bulk soft delete return error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
						{
							ID: 1,
						},
					}, nil)
					srs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.Filter")).Return([]subscriptionreceiver.Relation{
						{
							ReceiverID: 2,
						},
					}, nil)
					srs.EXPECT().BulkUpsert(mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("[]subscriptionreceiver.Relation")).Return(nil)
					srs.EXPECT().BulkSoftDelete(mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("subscriptionreceiver.DeleteFilter")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("*context.cancelCtx"), errors.New("some error")).Return(nil)
				},
				ErrString: "some error",
			},
			{
				Description: "should return no error if update subscription return no error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
						{
							ID: 1,
						},
					}, nil)
					srs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.Filter")).Return([]subscriptionreceiver.Relation{
						{
							ReceiverID: 2,
						},
					}, nil)
					srs.EXPECT().BulkUpsert(mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("[]subscriptionreceiver.Relation")).Return(nil)
					srs.EXPECT().BulkSoftDelete(mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("subscriptionreceiver.DeleteFilter")).Return(nil)
					sr.EXPECT().Commit(mock.AnythingOfType("*context.cancelCtx")).Return(nil)
				},
			},
			{
				Description:  "should return no error if update subscription without receiver succeed",
				Subscription: &subscription.Subscription{},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					sr.EXPECT().Commit(mock.AnythingOfType("context.todoCtx")).Return(nil)
				},
			},
			{
				Description: "should return error if update subscription return no error but commit error",
				Subscription: &subscription.Subscription{
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					sr.EXPECT().Update(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("*subscription.Subscription")).Return(nil)
					rs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("receiver.Filter")).Return([]receiver.Receiver{
						{
							ID: 1,
						},
					}, nil)
					srs.EXPECT().List(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.Filter")).Return([]subscriptionreceiver.Relation{
						{
							ReceiverID: 2,
						},
					}, nil)
					srs.EXPECT().BulkUpsert(mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("[]subscriptionreceiver.Relation")).Return(nil)
					srs.EXPECT().BulkSoftDelete(mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("subscriptionreceiver.DeleteFilter")).Return(nil)
					sr.EXPECT().Commit(mock.AnythingOfType("*context.cancelCtx")).Return(errors.New("some error"))
				},
				ErrString: "some error",
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock          = new(mocks.Repository)
				subscriptionServiceMock = new(mocks.SubscriptionReceiverService)
				logServiceMock          = new(mocks.LogService)
				receiverServiceMock     = new(mocks.ReceiverService)
			)
			svc := subscription.NewService(
				repositoryMock,
				logServiceMock,
				nil,
				receiverServiceMock,
				subscriptionServiceMock,
			)
			tc.Setup(repositoryMock, subscriptionServiceMock, receiverServiceMock)

			err := svc.UpdateV2(ctx, tc.Subscription)
			if tc.ErrString != "" {
				if tc.ErrString != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}

			repositoryMock.AssertExpectations(t)
			subscriptionServiceMock.AssertExpectations(t)
			receiverServiceMock.AssertExpectations(t)
		})
	}
}

func TestService_DeleteV2(t *testing.T) {
	type testCase struct {
		Description  string
		Subscription *subscription.Subscription
		Setup        func(*mocks.Repository, *mocks.SubscriptionReceiverService)
		ErrString    string
	}

	var (
		ctx       = context.TODO()
		testCases = []testCase{
			{
				Description: "repository: should return error if bulk soft delete return some error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					srs.EXPECT().BulkDelete(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.DeleteFilter")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(nil)
				},
				ErrString: "some error",
			},

			{
				Description: "repository rollback error: should return error if bulk soft delete return some error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					srs.EXPECT().BulkDelete(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.DeleteFilter")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(errors.New("some error rollback"))
				},
				ErrString: "some error rollback",
			},

			{
				Description: "subscription receiver service: return error if delete subscription return error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					srs.EXPECT().BulkDelete(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.DeleteFilter")).Return(nil)
					sr.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(nil)
				},
				ErrString: "some error",
			},
			{
				Description: "subscription receiver service rollback error: return error if delete subscription return error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					srs.EXPECT().BulkDelete(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.DeleteFilter")).Return(nil)
					sr.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(errors.New("some error"))
					sr.EXPECT().Rollback(mock.AnythingOfType("context.todoCtx"), errors.New("some error")).Return(errors.New("some error rollback"))
				},
				ErrString: "some error rollback",
			},
			{
				Description: "should return no error if delete subscription return no error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					srs.EXPECT().BulkDelete(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.DeleteFilter")).Return(nil)
					sr.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(nil)
					sr.EXPECT().Commit(mock.AnythingOfType("context.todoCtx")).Return(nil)
				},
			},
			{
				Description: "should return error if delete subscription return no error but commit error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService) {
					sr.EXPECT().WithTransaction(mock.AnythingOfType("context.todoCtx")).Return(context.TODO())
					srs.EXPECT().BulkDelete(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscriptionreceiver.DeleteFilter")).Return(nil)
					sr.EXPECT().Delete(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("uint64")).Return(nil)
					sr.EXPECT().Commit(mock.AnythingOfType("context.todoCtx")).Return(errors.New("some error"))
				},
				ErrString: "some error",
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock          = new(mocks.Repository)
				subscriptionServiceMock = new(mocks.SubscriptionReceiverService)
				logServiceMock          = new(mocks.LogService)
			)
			svc := subscription.NewService(
				repositoryMock,
				logServiceMock,
				nil,
				nil,
				subscriptionServiceMock,
			)
			tc.Setup(repositoryMock, subscriptionServiceMock)

			err := svc.DeleteV2(ctx, 1)
			if tc.ErrString != "" {
				if tc.ErrString != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}

			repositoryMock.AssertExpectations(t)
			subscriptionServiceMock.AssertExpectations(t)
		})
	}
}

func TestService_MatchByLabelsV2(t *testing.T) {
	type testCase struct {
		Description  string
		Subscription *subscription.Subscription
		Setup        func(*mocks.Repository, *mocks.SubscriptionReceiverService, *mocks.ReceiverService)
		ErrString    string
	}

	var (
		ctx       = context.TODO()
		testCases = []testCase{
			{
				Description: "should return error if match labels subscription return some error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().MatchLabelsFetchReceivers(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscription.Filter")).Return(nil, errors.New("some error"))
				},
				ErrString: "some error",
			},
			{
				Description: "should return no error if match labels subscription return no error but post db hook receiver return error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().MatchLabelsFetchReceivers(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscription.Filter")).Return([]subscription.ReceiverView{
						{
							ID: 1,
							Configurations: map[string]any{
								"token": "key",
							},
						},
						{
							ID: 2,
							Configurations: map[string]any{
								"token": "key",
							},
						},
					}, nil)
					rs.EXPECT().PostHookDBTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("string"), map[string]any{
						"token": "key",
					}).Times(1).Return(nil, errors.New("decrypt error"))
				},
			},
			{
				Description: "should return no error if match labels subscription return no error",
				Setup: func(sr *mocks.Repository, srs *mocks.SubscriptionReceiverService, rs *mocks.ReceiverService) {
					sr.EXPECT().MatchLabelsFetchReceivers(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("subscription.Filter")).Return([]subscription.ReceiverView{
						{
							ID: 1,
							Configurations: map[string]any{
								"token": "key",
							},
						},
						{
							ID: 2,
							Configurations: map[string]any{
								"token": "key",
							},
						},
					}, nil)
					rs.EXPECT().PostHookDBTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("string"), map[string]any{
						"token": "key",
					}).Times(2).Return(map[string]any{
						"token": "key",
					}, nil)
				},
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock          = new(mocks.Repository)
				subscriptionServiceMock = new(mocks.SubscriptionReceiverService)
				receiverServiceMock     = new(mocks.ReceiverService)
				logServiceMock          = new(mocks.LogService)
			)
			svc := subscription.NewService(
				repositoryMock,
				logServiceMock,
				nil,
				receiverServiceMock,
				subscriptionServiceMock,
			)
			tc.Setup(repositoryMock, subscriptionServiceMock, receiverServiceMock)

			_, err := svc.MatchByLabelsV2(ctx, 1, map[string]string{})
			if tc.ErrString != "" {
				if tc.ErrString != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}

			repositoryMock.AssertExpectations(t)
			subscriptionServiceMock.AssertExpectations(t)
			receiverServiceMock.AssertExpectations(t)
		})
	}
}

func TestClassifyReceivers(t *testing.T) {

	type testCase struct {
		Description       string
		NewRelations      []subscriptionreceiver.Relation
		ExistingRelations []subscriptionreceiver.Relation
		ExpectedToUpsert  []subscriptionreceiver.Relation
		ExpectedToDelete  []subscriptionreceiver.Relation
	}

	var testCases = []testCase{
		{
			Description: "should correctly distribute receivers of a subscription to upsert and to delete",
			NewRelations: []subscriptionreceiver.Relation{
				{
					ReceiverID: 1,
				},
				{
					ReceiverID: 2,
				},
				{
					ReceiverID: 3,
				},
			},
			ExistingRelations: []subscriptionreceiver.Relation{
				{
					ReceiverID: 2,
				},
				{
					ReceiverID: 3,
				},
				{
					ReceiverID: 4,
				},
			},
			ExpectedToUpsert: []subscriptionreceiver.Relation{
				{
					ReceiverID: 1,
				},
				{
					ReceiverID: 2,
				},
				{
					ReceiverID: 3,
				},
			},
			ExpectedToDelete: []subscriptionreceiver.Relation{
				{
					ReceiverID: 4,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			tu, td := subscription.ClassifyReceivers(testCase.NewRelations, testCase.ExistingRelations)
			if diff := cmp.Diff(tu, testCase.ExpectedToUpsert,
				cmpopts.IgnoreFields(subscriptionreceiver.Relation{}, "ID")); diff != "" {
				t.Fatalf("got diff to update %+v", diff)
			}
			if diff := cmp.Diff(td, testCase.ExpectedToDelete,
				cmpopts.IgnoreFields(subscriptionreceiver.Relation{}, "ID")); diff != "" {
				t.Fatalf("got diff to delete %+v", diff)
			}
		})
	}
}
