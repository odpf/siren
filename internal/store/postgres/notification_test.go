package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/goto/salt/dockertestx"
	"github.com/goto/salt/log"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/internal/store/postgres"
	"github.com/goto/siren/pkg/pgc"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
)

type NotificationRepositoryTestSuite struct {
	suite.Suite
	ctx        context.Context
	client     *pgc.Client
	pool       *dockertest.Pool
	resource   *dockertest.Resource
	repository *postgres.NotificationRepository
}

func (s *NotificationRepositoryTestSuite) SetupSuite() {
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
	s.Require().NoError(migrate(s.ctx, logger, s.client, dbConfig))
	s.repository = postgres.NewNotificationRepository(s.client)
}

func (s *NotificationRepositoryTestSuite) SetupTest() {
	var err error
	_, err = bootstrapNotification(s.client)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *NotificationRepositoryTestSuite) TearDownSuite() {
	// Clean tests
	if err := purgeDocker(s.pool, s.resource); err != nil {
		s.T().Fatal(err)
	}
}

func (s *NotificationRepositoryTestSuite) TearDownTest() {
	if err := s.cleanup(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *NotificationRepositoryTestSuite) cleanup() error {
	queries := []string{
		"TRUNCATE TABLE notifications RESTART IDENTITY CASCADE",
	}
	return execQueries(context.TODO(), s.client, queries)
}

func (s *NotificationRepositoryTestSuite) TestCreate() {
	type testCase struct {
		Description          string
		NotificationToCreate notification.Notification
		ErrString            string
	}

	var testCases = []testCase{
		{
			Description: "should create a notification",
			NotificationToCreate: notification.Notification{
				NamespaceID: 1,
				Type:        notification.TypeAlert,
				Data:        map[string]any{},
				Labels:      map[string]string{},
				CreatedAt:   time.Now(),
			},
		},
		{
			Description: "should return error if a notification is invalid",
			NotificationToCreate: notification.Notification{
				NamespaceID: 1,
				Type:        notification.TypeAlert,
				Data: map[string]any{
					"k1": func(x chan struct{}) {
						<-x
					},
				},
				Labels:    map[string]string{},
				CreatedAt: time.Now(),
			},
			ErrString: "sql: converting argument $3 type: json: unsupported type: func(chan struct {})",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			_, err := s.repository.Create(s.ctx, tc.NotificationToCreate)
			if err != nil {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
		})
	}
}

func (s *NotificationRepositoryTestSuite) TestList() {
	type testCase struct {
		Description           string
		Filter                notification.Filter
		ExpectedNotifications []notification.Notification
		ErrString             string
	}

	var testCases = []testCase{
		{
			Description: "should get all notifications with receiver_selector filter",
			Filter: notification.Filter{
				ReceiverSelector: map[string]string{
					"team": "gotocompany-infra",
				},
			},
			ExpectedNotifications: []notification.Notification{
				{
					ID:          "789",
					NamespaceID: 1,
					Type:        "alert",
					Data: map[string]interface{}{
						"data-key": "data-value",
					},
					Labels: map[string]string{
						"label-key": "label-value",
					},
					Template: "",
					ReceiverSelectors: []map[string]string{
						{
							"team":     "gotocompany-infra",
							"severity": "WARNING",
						},
						{
							"id": "2",
						},
					},
				},
			},
		},
		{
			Description: "should get all notifications with type filter",
			Filter: notification.Filter{
				Type: "event",
			},
			ExpectedNotifications: []notification.Notification{
				{
					ID:          "123",
					NamespaceID: 1,
					Type:        "event",
					Data: map[string]interface{}{
						"data-key": "data-value",
					},
					Labels:        map[string]string{},
					Template:      "",
					ValidDuration: 0,
					UniqueKey:     "",
				},
				{
					ID:          "10911",
					NamespaceID: 2,
					Type:        "event",
					Data: map[string]any{
						"data-key": "data-value",
					},
					Labels:        map[string]string{},
					ValidDuration: time.Duration(0),
					Template:      "expiry-alert",
					UniqueKey:     "",
				},
			},
		},
		{
			Description: "should get all notifications with template filter",
			Filter: notification.Filter{
				Template: "expiry-alert",
			},
			ExpectedNotifications: []notification.Notification{
				{
					ID:          "10911",
					NamespaceID: 2,
					Type:        "event",
					Data: map[string]any{
						"data-key": "data-value",
					},
					Labels:        map[string]string{},
					ValidDuration: time.Duration(0),
					Template:      "expiry-alert",
					UniqueKey:     "",
				},
			},
		},
		{
			Description: "should get all notifications with lable filter",
			Filter: notification.Filter{
				Labels: map[string]string{
					"label-key": "label-value",
				},
			},
			ExpectedNotifications: []notification.Notification{
				{
					ID:          "456",
					NamespaceID: 1,
					Type:        "alert",
					Data: map[string]interface{}{
						"data-key": "data-value",
					},
					Labels:   map[string]string{"label-key": "label-value"},
					Template: "",
				},
				{
					ID:          "789",
					NamespaceID: 1,
					Type:        "alert",
					Data: map[string]interface{}{
						"data-key": "data-value",
					},
					Labels: map[string]string{
						"label-key": "label-value",
					},
					Template: "",
					ReceiverSelectors: []map[string]string{
						{
							"team":     "gotocompany-infra",
							"severity": "WARNING",
						},
						{
							"id": "2",
						},
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

			if diff := cmp.Diff(got, tc.ExpectedNotifications, cmpopts.IgnoreFields(notification.Notification{}, "CreatedAt", "ID")); diff != "" {
				s.T().Fatalf("got diff %+v", diff)
			}
		})
	}
}

func TestNotificationRepository(t *testing.T) {
	suite.Run(t, new(NotificationRepositoryTestSuite))
}
