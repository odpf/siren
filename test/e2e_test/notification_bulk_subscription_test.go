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

	"github.com/google/go-cmp/cmp"
	"github.com/goto/salt/db"
	"github.com/goto/siren/config"
	"github.com/goto/siren/core/notification"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/internal/server"
	"github.com/goto/siren/plugins"
	cortexv1plugin "github.com/goto/siren/plugins/providers/cortex/v1"
	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	testdatatemplate_test "github.com/goto/siren/test/e2e_test/testdata/templates"
	"github.com/mcuadros/go-defaults"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/yaml.v3"
)

type BulkNotificationSubscriptionTestSuite struct {
	suite.Suite
	cancelContext context.CancelFunc
	grpcClient    sirenv1beta1.SirenServiceClient
	dbClient      *db.Client
	cancelClient  func()
	appConfig     *config.Config
	testBench     *CortexTest
}

func (s *BulkNotificationSubscriptionTestSuite) SetupTest() {
	apiHTTPPort, err := getFreePort()
	s.Require().Nil(err)
	apiGRPCPort, err := getFreePort()
	s.Require().Nil(err)

	s.appConfig = &config.Config{}

	defaults.SetDefaults(s.appConfig)

	s.appConfig.Log.Level = "error"
	s.appConfig.Service = server.Config{
		Port: apiHTTPPort,
		GRPC: server.GRPCConfig{
			Port: apiGRPCPort,
		},
		EncryptionKey:         testEncryptionKey,
		SubscriptionV2Enabled: true,
	}
	s.appConfig.Notification = notification.Config{
		MessageHandler: notification.HandlerConfig{
			Enabled: false,
		},
		DLQHandler: notification.HandlerConfig{
			Enabled: false,
		},
		GroupBy: []string{
			"team",
			"service",
		},
		// SubscriptionV2Enabled: true,
	}
	s.appConfig.Telemetry.OpenTelemetry.Enabled = false

	s.testBench, err = InitCortexEnvironment(s.appConfig)
	s.Require().NoError(err)

	// setup custom cortex config
	// TODO host.docker.internal only works for docker-desktop to call a service in host (siren)
	pathProject, _ := os.Getwd()
	rootProject := filepath.Dir(filepath.Dir(pathProject))
	s.appConfig.Providers.PluginPath = filepath.Join(rootProject, "plugin")
	s.appConfig.Providers.Plugins = make(map[string]plugins.PluginConfig, 0)
	s.appConfig.Providers.Plugins["cortex"] = plugins.PluginConfig{
		Handshake: plugins.HandshakeConfig{
			ProtocolVersion:  cortexv1plugin.Handshake.ProtocolVersion,
			MagicCookieKey:   cortexv1plugin.Handshake.MagicCookieKey,
			MagicCookieValue: cortexv1plugin.Handshake.MagicCookieValue,
		},
		ServiceConfig: map[string]interface{}{
			"webhook_base_api": fmt.Sprintf("http://test:%d/v1beta1/alerts/cortex", apiHTTPPort),
			"group_wait":       "1s",
			"group_interval":   "1s",
			"repeat_interval":  "1s",
		},
	}

	// enable worker
	s.appConfig.Notification.MessageHandler.Enabled = true
	s.appConfig.Notification.DLQHandler.Enabled = true

	ctx, cancel := context.WithCancel(context.Background())
	s.cancelContext = cancel

	StartSirenServer(ctx, *s.appConfig)

	s.grpcClient, s.cancelClient, err = CreateClient(ctx, fmt.Sprintf("localhost:%d", apiGRPCPort))
	s.Require().NoError(err)

	_, err = s.grpcClient.CreateProvider(ctx, &sirenv1beta1.CreateProviderRequest{
		Host: fmt.Sprintf("http://%s", s.testBench.NginxHost),
		Urn:  "cortex-test",
		Name: "cortex-test",
		Type: "cortex",
	})
	s.Require().NoError(err)

	s.dbClient, err = db.New(s.testBench.PGConfig)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *BulkNotificationSubscriptionTestSuite) TearDownTest() {
	s.cancelClient()

	// Clean tests
	err := s.testBench.CleanUp()
	s.Require().NoError(err)

	s.cancelContext()
}

