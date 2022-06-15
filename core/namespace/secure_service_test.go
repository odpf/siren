package namespace_test

import (
	"errors"
	testing "testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/odpf/siren/core/namespace"
	"github.com/odpf/siren/core/namespace/mocks"
	mock "github.com/stretchr/testify/mock"
)

func TestSecureService_ListNamespaces(t *testing.T) {
	type testCase struct {
		Description        string
		ExpectedNamespaces []*namespace.Namespace
		Setup              func(*mocks.NamespaceRepository, *mocks.Encryptor, testCase)
		Err                error
	}
	var (
		timeNow   = time.Now()
		testCases = []testCase{

			{
				Description: "should return error if List repository error",
				Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().List().Return(nil, errors.New("some error"))
				},
				Err: errors.New("secureService.repository.List: some error"),
			},
			{
				Description: "should return error if List repository success and decrypt error",
				Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().List().Return([]*namespace.EncryptedNamespace{
						{
							Namespace: &namespace.Namespace{
								Id:        1,
								Provider:  1,
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							Credentials: `encrypted-text-1`,
						},
						{
							Namespace: &namespace.Namespace{
								Id:        2,
								Provider:  1,
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							Credentials: `encrypted-text-2`,
						},
					}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("string")).Return("", errors.New("decrypt error"))
				},
				Err: errors.New("secureService.cryptoClient.Decrypt: decrypt error"),
			},
			{
				Description: "should return error if list repository success and decrypted object is not json",
				Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().List().Return([]*namespace.EncryptedNamespace{
						{
							Namespace: &namespace.Namespace{
								Id:        1,
								Provider:  1,
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							Credentials: `encrypted-text-1`,
						},
						{
							Namespace: &namespace.Namespace{
								Id:        2,
								Provider:  1,
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							Credentials: `encrypted-text-2`,
						},
					}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("string")).Return("", nil)
				},
				Err: errors.New("json.Unmarshal: unexpected end of JSON input"),
			},
			{
				Description: "should success if list repository and decrypt success",
				Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().List().Return([]*namespace.EncryptedNamespace{
						{
							Namespace: &namespace.Namespace{
								Id:        1,
								Provider:  1,
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							Credentials: `encrypted-text-1`,
						},
						{
							Namespace: &namespace.Namespace{
								Id:        2,
								Provider:  1,
								Name:      "foo",
								Labels:    map[string]string{"foo": "bar"},
								CreatedAt: timeNow,
								UpdatedAt: timeNow,
							},
							Credentials: `encrypted-text-2`,
						},
					}, nil)
					e.EXPECT().Decrypt(mock.AnythingOfType("string")).Return("{\"name\": \"a\"}", nil)
				},
				ExpectedNamespaces: []*namespace.Namespace{
					{
						Id:       1,
						Provider: 1,
						Name:     "foo",
						Labels:   map[string]string{"foo": "bar"},
						Credentials: map[string]interface{}{
							"name": "a",
						},
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
					},
					{
						Id:       2,
						Provider: 1,
						Name:     "foo",
						Labels:   map[string]string{"foo": "bar"},
						Credentials: map[string]interface{}{
							"name": "a",
						},
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
					},
				},
				Err: nil,
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock = new(mocks.NamespaceRepository)
				encryptorMock  = new(mocks.Encryptor)
			)
			svc := namespace.NewSecureService(encryptorMock, repositoryMock)

			tc.Setup(repositoryMock, encryptorMock, tc)

			got, err := svc.ListNamespaces()
			if tc.Err != err {
				if tc.Err.Error() != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.Err.Error())
				}
			}
			if !cmp.Equal(got, tc.ExpectedNamespaces) {
				t.Fatalf("got result %+v, expected was %+v", got, tc.ExpectedNamespaces)
			}
			repositoryMock.AssertExpectations(t)
			encryptorMock.AssertExpectations(t)
		})
	}
}

