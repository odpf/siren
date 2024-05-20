package postgres_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/goto/salt/dockertestx"
	"github.com/goto/salt/log"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/internal/store/postgres"
	"github.com/goto/siren/pkg/pgc"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
)

type SubscriptionV2RepositoryTestSuite struct {
	suite.Suite
	ctx        context.Context
	client     *pgc.Client
	pool       *dockertest.Pool
	resource   *dockertest.Resource
	repository *postgres.SubscriptionRepository
}

func (s *SubscriptionV2RepositoryTestSuite) SetupSuite() {
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

func (s *SubscriptionV2RepositoryTestSuite) SetupTest() {
	var err error
	_, err = bootstrapSubscription(s.client)
	if err != nil {
		s.T().Fatal(err)
	}
	_, err = bootstrapSubscriptionReceiver(s.client)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionV2RepositoryTestSuite) TearDownSuite() {
	// Clean tests
	if err := purgeDocker(s.pool, s.resource); err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionV2RepositoryTestSuite) TearDownTest() {
	if err := s.cleanup(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *SubscriptionV2RepositoryTestSuite) cleanup() error {
	queries := []string{
		"TRUNCATE TABLE subscriptions RESTART IDENTITY CASCADE",
		"TRUNCATE TABLE subscriptions_receivers RESTART IDENTITY CASCADE",
	}
	return execQueries(context.TODO(), s.client, queries)
}

func (s *SubscriptionV2RepositoryTestSuite) TestMatchLabelsFetchReceivers() {
	type testCase struct {
		Description           string
		Filter                subscription.Filter // match and namespace_id only
		ExpectedReceiverViews []subscription.ReceiverView
		ErrString             string
	}

	var testCases = []testCase{
		{
			Description: "should get filtered receiver views by namespace id",
			Filter: subscription.Filter{
				NamespaceID: 1,
				Match: map[string]string{
					"environment": "integration",
					"team":        "gotocompany-data",
					"k1":          "v1",
					"k2":          "v2",
				},
			},
			ExpectedReceiverViews: []subscription.ReceiverView{
				{
					ID:             3,
					SubscriptionID: 2,
					Name:           "gotocompany_pagerduty",
					Labels: map[string]string{
						"entity": "gotocompany",
						"team":   "siren-gotocompany",
					},
					Type: receiver.TypePagerDuty,
					Configurations: map[string]any{
						"service_key": "1212121212121212121212121",
					},
					Match: map[string]string{
						"environment": "integration",
						"team":        "gotocompany-data",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			got, err := s.repository.MatchLabelsFetchReceivers(s.ctx, tc.Filter)
			if err != nil && err.Error() != tc.ErrString {
				s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
			}
			if diff := cmp.Diff(got, tc.ExpectedReceiverViews, cmpopts.IgnoreFields(subscription.ReceiverView{}, "CreatedAt", "UpdatedAt")); diff != "" {
				s.T().Fatalf("got diff %+v", diff)
			}
		})
	}
}

func TestSubscriptionV2Repository(t *testing.T) {
	suite.Run(t, new(SubscriptionV2RepositoryTestSuite))
}