func (s *BulkNotificationSubscriptionTestSuite) TestSendBulkNotification() {
	ctx := context.Background()

	_, err := s.grpcClient.CreateNamespace(ctx, &sirenv1beta1.CreateNamespaceRequest{
		Name:        "new-gotocompany-1",
		Urn:         "new-gotocompany-1",
		Provider:    1,
		Credentials: nil,
		Labels: map[string]string{
			"key1": "value1",
		},
	})
	s.Require().NoError(err)

	s.Run("sending bulk notification with same group labels should trigger only 1 notification", func() {
		waitChan := make(chan struct{}, 1)

		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bodyBytes, err := io.ReadAll(r.Body)
			s.Assert().NoError(err)
			fmt.Println(string(bodyBytes))
			type sampleStruct struct {
				Title             string `json:"title"`
				Desription        string `json:"description"`
				MergedTeam        string `json:"merged_team"`
				MergedService     string `json:"merged_service"`
				MergedEnvironment string `json:"merged_environment"`
				MergedCategory    string `json:"merged_category"`
			}

			expectedNotification := sampleStruct{
				Title: "This is the test notification with template",
				Desription: `Plain flow scalars are picky about the (:) and (#) characters. 
They can be in the string, but (:) cannot appear before a space or newline.
And (#) cannot appear after a space or newline; doing this will cause a syntax error. 
If you need to use these characters you are probably better off using one of the quoted styles instead.
`,
				MergedCategory:    "httpreceiver",
				MergedEnvironment: "integration development production",
				MergedService:     "some-service some-service some-service",
				MergedTeam:        "gotocompany gotocompany gotocompany",
			}
			var (
				resultStruct sampleStruct
			)
			s.Assert().NoError(json.Unmarshal(bodyBytes, &resultStruct))

			if diff := cmp.Diff(expectedNotification, resultStruct); diff != "" {
				s.T().Errorf("got diff: %v", diff)
			}
			waitChan <- struct{}{}

		}))
		defer testServer.Close()

		configs, err := structpb.NewStruct(map[string]any{
			"url": testServer.URL,
		})
		s.Require().NoError(err)
		_, err = s.grpcClient.CreateReceiver(ctx, &sirenv1beta1.CreateReceiverRequest{
			Name: "gotocompany-http",
			Type: "http",
			Labels: map[string]string{
				"entity": "gotocompany",
				"kind":   "http",
				"id":     "1",
			},
			Configurations: configs,
		})
		s.Require().NoError(err)

		sampleTemplateFile, err := template.YamlStringToFile(testdatatemplate_test.SampleBulkMessageTemplate)
		s.Require().NoError(err)

		body, err := yaml.Marshal(sampleTemplateFile.Body)
		s.Require().NoError(err)

		variables := make([]*sirenv1beta1.TemplateVariables, 0)
		for _, variable := range sampleTemplateFile.Variables {
			variables = append(variables, &sirenv1beta1.TemplateVariables{
				Name:        variable.Name,
				Type:        variable.Type,
				Default:     variable.Default,
				Description: variable.Description,
			})
		}

		_, err = s.grpcClient.UpsertTemplate(ctx, &sirenv1beta1.UpsertTemplateRequest{
			Name:      sampleTemplateFile.Name,
			Body:      string(body),
			Tags:      sampleTemplateFile.Tags,
			Variables: variables,
		})
		s.Require().NoError(err)

		sub, err := s.grpcClient.CreateSubscription(ctx, &sirenv1beta1.CreateSubscriptionRequest{
			Urn:       "subscribe-http-three",
			Namespace: 1,
			Match: map[string]string{
				"team":    "gotocompany",
				"service": "some-service",
			},
		})
		s.Require().NoError(err)

		_, err = s.grpcClient.AddSubscriptionReceiver(ctx, &sirenv1beta1.AddSubscriptionReceiverRequest{
			SubscriptionId: sub.GetId(),
			ReceiverId:     1,
		})
		s.Require().NoError(err)

		data, err := structpb.NewStruct(map[string]any{
			"title":      "This is the test notification with template",
			"icon_emoji": ":smile:",
		})
		s.Require().NoError(err)

		_, err = s.grpcClient.PostBulkNotifications(ctx, &sirenv1beta1.PostBulkNotificationsRequest{
			Notifications: []*sirenv1beta1.Notification{
				{
					Data: data,
					Labels: map[string]string{
						"team":        "gotocompany",
						"service":     "some-service",
						"environment": "integration",
						"category":    "httpreceiver",
					},
					Template: sampleTemplateFile.Name,
				},
				{
					Data: data,
					Labels: map[string]string{
						"team":        "gotocompany",
						"service":     "some-service",
						"environment": "development",
					},
					Template: sampleTemplateFile.Name,
				},
				{
					Data: data,
					Labels: map[string]string{
						"team":        "gotocompany",
						"service":     "some-service",
						"environment": "production",
					},
					Template: sampleTemplateFile.Name,
				},
			},
		})
		s.Assert().NoError(err)

		<-waitChan
	})
}

func TestBulkNotificationSubscriptionTestSuite(t *testing.T) {
	suite.Run(t, new(BulkNotificationSubscriptionTestSuite))
}