func TestSecureService_CreateNamespace(t *testing.T) {
	type testCase struct {
		Description string
		NSpace      *namespace.Namespace
		Setup       func(*mocks.NamespaceRepository, *mocks.Encryptor, testCase)
		Err         error
	}
	var testCases = []testCase{
		{
			Description: "should return error if encrypt return error caused credential is not in json",
			Setup:       func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {},
			NSpace: &namespace.Namespace{
				Credentials: map[string]interface{}{
					"invalid": make(chan int),
				},
			},
			Err: errors.New("json.Marshal: json: unsupported type: chan int"),
		},
		{
			Description: "should return error if encrypt return error",
			Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
				e.EXPECT().Encrypt(mock.AnythingOfType("string")).Return("", errors.New("some error"))
			},
			NSpace: &namespace.Namespace{
				Credentials: map[string]interface{}{
					"credential": "value",
				},
			},
			Err: errors.New("secureService.cryptoClient.Encrypt: some error"),
		},
		{
			Description: "should return error if encrypt success and create repository error",
			Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
				e.EXPECT().Encrypt(mock.AnythingOfType("string")).Return("some-ciphertext", nil)
				rr.EXPECT().Create(&namespace.EncryptedNamespace{
					Namespace:   tc.NSpace,
					Credentials: "some-ciphertext",
				}).Return(errors.New("some error"))
			},
			NSpace: &namespace.Namespace{
				Credentials: map[string]interface{}{
					"credential": "value",
				},
			},
			Err: errors.New("secureService.repository.Create: some error"),
		},
		{
			Description: "should return nil error if encrypt success and create repository success",
			Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
				e.EXPECT().Encrypt(mock.AnythingOfType("string")).Return("some-ciphertext", nil)
				rr.EXPECT().Create(&namespace.EncryptedNamespace{
					Namespace:   tc.NSpace,
					Credentials: "some-ciphertext",
				}).Return(nil)
			},
			NSpace: &namespace.Namespace{
				Credentials: map[string]interface{}{
					"credential": "value",
				},
			},
			Err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock = new(mocks.NamespaceRepository)
				encryptorMock  = new(mocks.Encryptor)
			)
			svc := namespace.NewSecureService(encryptorMock, repositoryMock)

			tc.Setup(repositoryMock, encryptorMock, tc)

			err := svc.CreateNamespace(tc.NSpace)
			if tc.Err != err {
				if tc.Err.Error() != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.Err.Error())
				}
			}

			repositoryMock.AssertExpectations(t)
			encryptorMock.AssertExpectations(t)
		})
	}
}

func TestSecureService_GetNamespace(t *testing.T) {
	type testCase struct {
		Description string
		NSpace      *namespace.Namespace
		Setup       func(*mocks.NamespaceRepository, *mocks.Encryptor, testCase)
		Err         error
	}
	var (
		testID    = uint64(10)
		testCases = []testCase{
			{
				Description: "should return error if Get repository error",
				Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().Get(testID).Return(nil, errors.New("some error"))
				},
				Err: errors.New("secureService.repository.Get: some error"),
			},
			{
				Description: "should return error if Get repository success and decrypt return error",
				Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().Get(testID).Return(&namespace.EncryptedNamespace{
						Namespace:   tc.NSpace,
						Credentials: "some-ciphertext",
					}, nil)
					e.EXPECT().Decrypt("some-ciphertext").Return("", errors.New("some error"))
				},
				Err: errors.New("secureService.cryptoClient.Decrypt: some error"),
			},
			{
				Description: "should return error if Get repository success and decrypted credentials is not json marshallable",
				Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().Get(testID).Return(&namespace.EncryptedNamespace{
						Namespace:   tc.NSpace,
						Credentials: "some-ciphertext",
					}, nil)
					e.EXPECT().Decrypt("some-ciphertext").Return("", nil)
				},
				Err: errors.New("json.Unmarshal: unexpected end of JSON input"),
			},
			{
				Description: "should return nil error if Get repository success and decrypt success",
				Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
					rr.EXPECT().Get(testID).Return(&namespace.EncryptedNamespace{
						Namespace:   tc.NSpace,
						Credentials: "some-ciphertext",
					}, nil)
					e.EXPECT().Decrypt("some-ciphertext").Return("{ \"key\": \"value\" }", nil)
				},
				NSpace: &namespace.Namespace{
					Credentials: map[string]interface{}{
						"key": "value",
					},
				},
				Err: nil,
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock = new(mocks.NamespaceRepository)
				encryptorMock  = new(mocks.Encryptor)
			)
			svc := namespace.NewSecureService(encryptorMock, repositoryMock)

			tc.Setup(repositoryMock, encryptorMock, tc)

			got, err := svc.GetNamespace(testID)
			if tc.Err != err {
				if tc.Err.Error() != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.Err.Error())
				}
			}
			if !cmp.Equal(got, tc.NSpace) {
				t.Fatalf("got result %+v, expected was %+v", got, tc.NSpace)
			}
			repositoryMock.AssertExpectations(t)
			encryptorMock.AssertExpectations(t)
		})
	}
}

