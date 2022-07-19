package postgres_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/odpf/salt/log"
	"github.com/odpf/siren/core/subscription"
	"github.com/odpf/siren/internal/store/postgres"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
)

type SubscriptionRepositoryTestSuite struct {
	suite.Suite
	ctx        context.Context
	client     *postgres.Client
	pool       *dockertest.Pool
	resource   *dockertest.Resource
	repository *postgres.SubscriptionRepository
}

func (s *SubscriptionRepositoryTestSuite) SetupSuite() {
	var err error

	logger := log.NewZap()
	s.client, s.pool, s.resource, err = newTestClient(logger)
	if err != nil {
		s.T().Fatal(err)
	}

	s.ctx = context.TODO()
	s.repository = postgres.NewSubscriptionRepository(s.client)

	_, err = bootstrapProvider(s.client)
	if err != nil {
		s.T().Fatal(err)
	}

	_, err = bootstrapNamespace(s.client)
	if err != nil {
		s.T().Fatal(err)
	}

	_, err = bootstrapReceiver(s.client)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionRepositoryTestSuite) SetupTest() {
	var err error
	_, err = bootstrapSubscription(s.client)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionRepositoryTestSuite) TearDownSuite() {
	// Clean tests
	if err := purgeDocker(s.pool, s.resource); err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionRepositoryTestSuite) TearDownTest() {
	if err := s.cleanup(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionRepositoryTestSuite) cleanup() error {
	queries := []string{
		"TRUNCATE TABLE subscriptions RESTART IDENTITY CASCADE",
	}
	return execQueries(context.TODO(), s.client, queries)
}

func (s *SubscriptionRepositoryTestSuite) TestList() {
	type testCase struct {
		Description           string
		Filter                subscription.Filter
		ExpectedSubscriptions []subscription.Subscription
		ErrString             string
	}

	var testCases = []testCase{
		{
			Description: "should get all subscriptions",
			ExpectedSubscriptions: []subscription.Subscription{
				{
					ID:        1,
					URN:       "alert-history-odpf",
					Namespace: 2,
					Receivers: []subscription.Receiver{
						{
							ID: 1,
						},
					},
				},
				{
					ID:        2,
					URN:       "odpf-data-warning",
					Namespace: 1,
					Receivers: []subscription.Receiver{
						{
							ID: 1,
							Configuration: map[string]string{
								"channel_name": "odpf-data",
							},
						},
					},
					Match: map[string]string{
						"environment": "integration",
						"team":        "odpf-data",
					},
				},
				{
					ID:        3,
					URN:       "odpf-pd",
					Namespace: 2,
					Receivers: []subscription.Receiver{
						{
							ID: 31,
						},
					},
					Match: map[string]string{
						"environment": "production",
						"severity":    "CRITICAL",
						"team":        "odpf-app",
					},
				},
			},
		},
		{
			Description: "should get filtered subscriptions",
			Filter: subscription.Filter{
				NamespaceID: 1,
			},
			ExpectedSubscriptions: []subscription.Subscription{
				{
					ID:        2,
					URN:       "odpf-data-warning",
					Namespace: 1,
					Receivers: []subscription.Receiver{
						{
							ID: 1,
							Configuration: map[string]string{
								"channel_name": "odpf-data",
							},
						},
					},
					Match: map[string]string{
						"environment": "integration",
						"team":        "odpf-data",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			got, err := s.repository.List(s.ctx, tc.Filter)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
			if !cmp.Equal(got, tc.ExpectedSubscriptions, cmpopts.IgnoreFields(subscription.Subscription{}, "CreatedAt", "UpdatedAt")) {
				s.T().Fatalf("got result %+v, expected was %+v", got, tc.ExpectedSubscriptions)
			}
		})
	}
}

func (s *SubscriptionRepositoryTestSuite) TestCreate() {
	type testCase struct {
		Description          string
		SubscriptionToUpsert *subscription.Subscription
		ExpectedID           uint64
		ErrString            string
	}

	var testCases = []testCase{
		{
			Description: "should create a subscription if doesn't exist",
			SubscriptionToUpsert: &subscription.Subscription{
				Namespace: 1,
				URN:       "foo",
				Match: map[string]string{
					"foo": "bar",
				},
				Receivers: []subscription.Receiver{
					{
						ID:            2,
						Configuration: map[string]string{},
					},
					{
						ID: 1,
						Configuration: map[string]string{
							"channel_name": "test",
						},
					},
				},
			},
			ExpectedID: uint64(4), // autoincrement in db side
		},
		{
			Description: "should return duplicate error if urn already exist",
			SubscriptionToUpsert: &subscription.Subscription{
				Namespace: 1,
				URN:       "foo",
				Match: map[string]string{
					"foo": "bar",
				},
				Receivers: []subscription.Receiver{
					{
						ID:            2,
						Configuration: map[string]string{},
					},
					{
						ID: 1,
						Configuration: map[string]string{
							"channel_name": "test",
						},
					},
				},
			},
			ErrString: "urn already exist",
		}, {
			Description: "should return relation error if namespace id does not exist",
			SubscriptionToUpsert: &subscription.Subscription{
				Namespace: 1000,
				URN:       "new-foo",
				Match: map[string]string{
					"foo": "bar",
				},
				Receivers: []subscription.Receiver{
					{
						ID:            2,
						Configuration: map[string]string{},
					},
					{
						ID: 1,
						Configuration: map[string]string{
							"channel_name": "test",
						},
					},
				},
			},
			ErrString: "namespace id does not exist",
		},
		{
			Description: "should return error if subscription is nil",
			ErrString:   "subscription domain is nil",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.Create(s.ctx, tc.SubscriptionToUpsert)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
		})
	}
}

func (s *SubscriptionRepositoryTestSuite) TestGet() {
	type testCase struct {
		Description          string
		PassedID             uint64
		ExpectedSubscription *subscription.Subscription
		ErrString            string
	}

	var testCases = []testCase{
		{
			Description: "should get a subscription",
			PassedID:    uint64(3),
			ExpectedSubscription: &subscription.Subscription{
				ID:        3,
				URN:       "odpf-pd",
				Namespace: 2,
				Receivers: []subscription.Receiver{
					{
						ID: 31,
					},
				},
				Match: map[string]string{
					"environment": "production",
					"severity":    "CRITICAL",
					"team":        "odpf-app",
				},
			},
		},
		{
			Description: "should return not found if id not found",
			PassedID:    uint64(1000),
			ErrString:   "subscription with id 1000 not found",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			got, err := s.repository.Get(s.ctx, tc.PassedID)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
			if !cmp.Equal(got, tc.ExpectedSubscription, cmpopts.IgnoreFields(subscription.Subscription{}, "CreatedAt", "UpdatedAt")) {
				s.T().Fatalf("got result %+v, expected was %+v", got, tc.ExpectedSubscription)
			}
		})
	}
}

func (s *SubscriptionRepositoryTestSuite) TestUpdate() {
	type testCase struct {
		Description          string
		SubscriptionToUpsert *subscription.Subscription
		ExpectedID           uint64
		ErrString            string
	}

	var testCases = []testCase{
		{
			Description: "should update a subscription",
			SubscriptionToUpsert: &subscription.Subscription{
				ID:        3,
				URN:       "odpf-pd",
				Namespace: 2,
				Receivers: []subscription.Receiver{
					{
						ID: 3100,
					},
				},
				Match: map[string]string{
					"key": "label",
				},
			},
			ExpectedID: uint64(3),
		},
		{
			Description: "should return duplicate error if urn already exist",
			SubscriptionToUpsert: &subscription.Subscription{
				ID:        1,
				URN:       "odpf-pd",
				Namespace: 2,
				Receivers: []subscription.Receiver{
					{
						ID: 3100,
					},
				},
				Match: map[string]string{
					"key": "label",
				},
			},
			ErrString: "urn already exist",
		},
		{
			Description: "should return relation error if namespace id does not exist",
			SubscriptionToUpsert: &subscription.Subscription{
				ID:        3,
				URN:       "odpf-pd",
				Namespace: 1000,
				Receivers: []subscription.Receiver{
					{
						ID: 3100,
					},
				},
				Match: map[string]string{
					"key": "label",
				},
			},
			ErrString: "namespace id does not exist",
		},
		{
			Description: "should return not found error if id not found",
			SubscriptionToUpsert: &subscription.Subscription{
				ID:        3000,
				URN:       "odpf-pd",
				Namespace: 1,
				Receivers: []subscription.Receiver{
					{
						ID: 3100,
					},
				},
				Match: map[string]string{
					"key": "label",
				},
			},
			ErrString: "subscription with id 3000 not found",
		},
		{
			Description: "should return error if subscription is nil",
			ErrString:   "subscription domain is nil",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.Update(s.ctx, tc.SubscriptionToUpsert)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
		})
	}
}

