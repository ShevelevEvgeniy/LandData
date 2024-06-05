package retry_func

import (
	"context"
	"github.com/ShevelevEvgeniy/app/config"
	"github.com/avast/retry-go"
	"log/slog"
	"time"
)

type RetryFunc struct {
	cfg      *config.RetryConfig
	log      *slog.Logger
	Attempts uint
	Delay    time.Duration
}

func NewRetryFunc(cfg *config.RetryConfig, log *slog.Logger) *RetryFunc {
	return &RetryFunc{
		cfg:      cfg,
		log:      log,
		Attempts: cfg.Attempts,
		Delay:    cfg.Delay,
	}
}

func (r *RetryFunc) Do(ctx context.Context, task func() error) error {
	return retry.Do(
		task,
		retry.Attempts(r.Attempts),
		retry.DelayType(func(n uint, err error, config *retry.Config) time.Duration {
			return time.Duration(n) * r.Delay
		}),
		retry.OnRetry(func(n uint, err error) {
			r.log.Info("Retrying", slog.Int("attempt", int(n)), slog.String("error", err.Error()))
		}),
		retry.Context(ctx),
	)
}
