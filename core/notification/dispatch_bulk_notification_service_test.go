package notification_test

// import (
// 	"context"
// 	"sort"
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// 	saltlog "github.com/goto/salt/log"
// 	"github.com/goto/siren/core/log"
// 	"github.com/goto/siren/core/notification"
// 	"github.com/goto/siren/core/notification/mocks"
// 	"github.com/goto/siren/core/template"
// 	"github.com/goto/siren/pkg/errors"
// 	"github.com/stretchr/testify/mock"
// )

// func TestDispatchBulkNotificationServiceService_Dispatch(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		n             []notification.Notification
// 		setup         func(*mocks.Repository, *mocks.AlertRepository, *mocks.Router, *mocks.LogService, *mocks.Queuer, *mocks.TemplateService, *mocks.Notifier)
// 		wantErrString string
// 	}{
// 		{
// 			name: "should return error if repository return error",
// 			n: []notification.Notification{
// 				{
// 					Type: notification.TypeEvent,
// 					Labels: map[string]string{
// 						"k1": "v1",
// 					},
// 				},
// 			},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer, ts *mocks.TemplateService, nt *mocks.Notifier) {
// 				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return(nil, errors.New("some error"))
// 			},
// 			wantErrString: "some error",
// 		},
// 		{
// 			name: "should return error if notification is not valid subscriber router",
// 			n:    []notification.Notification{},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer, ts *mocks.TemplateService, nt *mocks.Notifier) {
// 				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
// 					{
// 						Type:   notification.TypeEvent,
// 						Labels: map[string]string{},
// 					},
// 				}, nil)
// 			},
// 			wantErrString: "notification type subscriber should have labels: { 0 event map[] map[] 0s   [] 0001-01-01 00:00:00 +0000 utc []}",
// 		},
// 		{
// 			name: "should return error if router prepare meta messages return error",
// 			n:    []notification.Notification{},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer, ts *mocks.TemplateService, nt *mocks.Notifier) {
// 				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
// 					{
// 						Type: notification.TypeEvent,
// 						Labels: map[string]string{
// 							"k1": "v1",
// 						},
// 					},
// 				}, nil)
// 				mr.EXPECT().PrepareMetaMessages(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, nil, errors.New("some error"))
// 			},
// 			wantErrString: "some error",
// 		},
// 		{
// 			name: "should return error if log notifications return error",
// 			n:    []notification.Notification{},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer, ts *mocks.TemplateService, nt *mocks.Notifier) {
// 				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
// 					{
// 						Type: notification.TypeEvent,
// 						Labels: map[string]string{
// 							"k1": "v1",
// 						},
// 					},
// 				}, nil)
// 				mr.EXPECT().PrepareMetaMessages(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, []log.Notification{
// 					{
// 						NamespaceID:    3,
// 						SubscriptionID: 123,
// 						ReceiverID:     12,
// 					},
// 				}, nil)
// 				ls.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(errors.New("some error 3"))

