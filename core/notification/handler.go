package notification

import (
	"context"
	"fmt"
	"time"

	"github.com/goto/salt/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/goto/siren/pkg/errors"
)

const (
	defaultBatchSize = 1
)

// Handler is a process to handle message publishing
type Handler struct {
	logger                 log.Logger
	q                      Queuer
	identifier             string
	notifierRegistry       map[string]Notifier
	supportedReceiverTypes []string
	batchSize              int
	metricHistMQDuration   metric.Int64Histogram
}

// NewHandler creates a new handler with some supported type of Notifiers
func NewHandler(cfg HandlerConfig, logger log.Logger, q Queuer, registry map[string]Notifier, opts ...HandlerOption) *Handler {
	metricHistMQDuration, err := otel.Meter("github.com/goto/siren/core/notification").
		Int64Histogram("siren.notification.queue.duration")
	if err != nil {
		otel.Handle(err)
	}
	h := &Handler{
		batchSize: defaultBatchSize,

		logger:               logger,
		notifierRegistry:     registry,
		q:                    q,
		metricHistMQDuration: metricHistMQDuration,
	}

	if cfg.BatchSize != 0 {
		h.batchSize = cfg.BatchSize
	}
	registeredReceivers := make([]string, 0, len(h.notifierRegistry))
	for k := range h.notifierRegistry {
		registeredReceivers = append(registeredReceivers, k)
	}
	h.supportedReceiverTypes = registeredReceivers

	if len(cfg.ReceiverTypes) != 0 {
		newSupportedReceiverTypes := []string{}
		for _, rt := range cfg.ReceiverTypes {
			found := false
			for _, k := range registeredReceivers {
				if rt == k {
					found = true
					break
				}
			}
			if found {
				newSupportedReceiverTypes = append(newSupportedReceiverTypes, rt)
			}
		}
		h.supportedReceiverTypes = newSupportedReceiverTypes
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func (h *Handler) getNotifierPlugin(receiverType string) (Notifier, error) {
	receiverPlugin, exist := h.notifierRegistry[receiverType]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported receiver type: %q on handler %s", receiverType, h.identifier)
	}
	return receiverPlugin, nil
}

func (h *Handler) Process(ctx context.Context, runAt time.Time) error {
	receiverTypes := h.supportedReceiverTypes
	if len(receiverTypes) == 0 {
		return errors.New("no receiver type plugin registered, skipping dequeue")
	} else {
		if err := h.q.Dequeue(ctx, receiverTypes, h.batchSize, h.MessageHandler); err != nil {
			if !errors.Is(err, ErrNoMessage) {
				return fmt.Errorf("dequeue failed on handler with id %s: %w", h.identifier, err)
			}
		}
	}
	return nil
}

func (h *Handler) errorMessageHandler(ctx context.Context, retryable bool, herr error, msg *Message) error {
	msg.MarkFailed(time.Now(), retryable, herr)
	if err := h.q.ErrorCallback(ctx, *msg); err != nil {
		return fmt.Errorf("failed to execute error callback with receiver type %s and error %w", msg.ReceiverType, err)
	}
	return herr
}

// MessageHandler is a function to handler dequeued message
func (h *Handler) MessageHandler(ctx context.Context, messages []Message) error {
	for _, msg := range messages {
		if err := h.SingleMessageHandler(ctx, &msg); err != nil {
			h.logger.Error(err.Error())
		}
	}
	return nil
}

func (h *Handler) SingleMessageHandler(ctx context.Context, msg *Message) error {

	defer func() {
		h.instrumentMQDuration(ctx, msg)
	}()

	notifier, err := h.getNotifierPlugin(msg.ReceiverType)
	if err != nil {
		return h.errorMessageHandler(ctx, false, err, msg)
	}

	msg.MarkPending(time.Now())

	newConfig, err := notifier.PostHookQueueTransformConfigs(ctx, msg.Configs)
	if err != nil {
		return h.errorMessageHandler(ctx, false, err, msg)
	}
	msg.Configs = newConfig

	if retryable, err := notifier.Send(ctx, *msg); err != nil {
		return h.errorMessageHandler(ctx, retryable, err, msg)
	}

	msg.MarkPublished(time.Now())

	if err := h.q.SuccessCallback(ctx, *msg); err != nil {
		return err
	}

	return nil
}

func (h *Handler) instrumentMQDuration(ctx context.Context, msg *Message) {
	h.metricHistMQDuration.Record(
		ctx, time.Since(msg.CreatedAt).Milliseconds(), metric.WithAttributes(
			attribute.String("receiver_type", msg.ReceiverType),
			attribute.String("status", string(msg.Status)),
			attribute.Int("try_count", msg.TryCount),
			attribute.Int("max_try", msg.MaxTries),
			attribute.Bool("retryable", msg.Retryable),
		),
	)
}
