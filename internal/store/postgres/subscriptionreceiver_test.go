package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/goto/salt/dockertestx"
	"github.com/goto/salt/log"
	"github.com/goto/siren/core/subscriptionreceiver"
	"github.com/goto/siren/internal/store/postgres"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/pgc"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
)

type SubscriptionReceiverRepositoryTestSuite struct {
	suite.Suite
	ctx        context.Context
	client     *pgc.Client
	pool       *dockertest.Pool
	resource   *dockertest.Resource
	repository *postgres.SubscriptionReceiverRepository
}

func (s *SubscriptionReceiverRepositoryTestSuite) SetupSuite() {
	var err error

	logger := log.NewZap()
	dpg, err := dockertestx.CreatePostgres(
		dockertestx.PostgresWithDetail(
			pgUser, pgPass, pgDBName,
		),
		dockertestx.PostgresWithVersionTag("13"),
	)
	if err != nil {
		s.T().Fatal(err)
	}

	s.pool = dpg.GetPool()
	s.resource = dpg.GetResource()

	dbConfig.URL = dpg.GetExternalConnString()
	s.client, err = pgc.NewClient(logger, dbConfig)
	if err != nil {
		s.T().Fatal(err)
	}

	s.ctx = context.TODO()
	s.Require().NoError(migrate(s.ctx, s.client, dbConfig))

	s.repository = postgres.NewSubscriptionReceiverRepository(s.client)

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

	_, err = bootstrapSubscription(s.client)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionReceiverRepositoryTestSuite) SetupTest() {
	if _, err := bootstrapSubscriptionReceiver(s.client); err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionReceiverRepositoryTestSuite) TearDownSuite() {
	// Clean tests
	if err := purgeDocker(s.pool, s.resource); err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionReceiverRepositoryTestSuite) TearDownTest() {
	if err := s.cleanup(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionReceiverRepositoryTestSuite) cleanup() error {
	queries := []string{
		"TRUNCATE TABLE subscriptions_receivers RESTART IDENTITY CASCADE",
	}
	return execQueries(context.TODO(), s.client, queries)
}

func (s *SubscriptionReceiverRepositoryTestSuite) TestList() {
	type testCase struct {
		Description string
		Filter      subscriptionreceiver.Filter
		Expected    []subscriptionreceiver.Relation
		ErrString   string
	}

	var testCases = []testCase{
		{
			Description: "should get all subscriptions receivers id 1 and 2",
			Filter: subscriptionreceiver.Filter{
				SubscriptionIDs: []uint64{1, 2},
			},
			Expected: []subscriptionreceiver.Relation{
				{
					SubscriptionID: 1,
					ReceiverID:     1,
					Labels:         map[string]string{},
				},
				{
					SubscriptionID: 1,
					ReceiverID:     2,
					Labels:         map[string]string{},
				},
				{
					SubscriptionID: 1,
					ReceiverID:     3,
					Labels: map[string]string{
						"lk1": "lv1",
						"lk2": "lv2",
					},
				},
				{
					SubscriptionID: 2,
					ReceiverID:     3,
					Labels:         map[string]string{},
				},
			},
		},
		{
			Description: "should get all subscriptions receivers id 2 with receiver id 3",
			Filter: subscriptionreceiver.Filter{
				SubscriptionIDs: []uint64{2},
				ReceiverID:      3,
			},
			Expected: []subscriptionreceiver.Relation{
				{
					SubscriptionID: 2,
					ReceiverID:     3,
					Labels:         map[string]string{},
				},
			},
		},
		{
			Description: "should get all subscriptions receivers with filter label",
			Filter: subscriptionreceiver.Filter{
				SubscriptionIDs: []uint64{1},
				Labels: map[string]string{
					"lk1": "lv1",
				},
			},
			Expected: []subscriptionreceiver.Relation{
				{
					SubscriptionID: 1,
					ReceiverID:     3,
					Labels: map[string]string{
						"lk1": "lv1",
						"lk2": "lv2",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			got, err := s.repository.List(s.ctx, tc.Filter)
			if err != nil && err.Error() != tc.ErrString {
				s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
			}
			if diff := cmp.Diff(got, tc.Expected, cmpopts.IgnoreFields(subscriptionreceiver.Relation{}, "ID", "CreatedAt", "UpdatedAt")); diff != "" {
				s.T().Fatalf("got diff %+v", diff)
			}
		})
	}
}

func (s *SubscriptionReceiverRepositoryTestSuite) TestBulkCreate() {
	type testCase struct {
		Description    string
		ToUpsert       []subscriptionreceiver.Relation
		TesterFunction func(t *testing.T, tc testCase, r *postgres.SubscriptionReceiverRepository)
		ErrString      string
	}

	var testCases = []testCase{
		{
			Description: "receiver 2 and 3 should subscribe to subscription 2",
			ToUpsert: []subscriptionreceiver.Relation{
				{
					SubscriptionID: 2,
					ReceiverID:     2,
					Labels:         map[string]string{},
				},
				{
					SubscriptionID: 2,
					ReceiverID:     4,
					Labels:         map[string]string{},
				},
			},
			TesterFunction: func(t *testing.T, tc testCase, r *postgres.SubscriptionReceiverRepository) {
				sr, err := r.List(context.Background(), subscriptionreceiver.Filter{
					SubscriptionIDs: []uint64{2},
				})
				if err != nil {
					t.Fatal(err)
				}
				expectedRelations := []subscriptionreceiver.Relation{
					{
						SubscriptionID: 2,
						ReceiverID:     2,
						Labels:         map[string]string{},
					},
					{
						SubscriptionID: 2,
						ReceiverID:     3,
						Labels:         map[string]string{},
					},
					{
						SubscriptionID: 2,
						ReceiverID:     4,
						Labels:         map[string]string{},
					},
				}
				if diff := cmp.Diff(sr, expectedRelations, cmpopts.IgnoreFields(subscriptionreceiver.Relation{}, "ID", "CreatedAt", "UpdatedAt")); diff != "" {
					t.Error(diff)
				}
			},
		},
		{
			Description: "should return error if subscription not exist",
			ToUpsert: []subscriptionreceiver.Relation{
				{
					SubscriptionID: 999,
					ReceiverID:     2,
				},
			},
			ErrString:      "subscription or receiver id does not exist",
			TesterFunction: func(t *testing.T, tc testCase, r *postgres.SubscriptionReceiverRepository) {},
		},
		{
			Description: "should return error if receiver not exist",
			ToUpsert: []subscriptionreceiver.Relation{
				{
					SubscriptionID: 2,
					ReceiverID:     999,
				},
			},
			ErrString: "subscription or receiver id does not exist",

			TesterFunction: func(t *testing.T, tc testCase, r *postgres.SubscriptionReceiverRepository) {},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.BulkCreate(s.ctx, tc.ToUpsert)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
			if tc.TesterFunction != nil {
				tc.TesterFunction(s.T(), tc, s.repository)
			}
		})
	}
}

func (s *SubscriptionReceiverRepositoryTestSuite) TestBulkUpsert() {
	type testCase struct {
		Description    string
		TesterFunction func(t *testing.T, tc testCase, r *postgres.SubscriptionReceiverRepository)
		ToUpsert       []subscriptionreceiver.Relation
		ErrString      string
	}

	var testCases = []testCase{
		{
			Description: "subscription 2 should be created and subscription 1 should be updated",
			ToUpsert: []subscriptionreceiver.Relation{
				{
					SubscriptionID: 2,
					ReceiverID:     1,
					Labels:         map[string]string{},
				},
				{
					SubscriptionID: 1,
					ReceiverID:     1,
					Labels:         map[string]string{},
				},
			},
			TesterFunction: func(t *testing.T, tc testCase, r *postgres.SubscriptionReceiverRepository) {
				sr2, err := r.List(context.Background(), subscriptionreceiver.Filter{
					SubscriptionIDs: []uint64{2},
				})
				if err != nil {
					t.Fatal(err)
				}

				if diff := cmp.Diff(sr2, []subscriptionreceiver.Relation{
					{
						SubscriptionID: 2,
						ReceiverID:     1,
						Labels:         map[string]string{},
					},
					{
						SubscriptionID: 2,
						ReceiverID:     3,
						Labels:         map[string]string{},
					},
				}, cmpopts.IgnoreFields(subscriptionreceiver.Relation{}, "ID", "CreatedAt", "UpdatedAt")); diff != "" {
					t.Error(diff)
				}

				sr1, err := r.List(context.Background(), subscriptionreceiver.Filter{
					SubscriptionIDs: []uint64{1},
				})
				if err != nil {
					t.Fatal(err)
				}
				var expectedSR1 = []subscriptionreceiver.Relation{
					{
						SubscriptionID: 1,
						ReceiverID:     1,
						Labels:         map[string]string{},
					},
					{
						SubscriptionID: 1,
						ReceiverID:     2,
						Labels:         map[string]string{},
					},
					{
						SubscriptionID: 1,
						ReceiverID:     3,
						Labels: map[string]string{
							"lk1": "lv1",
							"lk2": "lv2",
						},
					},
				}
				if diff := cmp.Diff(sr1, expectedSR1,
					cmpopts.IgnoreFields(subscriptionreceiver.Relation{}, "ID", "CreatedAt", "UpdatedAt"),
				); diff != "" {
					t.Error(diff)
				}
			},
		},
		{
			Description: "should return error if subscription not exist",
			ToUpsert: []subscriptionreceiver.Relation{
				{
					SubscriptionID: 999,
					ReceiverID:     2,
				},
			},
			ErrString: "subscription or receiver id does not exist",
		},
		{
			Description: "should return error if receiver not exist",
			ToUpsert: []subscriptionreceiver.Relation{
				{
					SubscriptionID: 1,
					ReceiverID:     999,
				},
			},
			ErrString: "subscription or receiver id does not exist",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.BulkUpsert(s.ctx, tc.ToUpsert)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
			if tc.TesterFunction != nil {
				tc.TesterFunction(s.T(), tc, s.repository)
			}
		})
	}
}

func (s *SubscriptionReceiverRepositoryTestSuite) TestUpdate() {
	type testCase struct {
		Description string
		ToUpdate    *subscriptionreceiver.Relation
		ErrString   string
	}

	var testCases = []testCase{
		{
			Description: "should update a subscription receiver relation",
			ToUpdate: &subscriptionreceiver.Relation{
				SubscriptionID: 1,
				ReceiverID:     2,
				Labels: map[string]string{
					"newkey": "newvalue",
				},
			},
		},
		{
			Description: "should return relation error if subscription id does not exist",
			ToUpdate: &subscriptionreceiver.Relation{
				SubscriptionID: 100,
				ReceiverID:     2,
				Labels: map[string]string{
					"newkey": "newvalue",
				},
			},
			ErrString: "subscription with id 100 and receiver with id 2 not found",
		},
		{
			Description: "should return relation error if receiver id does not exist",
			ToUpdate: &subscriptionreceiver.Relation{
				SubscriptionID: 1,
				ReceiverID:     200,
				Labels: map[string]string{
					"newkey": "newvalue",
				},
			},
			ErrString: "subscription with id 1 and receiver with id 200 not found",
		},
		{
			Description: "should return error if subscription receiver relation is nil",
			ErrString:   "subscription receiver relation is nil",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.Update(s.ctx, tc.ToUpdate)
			if err != nil && err.Error() != tc.ErrString {
				s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
			}
		})
	}
}

