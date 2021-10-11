package receiver

import (
	"errors"
	"github.com/odpf/siren/domain"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ServiceTestSuite struct {
	suite.Suite
	exchangerMock  *MockExchanger
	repositoryMock *MockReceiverRepository
}

func (s *ServiceTestSuite) SetupTest() {
	s.exchangerMock = &MockExchanger{}
	s.repositoryMock = &MockReceiverRepository{}
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) TestService_ListReceivers() {
	configurations := make(StringInterfaceMap)
	labels := make(StringStringMap)
	labels["foo"] = "bar"

	s.Run("should call repository List method and return result in domain's type", func() {
		dummyService := Service{repository: s.repositoryMock, exchanger: s.exchangerMock}
		dummyReceivers := []*domain.Receiver{
			{
				Id:             10,
				Name:           "foo",
				Type:           "slack",
				Labels:         labels,
				Configurations: configurations,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
		}
		receivers := []*Receiver{
			{
				Id:             10,
				Name:           "foo",
				Type:           "slack",
				Labels:         labels,
				Configurations: configurations,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
		}
		s.repositoryMock.On("List").Return(receivers, nil).Once()

		result, err := dummyService.ListReceivers()
		s.Nil(err)
		s.Equal(len(dummyReceivers), len(result))
		s.Equal(dummyReceivers[0].Name, result[0].Name)
		s.repositoryMock.AssertCalled(s.T(), "List")
	})

	s.Run("should call repository List method and return error if any", func() {
		dummyService := Service{repository: s.repositoryMock, exchanger: s.exchangerMock}
		s.repositoryMock.On("List").
			Return(nil, errors.New("random error")).Once()

		result, err := dummyService.ListReceivers()
		s.Nil(result)
		s.EqualError(err, "service.repository.List: random error")
		s.repositoryMock.AssertCalled(s.T(), "List")
	})
}

func (s *ServiceTestSuite) TestService_CreateReceiver() {
	configurations := make(StringInterfaceMap)
	configurations["client_id"] = "foo"
	configurations["client_secret"] = "bar"
	configurations["auth_code"] = "foo"

	labels := make(StringStringMap)
	labels["foo"] = "bar"
	timenow := time.Now()

	receiver := &domain.Receiver{
		Id:             10,
		Name:           "foo",
		Type:           "slack",
		Labels:         labels,
		Configurations: configurations,
		CreatedAt:      timenow,
		UpdatedAt:      timenow,
	}

	receiverRequest := &Receiver{
		Id:     10,
		Name:   "foo",
		Type:   "slack",
		Labels: labels,
		Configurations: map[string]interface{}{
			"workspace": "test-name",
			"token":     "token",
		},
		CreatedAt: timenow,
		UpdatedAt: timenow,
	}

	receiverResponse := &domain.Receiver{
		Id:     10,
		Name:   "foo",
		Type:   "slack",
		Labels: labels,
		Configurations: map[string]interface{}{
			"workspace": "test-name",
			"token":     "token",
		},
		CreatedAt: timenow,
		UpdatedAt: timenow,
	}

	codeExchangeHTTPResponse := CodeExchangeHTTPResponse{
		AccessToken: "token",
		Team: struct {
			Name string `json:"name"`
		}{
			Name: "test-name",
		},
	}

	s.Run("should call repository Create method and return result in domain's type", func() {
		dummyService := Service{repository: s.repositoryMock, exchanger: s.exchangerMock}
		s.repositoryMock.On("Create", receiverRequest).Return(receiverRequest, nil).Once()
		s.exchangerMock.On("Exchange", "foo", "foo", "bar").
			Return(codeExchangeHTTPResponse, nil).Once()

		result, err := dummyService.CreateReceiver(receiver)
		s.Nil(err)
		s.Equal(receiverResponse, result)
		s.repositoryMock.AssertCalled(s.T(), "Create", receiverRequest)
	})

	s.Run("should call repository Create method and return error if slack exchange failed", func() {
		dummyService := Service{repository: s.repositoryMock, exchanger: s.exchangerMock}
		s.exchangerMock.On("Exchange", "foo", "foo", "bar").
			Return(CodeExchangeHTTPResponse{}, errors.New("random error")).Once()

		result, err := dummyService.CreateReceiver(receiver)
		s.Nil(result)
		s.EqualError(err, "failed to exchange code with slack OAuth server: random error")
	})

	s.Run("should call repository Create method and return error if any", func() {
		dummyService := Service{repository: s.repositoryMock, exchanger: s.exchangerMock}
		s.exchangerMock.On("Exchange", "foo", "foo", "bar").
			Return(codeExchangeHTTPResponse, nil).Once()
		s.repositoryMock.On("Create", receiverRequest).
			Return(nil, errors.New("random error")).Once()

		result, err := dummyService.CreateReceiver(receiver)
		s.Nil(result)
		s.EqualError(err, "service.repository.Create: random error")
		s.repositoryMock.AssertCalled(s.T(), "Create", receiverRequest)
	})
}

func (s *ServiceTestSuite) TestService_GetReceiver() {
	receiverID := uint64(10)
	configurations := make(StringInterfaceMap)
	configurations["foo"] = "bar"

	labels := make(StringStringMap)
	labels["foo"] = "bar"
	timenow := time.Now()
	dummyReceiver := &domain.Receiver{
		Id:             10,
		Name:           "foo",
		Type:           "slack",
		Labels:         labels,
		Configurations: configurations,
		CreatedAt:      timenow,
		UpdatedAt:      timenow,
	}
	receiver := &Receiver{
		Id:             10,
		Name:           "foo",
		Type:           "slack",
		Labels:         labels,
		Configurations: configurations,
		CreatedAt:      timenow,
		UpdatedAt:      timenow,
	}

	s.Run("should call repository Get method and return result in domain's type", func() {
		dummyService := Service{repository: s.repositoryMock}
		s.repositoryMock.On("Get", receiverID).Return(receiver, nil).Once()

		result, err := dummyService.GetReceiver(receiverID)
		s.Nil(err)
		s.Equal(dummyReceiver, result)
		s.repositoryMock.AssertCalled(s.T(), "Get", receiverID)
	})

	s.Run("should call repository Get method and return error if any", func() {
		dummyService := Service{repository: s.repositoryMock}
		s.repositoryMock.On("Get", receiverID).
			Return(nil, errors.New("random error")).Once()

		result, err := dummyService.GetReceiver(receiverID)
		s.Nil(result)
		s.EqualError(err, "random error")
		s.repositoryMock.AssertCalled(s.T(), "Get", receiverID)
	})

}

func (s *ServiceTestSuite) TestService_UpdateReceiver() {
	timenow := time.Now()
	configurations := make(StringInterfaceMap)
	configurations["client_id"] = "foo"
	configurations["client_secret"] = "bar"
	configurations["auth_code"] = "foo"

	labels := make(StringStringMap)
	labels["foo"] = "bar"
	receiver := &domain.Receiver{
		Id:             10,
		Name:           "foo",
		Type:           "slack",
		Labels:         labels,
		Configurations: configurations,
		CreatedAt:      timenow,
		UpdatedAt:      timenow,
	}
	receiverRequest := &Receiver{
		Id:     10,
		Name:   "foo",
		Type:   "slack",
		Labels: labels,
		Configurations: map[string]interface{}{
			"workspace": "test-name",
			"token":     "token",
		},
		CreatedAt: timenow,
		UpdatedAt: timenow,
	}
	receiverResponse := &domain.Receiver{
		Id:     10,
		Name:   "foo",
		Type:   "slack",
		Labels: labels,
		Configurations: map[string]interface{}{
			"workspace": "test-name",
			"token":     "token",
		},
		CreatedAt: timenow,
		UpdatedAt: timenow,
	}

	codeExchangeHTTPResponse := CodeExchangeHTTPResponse{
		AccessToken: "token",
		Team: struct {
			Name string `json:"name"`
		}{
			Name: "test-name",
		},
	}

	s.Run("should call repository Update method and return result in domain's type", func() {
		dummyService := Service{repository: s.repositoryMock, exchanger: s.exchangerMock}
		s.repositoryMock.On("Update", receiverRequest).Return(receiverRequest, nil).Once()
		s.exchangerMock.On("Exchange", "foo", "foo", "bar").
			Return(codeExchangeHTTPResponse, nil).Once()

		result, err := dummyService.UpdateReceiver(receiver)
		s.Nil(err)
		s.Equal(receiverResponse, result)
		s.repositoryMock.AssertCalled(s.T(), "Update", receiverRequest)
	})

	s.Run("should call repository Create method and return error if slack exchange failed", func() {
		dummyService := Service{repository: s.repositoryMock, exchanger: s.exchangerMock}
		s.exchangerMock.On("Exchange", "foo", "foo", "bar").
			Return(CodeExchangeHTTPResponse{}, errors.New("random error")).Once()

		result, err := dummyService.UpdateReceiver(receiver)
		s.Nil(result)
		s.EqualError(err, "failed to exchange code with slack OAuth server: random error")
	})

	s.Run("should call repository Update method and return error if any", func() {
		dummyService := Service{repository: s.repositoryMock, exchanger: s.exchangerMock}
		s.repositoryMock.On("Update", receiverRequest).
			Return(nil, errors.New("random error")).Once()
		s.exchangerMock.On("Exchange", "foo", "foo", "bar").
			Return(codeExchangeHTTPResponse, nil).Once()

		result, err := dummyService.UpdateReceiver(receiver)
		s.Nil(result)
		s.EqualError(err, "random error")
		s.repositoryMock.AssertCalled(s.T(), "Update", receiverRequest)
	})
}

func (s *ServiceTestSuite) TestService_DeleteReceiver() {
	configurations := make(StringInterfaceMap)
	configurations["foo"] = "bar"

	labels := make(StringStringMap)
	labels["foo"] = "bar"
	receiverID := uint64(10)

	s.Run("should call repository Delete method and return nil if no error", func() {
		dummyService := Service{repository: s.repositoryMock}
		s.repositoryMock.On("Delete", receiverID).Return(nil).Once()

		err := dummyService.DeleteReceiver(receiverID)
		s.Nil(err)
		s.repositoryMock.AssertCalled(s.T(), "Delete", receiverID)
	})

	s.Run("should call repository Delete method and return error if any", func() {
		dummyService := Service{repository: s.repositoryMock}
		s.repositoryMock.On("Delete", receiverID).
			Return(errors.New("random error")).Once()

		err := dummyService.DeleteReceiver(receiverID)
		s.EqualError(err, "random error")
		s.repositoryMock.AssertCalled(s.T(), "Delete", receiverID)
	})
}

func (s *ServiceTestSuite) TestService_Migrate() {
	s.Run("should call repository Migrate method and return result", func() {
		dummyService := Service{repository: s.repositoryMock}
		s.repositoryMock.On("Migrate").Return(nil).Once()

		err := dummyService.Migrate()
		s.Nil(err)
		s.repositoryMock.AssertCalled(s.T(), "Migrate")
	})
}