func TestSecureService_UpdateNamespace(t *testing.T) {
	type testCase struct {
		Description string
		NSpace      *namespace.Namespace
		Setup       func(*mocks.NamespaceRepository, *mocks.Encryptor, testCase)
		Err         error
	}
	var testCases = []testCase{
		{
			Description: "should return error if encrypt return error caused credential is not in json",
			Setup:       func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {},
			NSpace: &namespace.Namespace{
				Credentials: map[string]interface{}{
					"invalid": make(chan int),
				},
			},
			Err: errors.New("json.Marshal: json: unsupported type: chan int"),
		},
		{
			Description: "should return error if encrypt return error",
			Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
				e.EXPECT().Encrypt(mock.AnythingOfType("string")).Return("", errors.New("some error"))
			},
			NSpace: &namespace.Namespace{
				Credentials: map[string]interface{}{
					"credential": "value",
				},
			},
			Err: errors.New("secureService.cryptoClient.Encrypt: some error"),
		},
		{
			Description: "should return error if encrypt success and update repository error",
			Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
				e.EXPECT().Encrypt(mock.AnythingOfType("string")).Return("some-ciphertext", nil)
				rr.EXPECT().Update(&namespace.EncryptedNamespace{
					Namespace:   tc.NSpace,
					Credentials: "some-ciphertext",
				}).Return(errors.New("some error"))
			},
			NSpace: &namespace.Namespace{
				Credentials: map[string]interface{}{
					"credential": "value",
				},
			},
			Err: errors.New("secureService.repository.Update: some error"),
		},
		{
			Description: "should return nil error if encrypt success and update repository success",
			Setup: func(rr *mocks.NamespaceRepository, e *mocks.Encryptor, tc testCase) {
				e.EXPECT().Encrypt(mock.AnythingOfType("string")).Return("some-ciphertext", nil)
				rr.EXPECT().Update(&namespace.EncryptedNamespace{
					Namespace:   tc.NSpace,
					Credentials: "some-ciphertext",
				}).Return(nil)
			},
			NSpace: &namespace.Namespace{
				Credentials: map[string]interface{}{
					"credential": "value",
				},
			},
			Err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			var (
				repositoryMock = new(mocks.NamespaceRepository)
				encryptorMock  = new(mocks.Encryptor)
			)
			svc := namespace.NewSecureService(encryptorMock, repositoryMock)

			tc.Setup(repositoryMock, encryptorMock, tc)

			err := svc.UpdateNamespace(tc.NSpace)
			if tc.Err != err {
				if tc.Err.Error() != err.Error() {
					t.Fatalf("got error %s, expected was %s", err.Error(), tc.Err.Error())
				}
			}

			repositoryMock.AssertExpectations(t)
			encryptorMock.AssertExpectations(t)
		})
	}
}
