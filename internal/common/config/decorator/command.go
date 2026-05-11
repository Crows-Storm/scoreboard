package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

type CommandHandler[C, R any] interface {
	Handle(ctx context.Context, cmd C) (result R, err error)
}

func ApplyCommandDecorator[C, R any](handler CommandHandler[C, R], logger *logrus.Entry, metricsClient MetricsClient) CommandHandler[C, R] {
	return mutationsLoggerDecorator[C, R]{
		logger: logger,
		base: mutationsMetricsDecorator[C, R]{
			base:   handler,
			client: metricsClient,
		},
	}
}