// 			},
// 			wantErrString: "failed logging notifications: some error 3",
// 		},
// 		{
// 			name: "should return no error if successfully enqueue messages",
// 			n:    []notification.Notification{},
// 			setup: func(r *mocks.Repository, ar *mocks.AlertRepository, mr *mocks.Router, ls *mocks.LogService, q *mocks.Queuer, ts *mocks.TemplateService, nt *mocks.Notifier) {
// 				r.EXPECT().BulkCreate(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("[]notification.Notification")).Return([]notification.Notification{
// 					{
// 						Type: notification.TypeEvent,
// 						Labels: map[string]string{
// 							"k1": "v1",
// 						},
// 					},
// 				}, nil)
// 				mr.EXPECT().PrepareMetaMessages(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("notification.Notification")).Return(nil, []log.Notification{
// 					{
// 						NamespaceID:    3,
// 						SubscriptionID: 123,
// 						ReceiverID:     12,
// 					},
// 				}, nil)
// 				ls.EXPECT().LogNotifications(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("log.Notification")).Return(nil)
// 				nt.EXPECT().PreHookQueueTransformConfigs(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("map[string]any")).Return(map[string]any{}, nil)
// 				nt.EXPECT().GetSystemDefaultTemplate().Return("system-template")
// 				ts.EXPECT().GetByName(mock.AnythingOfType("context.todoCtx"), mock.AnythingOfType("string")).Return(&template.Template{}, nil)
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var (
// 				mockQueuer          = new(mocks.Queuer)
// 				mockRepository      = new(mocks.Repository)
// 				mockNotifier        = new(mocks.Notifier)
// 				mockTemplateService = new(mocks.TemplateService)
// 				mockAlertRepository = new(mocks.AlertRepository)
// 				mockLogService      = new(mocks.LogService)
// 				mockRouter          = new(mocks.Router)
// 			)

// 			if tt.setup != nil {
// 				tt.setup(mockRepository, mockAlertRepository, mockRouter, mockLogService, mockQueuer, mockTemplateService, mockNotifier)
// 			}

// 			s := notification.NewDispatchBulkNotificationService(
// 				notification.Deps{
// 					Cfg: notification.Config{
// 						EnableSilenceFeature: true,
// 					},
// 					Logger:          saltlog.NewNoop(),
// 					Repository:      mockRepository,
// 					Q:               mockQueuer,
// 					AlertRepository: mockAlertRepository,
// 					LogService:      mockLogService,
// 				},
// 				map[string]notification.Notifier{
// 					testType: mockNotifier,
// 				},
// 				map[string]notification.Router{
// 					testType:                      mockRouter,
// 					notification.RouterSubscriber: mockRouter,
// 				},
// 			)
// 			if _, err := s.Dispatch(context.TODO(), tt.n); err != nil {
// 				if err.Error() != tt.wantErrString {
// 					t.Fatalf("error = %v, wantErr %s", err, tt.wantErrString)
// 				}
// 			}
// 		})
// 	}
// }

// func TestReduceMetaMessages(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		metaMessages []notification.MetaMessage
// 		groupBy      []string
// 		want         []notification.MetaMessage
// 		wantErr      bool
// 	}{
// 		{
// 			name: "should group meta messages with receiver id",
// 			metaMessages: []notification.MetaMessage{
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"xx"},
// 					SubscriptionIDs: []uint64{1, 2},
// 					Data: map[string]any{
// 						"d1": "dv1",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "ab",
// 					},
// 				},
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"yy"},
// 					SubscriptionIDs: []uint64{3, 4},
// 					Data: map[string]any{
// 						"d2": "dv2",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "cd",
// 						"x1": "x1",
// 					},
// 				},
// 				{
// 					ReceiverID:      14,
// 					NotificationIDs: []string{"zz"},
// 					SubscriptionIDs: []uint64{3, 4},
// 					Data: map[string]any{
// 						"d3": "dv3",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-14",
// 						"k2": "ab",
// 					},
// 				},
// 			},
// 			want: []notification.MetaMessage{
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"xx", "yy"},
// 					SubscriptionIDs: []uint64{1, 2, 3, 4},
// 					Data: map[string]any{
// 						"d1": "dv1",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "ab",
// 					},
// 					MergedLabels: map[string][]string{
// 						"k1": {"receiver-13", "receiver-13"},
// 						"k2": {"ab", "cd"},
// 						"x1": {"x1"},
// 					},
// 				},
// 				{
// 					ReceiverID:      14,
// 					NotificationIDs: []string{"zz"},
// 					SubscriptionIDs: []uint64{3, 4},
// 					Data: map[string]any{
// 						"d3": "dv3",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-14",
// 						"k2": "ab",
// 					},
// 					MergedLabels: map[string][]string{
// 						"k1": {"receiver-14"},
// 						"k2": {"ab"},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "should not group meta messages if receiver and template are different",
// 			metaMessages: []notification.MetaMessage{
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"xx"},
// 					SubscriptionIDs: []uint64{1, 2},
// 					Data: map[string]any{
// 						"d1": "dv1",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "ab",
// 					},
// 					Template: "template-1",
// 				},
// 				{
// 					ReceiverID:      14,
// 					NotificationIDs: []string{"yy"},
// 					SubscriptionIDs: []uint64{3, 4},
// 					Data: map[string]any{
// 						"d2": "dv2",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "cd",
// 						"x1": "x1",
// 					},
// 					Template: "template-2",
// 				},
// 				{
// 					ReceiverID:      15,
// 					NotificationIDs: []string{"zz"},
// 					SubscriptionIDs: []uint64{3, 4},
// 					Data: map[string]any{
// 						"d3": "dv3",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-14",
// 						"k2": "ab",
// 					},
// 					Template: "template-1",
// 				},
// 			},
// 			want: []notification.MetaMessage{
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"xx"},
// 					SubscriptionIDs: []uint64{1, 2},
// 					Data: map[string]any{
// 						"d1": "dv1",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "ab",
// 					},
// 					Template:     "template-1",
// 					MergedLabels: map[string][]string{"k1": {"receiver-13"}, "k2": {"ab"}},
// 				},
// 				{
// 					ReceiverID:      14,
// 					NotificationIDs: []string{"yy"},
// 					SubscriptionIDs: []uint64{3, 4},
// 					Data: map[string]any{
// 						"d2": "dv2",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "cd",
// 						"x1": "x1",
// 					},
// 					Template:     "template-2",
// 					MergedLabels: map[string][]string{"k1": {"receiver-13"}, "k2": {"cd"}, "x1": {"x1"}},
// 				},
// 				{
// 					ReceiverID:      15,
// 					NotificationIDs: []string{"zz"},
// 					SubscriptionIDs: []uint64{3, 4},
// 					Data: map[string]any{
// 						"d3": "dv3",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-14",
// 						"k2": "ab",
// 					},
// 					Template:     "template-1",
// 					MergedLabels: map[string][]string{"k1": {"receiver-14"}, "k2": {"ab"}},
// 				},
// 			},
// 		},
// 		{
// 			name: "should group meta messages with group by labels, template and receiver id",
// 			groupBy: []string{
// 				"k1",
// 				"k2",
// 			},
// 			metaMessages: []notification.MetaMessage{
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"xx"},
// 					SubscriptionIDs: []uint64{1, 2},
// 					Data: map[string]any{
// 						"d1": "dv1",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "ab",
// 					},
// 				},
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"yy"},
// 					SubscriptionIDs: []uint64{3, 4},
// 					Data: map[string]any{
// 						"d2": "dv2",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "cd",
// 						"x1": "x1",
// 					},
// 				},
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"zz"},
// 					SubscriptionIDs: []uint64{3, 4},
// 					Data: map[string]any{
// 						"d3": "dv3",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "ab",
// 					},
// 				},
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"aa"},
// 					SubscriptionIDs: []uint64{5, 6},
// 					Data: map[string]any{
// 						"d3": "dv3",
// 					},
// 					Labels: map[string]string{
// 						"k2": "ab",
// 					},
// 				},
// 				{
// 					ReceiverID:      14,
// 					NotificationIDs: []string{"aa"},
// 					SubscriptionIDs: []uint64{5, 6},
// 					Data: map[string]any{
// 						"d3": "dv3",
// 					},
// 					Labels: map[string]string{
// 						"k2": "ab",
// 					},
// 				},
// 			},
// 			want: []notification.MetaMessage{
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"xx", "zz"},
// 					SubscriptionIDs: []uint64{1, 2, 3, 4},
// 					Data: map[string]any{
// 						"d1": "dv1",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "ab",
// 					},
// 					MergedLabels: map[string][]string{"k1": {"receiver-13", "receiver-13"}, "k2": {"ab", "ab"}},
// 				},
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"yy"},
// 					SubscriptionIDs: []uint64{3, 4},
// 					Data: map[string]any{
// 						"d2": "dv2",
// 					},
// 					Labels: map[string]string{
// 						"k1": "receiver-13",
// 						"k2": "cd",
// 						"x1": "x1",
// 					},
// 					MergedLabels: map[string][]string{"k1": {"receiver-13"}, "k2": {"cd"}, "x1": {"x1"}},
// 				},
// 				{
// 					ReceiverID:      13,
// 					NotificationIDs: []string{"aa"},
// 					SubscriptionIDs: []uint64{5, 6},
// 					Data: map[string]any{
// 						"d3": "dv3",
// 					},
// 					Labels: map[string]string{
// 						"k2": "ab",
// 					},
// 					MergedLabels: map[string][]string{"k2": {"ab"}},
// 				},
// 				{
// 					ReceiverID:      14,
// 					NotificationIDs: []string{"aa"},
// 					SubscriptionIDs: []uint64{5, 6},
// 					Data: map[string]any{
// 						"d3": "dv3",
// 					},
// 					Labels: map[string]string{
// 						"k2": "ab",
// 					},
// 					MergedLabels: map[string][]string{"k2": {"ab"}},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := notification.ReduceMetaMessages(tt.metaMessages, tt.groupBy)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ReduceMetaMessages() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			sort.Slice(got, func(i, j int) bool {
// 				return got[i].ReceiverID < got[j].ReceiverID
// 			})
// 			sort.Slice(tt.want, func(i, j int) bool {
// 				return tt.want[i].ReceiverID < tt.want[j].ReceiverID
// 			})
// 			if diff := cmp.Diff(got, tt.want); diff != "" {
// 				t.Errorf("ReduceMetaMessages() diff = %v", diff)
// 			}
// 		})
// 	}
// }