func (s *SubscriptionRepositoryTestSuite) TestDelete() {
	type testCase struct {
		Description string
		IDToDelete  uint64
		ErrString   string
	}

	var testCases = []testCase{
		{
			Description: "should delete a subscription",
			IDToDelete:  1,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.Delete(s.ctx, tc.IDToDelete, 1)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
		})
	}
}

func (s *SubscriptionRepositoryTestSuite) TestTransaction() {
	s.Run("successfully commit transaction", func() {
		ctx := s.repository.WithTransaction(context.Background())
		err := s.repository.Create(ctx, &subscription.Subscription{
			Namespace: 2,
			URN:       "foo-commit",
			Match: map[string]string{
				"foo": "bar",
			},
			Receivers: []subscription.Receiver{
				{
					ID:            2,
					Configuration: map[string]string{},
				},
				{
					ID: 1,
					Configuration: map[string]string{
						"channel_name": "test",
					},
				},
			},
		})
		s.NoError(err)

		err = s.repository.Commit(ctx)
		s.NoError(err)

		fetchedRules, err := s.repository.List(s.ctx, subscription.Filter{
			NamespaceID: 2,
		})
		s.NoError(err)
		s.Len(fetchedRules, 3)
	})

	s.Run("successfully rollback transaction", func() {
		ctx := s.repository.WithTransaction(context.Background())
		err := s.repository.Create(ctx, &subscription.Subscription{
			Namespace: 2,
			URN:       "foo-rollback",
			Match: map[string]string{
				"foo": "bar",
			},
			Receivers: []subscription.Receiver{
				{
					ID:            2,
					Configuration: map[string]string{},
				},
				{
					ID: 1,
					Configuration: map[string]string{
						"channel_name": "test",
					},
				},
			},
		})
		s.NoError(err)

		err = s.repository.Rollback(ctx, nil)
		s.NoError(err)

		fetchedRules, err := s.repository.List(s.ctx, subscription.Filter{
			NamespaceID: 2,
		})
		s.NoError(err)
		s.Len(fetchedRules, 3)
	})
}

func TestSubscriptionRepository(t *testing.T) {
	suite.Run(t, new(SubscriptionRepositoryTestSuite))
}