func (s *SubscriptionReceiverRepositoryTestSuite) TestBulkSoftDelete() {
	type testCase struct {
		Description    string
		Filter         subscriptionreceiver.DeleteFilter
		TesterFunction func(t *testing.T, tc testCase, r *postgres.SubscriptionReceiverRepository)
		ErrString      string
	}

	var testCases = []testCase{
		{
			Description: "deleting subscription by id should works",
			Filter: subscriptionreceiver.DeleteFilter{
				SubscriptionID: 2,
			},
			TesterFunction: func(t *testing.T, tc testCase, r *postgres.SubscriptionReceiverRepository) {
				sr2, err := r.List(context.Background(), subscriptionreceiver.Filter{
					SubscriptionIDs: []uint64{2},
					Deleted:         true,
				})
				if err != nil {
					t.Fatal(err)
				}

				s.Assert().NotEqual(sr2[0].DeletedAt, time.Time{})
			},
		},
		{
			Description: "deleting subscription by pair should works",
			Filter: subscriptionreceiver.DeleteFilter{
				Pair: []subscriptionreceiver.Relation{
					{
						SubscriptionID: 1,
						ReceiverID:     3,
					},
					{
						SubscriptionID: 1,
						ReceiverID:     2,
					},
				},
			},
			TesterFunction: func(t *testing.T, tc testCase, r *postgres.SubscriptionReceiverRepository) {
				sr1, err := r.List(context.Background(), subscriptionreceiver.Filter{
					SubscriptionIDs: []uint64{1},
					Deleted:         true,
				})
				if err != nil {
					t.Fatal(err)
				}

				for _, sr := range sr1 {
					if sr.ReceiverID == 3 || sr.ReceiverID == 2 {
						s.Assert().NotEqual(sr.DeletedAt, time.Time{})
					} else {
						s.Assert().Equal(sr.DeletedAt, time.Time{})
					}
				}
			},
		},
		{
			Description: "should return error if both filters are passed",
			Filter: subscriptionreceiver.DeleteFilter{
				Pair: []subscriptionreceiver.Relation{
					{
						SubscriptionID: 1,
						ReceiverID:     3,
					},
				},
				SubscriptionID: 2,
			},
			ErrString: "use either pairs of subscription id and receiver id or a single subscription id",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.BulkSoftDelete(s.ctx, tc.Filter)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
			if tc.TesterFunction != nil {
				tc.TesterFunction(s.T(), tc, s.repository)
			}
		})
	}
}

