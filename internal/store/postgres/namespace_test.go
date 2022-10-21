package postgres_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/odpf/salt/db"
	"github.com/odpf/salt/dockertest"
	"github.com/odpf/salt/log"
	"github.com/odpf/siren/core/namespace"
	"github.com/odpf/siren/core/provider"
	"github.com/odpf/siren/internal/store/postgres"
	"github.com/stretchr/testify/suite"
)

type NamespaceRepositoryTestSuite struct {
	suite.Suite
	ctx        context.Context
	client     *postgres.Client
	pool       *dockertest.Pool
	resource   *dockertest.Resource
	repository *postgres.NamespaceRepository
}

func (s *NamespaceRepositoryTestSuite) SetupSuite() {
	var err error

	logger := log.NewZap()
	dpg, err := dockertest.CreatePostgres(
		dockertest.PostgresWithDetail(
			pgUser, pgPass, pgDBName,
		),
	)
	if err != nil {
		s.T().Fatal(err)
	}

	s.pool = dpg.GetPool()
	s.resource = dpg.GetResource()

	dbConfig.URL = dpg.GetExternalConnString()
	dbc, err := db.New(dbConfig)
	if err != nil {
		s.T().Fatal(err)
	}

	s.client, err = postgres.NewClient(logger, dbc)
	if err != nil {
		s.T().Fatal(err)
	}
	s.ctx = context.TODO()
	migrate(s.ctx, logger, s.client, dbConfig)
	s.repository = postgres.NewNamespaceRepository(s.client)

	_, err = bootstrapProvider(s.client)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *NamespaceRepositoryTestSuite) SetupTest() {
	var err error
	_, err = bootstrapNamespace(s.client)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *NamespaceRepositoryTestSuite) TearDownSuite() {
	// Clean tests
	if err := purgeDocker(s.pool, s.resource); err != nil {
		s.T().Fatal(err)
	}
}

func (s *NamespaceRepositoryTestSuite) TearDownTest() {
	if err := s.cleanup(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *NamespaceRepositoryTestSuite) cleanup() error {
	queries := []string{
		"TRUNCATE TABLE namespaces RESTART IDENTITY CASCADE",
	}
	return execQueries(context.TODO(), s.client, queries)
}

func (s *NamespaceRepositoryTestSuite) TestList() {
	type testCase struct {
		Description        string
		ExpectedNamespaces []namespace.EncryptedNamespace
		ErrString          string
	}

	var testCases = []testCase{
		{
			Description: "should get all namespaces",
			ExpectedNamespaces: []namespace.EncryptedNamespace{
				{
					Namespace: &namespace.Namespace{
						ID:   1,
						Name: "odpf",
						URN:  "odpf",
						Provider: provider.Provider{
							ID:          1,
							Host:        "http://cortex-ingress.odpf.io",
							URN:         "odpf-cortex",
							Name:        "odpf-cortex",
							Type:        "cortex",
							Credentials: map[string]interface{}{},
							Labels:      map[string]string{},
						},
						Labels: map[string]string{},
					},
					CredentialString: "map[secret_key:odpf-secret-key-1]",
				},
				{
					Namespace: &namespace.Namespace{
						ID:   2,
						Name: "odpf",
						URN:  "odpf",
						Provider: provider.Provider{
							ID:          2,
							Host:        "http://prometheus-ingress.odpf.io",
							URN:         "odpf-prometheus",
							Name:        "odpf-prometheus",
							Type:        "prometheus",
							Credentials: map[string]interface{}{},
							Labels:      map[string]string{},
						},
						Labels: map[string]string{},
					},
					CredentialString: "map[secret_key:odpf-secret-key-2]",
				},
				{
					Namespace: &namespace.Namespace{
						ID:   3,
						Name: "instance-1",
						URN:  "instance-1",
						Provider: provider.Provider{
							ID:          2,
							Host:        "http://prometheus-ingress.odpf.io",
							URN:         "odpf-prometheus",
							Name:        "odpf-prometheus",
							Type:        "prometheus",
							Credentials: map[string]interface{}{},
							Labels:      map[string]string{},
						},
						Labels: map[string]string{},
					},
					CredentialString: "map[service_key:instance-1-service-key]",
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			got, err := s.repository.List(s.ctx)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
			if !cmp.Equal(got, tc.ExpectedNamespaces, cmpopts.IgnoreFields(namespace.EncryptedNamespace{},
				"Namespace.CreatedAt", "Namespace.UpdatedAt",
				"Namespace.Provider.CreatedAt", "Namespace.Provider.UpdatedAt")) {
				s.T().Fatalf("got result %+v, expected was %+v", got, tc.ExpectedNamespaces)
			}
		})
	}
}

func (s *NamespaceRepositoryTestSuite) TestGet() {
	type testCase struct {
		Description       string
		PassedID          uint64
		ExpectedNamespace *namespace.EncryptedNamespace
		ErrString         string
	}

	var testCases = []testCase{
		{
			Description: "should get a namespace",
			PassedID:    3,
			ExpectedNamespace: &namespace.EncryptedNamespace{
				Namespace: &namespace.Namespace{
					ID:   3,
					Name: "instance-1",
					URN:  "instance-1",
					Provider: provider.Provider{
						ID:          2,
						Host:        "http://prometheus-ingress.odpf.io",
						URN:         "odpf-prometheus",
						Name:        "odpf-prometheus",
						Type:        "prometheus",
						Credentials: map[string]interface{}{},
						Labels:      map[string]string{},
					},
					Labels: map[string]string{},
				},
				CredentialString: "map[service_key:instance-1-service-key]",
			},
		},
		{
			Description: "should return not found if id not found",
			PassedID:    1000,
			ErrString:   "namespace with id 1000 not found",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			got, err := s.repository.Get(s.ctx, tc.PassedID)
			if err != nil && err.Error() != tc.ErrString {
				s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
			}
			if !cmp.Equal(got, tc.ExpectedNamespace, cmpopts.IgnoreFields(namespace.EncryptedNamespace{},
				"Namespace.CreatedAt", "Namespace.UpdatedAt",
				"Namespace.Provider.CreatedAt", "Namespace.Provider.UpdatedAt")) {
				s.T().Fatalf("got result %+v, expected was %+v", got, tc.ExpectedNamespace)
			}
		})
	}
}

func (s *NamespaceRepositoryTestSuite) TestCreate() {
	type testCase struct {
		Description       string
		NamespaceToCreate *namespace.EncryptedNamespace
		ExpectedID        uint64
		ErrString         string
	}

	var testCases = []testCase{
		{
			Description: "should create a namespace",
			NamespaceToCreate: &namespace.EncryptedNamespace{
				Namespace: &namespace.Namespace{
					Name: "instance-2",
					URN:  "instance-2",
					Provider: provider.Provider{
						ID: 2,
					},
				},
				CredentialString: "xxx",
			},
			ExpectedID: uint64(4), // autoincrement in db side
		},
		{
			Description: "should return error foreign key if provider id does not exist",
			NamespaceToCreate: &namespace.EncryptedNamespace{
				Namespace: &namespace.Namespace{
					Name: "odpf-new",
					URN:  "odpf",
					Provider: provider.Provider{
						ID: 1000,
					},
				},
				CredentialString: "xxx",
			},
			ErrString: "provider id does not exist",
		},
		{
			Description: "should return error duplicate if URN and provider already exist",
			NamespaceToCreate: &namespace.EncryptedNamespace{
				Namespace: &namespace.Namespace{
					Name: "odpf-new",
					URN:  "odpf",
					Provider: provider.Provider{
						ID: 2,
					},
				},
				CredentialString: "xxx",
			},
			ErrString: "urn and provider pair already exist",
		},
		{
			Description: "should return error if namespace is nil",
			ErrString:   "nil encrypted namespace domain when converting to namespace model",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.Create(s.ctx, tc.NamespaceToCreate)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
		})
	}
}

