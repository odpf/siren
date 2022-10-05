package notification

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/odpf/salt/log"
	"github.com/odpf/siren/pkg/errors"
)

// Worker is a notification worker instance that runs one or more than one
// notification handler
type Worker struct {
	id      string
	logger  log.Logger
	handler *Handler
}

// NewWorker creates a new worker instance
func NewWorker(logger log.Logger, h *Handler) *Worker {
	return &Worker{
		id:      uuid.NewString(),
		logger:  logger,
		handler: h,
	}
}

// Run will execute and run one handler as goroutines
func (w *Worker) Run(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	w.logger.Info("running handler worker", "id", w.id)
	w.handler.RunHandler(ctx)
	w.logger.Info("handler worker exited", "id", w.id)

	<-ctx.Done()
	w.logger.Debug("worker context done", "err", ctx.Err(), "id", w.id)

	err := ctx.Err()
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return nil
	}
	return err
}