func (s *SubscriptionReceiverRepositoryTestSuite) TestBulkDelete() {
	type testCase struct {
		Description    string
		Filter         subscriptionreceiver.DeleteFilter
		TesterFunction func(t *testing.T, tc testCase, r *postgres.SubscriptionReceiverRepository)
		Err            error
	}

	var testCases = []testCase{
		{
			Description: "deleting subscription by id should works",
			Filter: subscriptionreceiver.DeleteFilter{
				SubscriptionID: 2,
			},
			Err: nil,
		},
		{
			Description: "deleting subscription by pair should works",
			Filter: subscriptionreceiver.DeleteFilter{
				Pair: []subscriptionreceiver.Relation{
					{
						SubscriptionID: 1,
						ReceiverID:     3,
					},
					{
						SubscriptionID: 1,
						ReceiverID:     2,
					},
				},
			},
			Err: nil,
		},
		{
			Description: "should return error if both filters are passed",
			Filter: subscriptionreceiver.DeleteFilter{
				Pair: []subscriptionreceiver.Relation{
					{
						SubscriptionID: 1,
						ReceiverID:     3,
					},
				},
				SubscriptionID: 2,
			},
			Err: errors.New("use either pairs of subscription id and receiver id or a single subscription id"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.BulkDelete(s.ctx, tc.Filter)
			s.Assert().Equal(err, tc.Err)
		})
	}
}

func TestSubscriptionReceiverRepository(t *testing.T) {
	suite.Run(t, new(SubscriptionReceiverRepositoryTestSuite))
}