func (s *NamespaceRepositoryTestSuite) TestUpdate() {
	type testCase struct {
		Description       string
		NamespaceToUpdate *namespace.EncryptedNamespace
		ExpectedID        uint64
		ErrString         string
	}

	var testCases = []testCase{
		{
			Description: "should update existing namespace",
			NamespaceToUpdate: &namespace.EncryptedNamespace{
				Namespace: &namespace.Namespace{
					ID:   1,
					Name: "instance-updated",
					URN:  "instance-updated",
					Provider: provider.Provider{
						ID: 2,
					},
				},
				CredentialString: "xxx",
			},
			ExpectedID: uint64(1),
		},
		{
			Description: "should return error duplicate if URN and provider already exist",
			NamespaceToUpdate: &namespace.EncryptedNamespace{
				Namespace: &namespace.Namespace{
					ID:   3,
					Name: "new-odpf",
					URN:  "odpf",
					Provider: provider.Provider{
						ID: 2,
					},
				},
				CredentialString: "xxx",
			},
			ErrString: "urn and provider pair already exist",
		},
		{
			Description: "should return error not found if id not found",
			NamespaceToUpdate: &namespace.EncryptedNamespace{
				Namespace: &namespace.Namespace{
					ID:   1000,
					Name: "new-odpf",
					URN:  "odpf",
					Provider: provider.Provider{
						ID: 2,
					},
				},
				CredentialString: "xxx",
			},
			ErrString: "namespace with id 1000 not found",
		},
		{
			Description: "should return error foreign key if provider id does not exist",
			NamespaceToUpdate: &namespace.EncryptedNamespace{
				Namespace: &namespace.Namespace{
					ID:   1,
					Name: "odpf-new",
					URN:  "odpf",
					Provider: provider.Provider{
						ID: 1000,
					},
				},
				CredentialString: "xxx",
			},
			ErrString: "provider id does not exist",
		},
		{
			Description: "should return error duplicate if URN and provider already exist",
			NamespaceToUpdate: &namespace.EncryptedNamespace{
				Namespace: &namespace.Namespace{
					ID:   1,
					Name: "odpf-new",
					URN:  "odpf",
					Provider: provider.Provider{
						ID: 2,
					},
				},
				CredentialString: "xxx",
			},
			ErrString: "urn and provider pair already exist",
		},
		{
			Description: "should return error if namespace is nil",
			ErrString:   "nil encrypted namespace domain when converting to namespace model",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.Update(s.ctx, tc.NamespaceToUpdate)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
		})
	}
}

func (s *NamespaceRepositoryTestSuite) TestDelete() {
	type testCase struct {
		Description string
		IDToDelete  uint64
		ErrString   string
	}

	var testCases = []testCase{
		{
			Description: "should delete a namespace",
			IDToDelete:  1,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			err := s.repository.Delete(s.ctx, tc.IDToDelete)
			if tc.ErrString != "" {
				if err.Error() != tc.ErrString {
					s.T().Fatalf("got error %s, expected was %s", err.Error(), tc.ErrString)
				}
			}
		})
	}
}

func TestNamespaceRepository(t *testing.T) {
	suite.Run(t, new(NamespaceRepositoryTestSuite))
}
