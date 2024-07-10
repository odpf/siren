package e2e_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/goto/siren/config"
	"github.com/goto/siren/internal/server"
	"github.com/goto/siren/plugins"
	cortexv1plugin "github.com/goto/siren/plugins/providers/cortex/v1"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"github.com/mcuadros/go-defaults"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/structpb"
)

type NotificationReceiverTestSuite struct {
	suite.Suite
	cancelContext      context.CancelFunc
	client             sirenv1beta1.SirenServiceClient
	cancelClient       func()
	appConfig          *config.Config
	testBench          *CortexTest
	closeWorkerChannel chan struct{}
}

func (s *NotificationReceiverTestSuite) SetupTest() {

	apiPort, err := getFreePort()
	s.Require().Nil(err)

	s.appConfig = &config.Config{}

	defaults.SetDefaults(s.appConfig)

	s.appConfig.Log.Level = "error"
	s.appConfig.Service = server.Config{
		GRPC: server.GRPCConfig{
			Port: apiPort,
		},
		EncryptionKey:         testEncryptionKey,
		SubscriptionV2Enabled: true,
	}
	s.appConfig.Notification.MessageHandler.Enabled = true
	s.appConfig.Notification.DLQHandler.Enabled = false
	// s.appConfig.Notification.SubscriptionV2Enabled = true

	s.testBench, err = InitCortexEnvironment(s.appConfig)
	s.Require().NoError(err)

	// TODO host.docker.internal only works for docker-desktop to call a service in host (siren)
	s.appConfig.Providers.Plugins = make(map[string]plugins.PluginConfig, 0)
	pathProject, _ := os.Getwd()
	rootProject := filepath.Dir(filepath.Dir(pathProject))
	s.appConfig.Providers.PluginPath = filepath.Join(rootProject, "plugin")
	s.appConfig.Providers.Plugins["cortex"] = plugins.PluginConfig{
		Handshake: plugins.HandshakeConfig{
			ProtocolVersion:  cortexv1plugin.Handshake.ProtocolVersion,
			MagicCookieKey:   cortexv1plugin.Handshake.MagicCookieKey,
			MagicCookieValue: cortexv1plugin.Handshake.MagicCookieValue,
		},
		ServiceConfig: map[string]interface{}{
			"webhook_base_api": fmt.Sprintf("http://host.docker.internal:%d/v1beta1/alerts/cortex", apiPort),
			"group_wait":       "1s",
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.cancelContext = cancel

	StartSirenServer(ctx, *s.appConfig)

	s.closeWorkerChannel = make(chan struct{}, 1)

	time.Sleep(500 * time.Millisecond)
	StartSirenMessageWorker(ctx, *s.appConfig, s.closeWorkerChannel)

	s.client, s.cancelClient, err = CreateClient(ctx, fmt.Sprintf("localhost:%d", apiPort))
	s.Require().NoError(err)

	bootstrapCortexTestData(&s.Suite, ctx, s.client, s.testBench.NginxHost)
}

func (s *NotificationReceiverTestSuite) TearDownTest() {
	s.cancelClient()

	// Clean tests
	err := s.testBench.CleanUp()
	s.Require().NoError(err)

	s.closeWorkerChannel <- struct{}{}

	s.cancelContext()
}

func (s *NotificationReceiverTestSuite) TestSendNotificationSuccess() {
	ctx := context.Background()

	controlChan := make(chan struct{}, 1)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		s.Assert().NoError(err)

		type sampleStruct struct {
			ID               string `json:"id"`
			IconEmoji        string `json:"icon_emoji"`
			NotificationType string `json:"notification_type"`
			ReceiverID       string `json:"receiver_id"`
			Text             string `json:"text"`
		}

		expectedNotification := `{"icon_emoji":":smile:","notification_type":"event","text":"test send notification"}`

		var (
			resultStruct   sampleStruct
			expectedStruct sampleStruct
		)
		s.Assert().NoError(json.Unmarshal(bodyBytes, &resultStruct))
		s.Assert().NoError(json.Unmarshal([]byte(expectedNotification), &expectedStruct))

		if diff := cmp.Diff(expectedStruct, resultStruct, cmpopts.IgnoreFields(sampleStruct{}, "ID")); diff != "" {
			s.T().Errorf("got diff: %v", diff)
		}
		controlChan <- struct{}{}

	}))
	defer testServer.Close()

	// add test server http receiver
	configs, err := structpb.NewStruct(map[string]any{
		"url": testServer.URL,
	})
	s.Require().NoError(err)
	rcv, err := s.client.CreateReceiver(ctx, &sirenv1beta1.CreateReceiverRequest{
		Name:           "notification-http",
		Type:           "http",
		Labels:         nil,
		Configurations: configs,
	})
	s.Require().NoError(err)

	time.Sleep(100 * time.Millisecond)

	_, err = s.client.PostNotification(ctx, &sirenv1beta1.PostNotificationRequest{
		Receivers: []*structpb.Struct{
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
		},
		Data: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"text":       structpb.NewStringValue("test send notification"),
				"icon_emoji": structpb.NewStringValue(":smile:"),
			},
		},
	})
	s.Require().NoError(err)

	<-controlChan

}

func (s *NotificationReceiverTestSuite) TestSendNotificationFailureLimited() {
	ctx := context.Background()

	// add test server http receiver
	configs, err := structpb.NewStruct(map[string]any{
		"url": "dummy",
	})
	s.Require().NoError(err)
	rcv, err := s.client.CreateReceiver(ctx, &sirenv1beta1.CreateReceiverRequest{
		Name:           "notification-http",
		Type:           "http",
		Labels:         nil,
		Configurations: configs,
	})
	s.Require().NoError(err)

	time.Sleep(100 * time.Millisecond)

	_, err = s.client.PostNotification(ctx, &sirenv1beta1.PostNotificationRequest{
		Receivers: []*structpb.Struct{
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
			{
				Fields: map[string]*structpb.Value{
					"id": structpb.NewStringValue(fmt.Sprintf("%d", rcv.GetId())),
				},
			},
		},
		Data: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"text":       structpb.NewStringValue("test send notification"),
				"icon_emoji": structpb.NewStringValue(":smile:"),
			},
		},
	})
	s.Require().Equal("rpc error: code = InvalidArgument desc = number of receiver selectors should be less than or equal threshold 10", err.Error())

}

func TestNotificationReceiverTestSuite(t *testing.T) {
	suite.Run(t, new(NotificationReceiverTestSuite))
}
