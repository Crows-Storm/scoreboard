package decorator

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type MetricsClient interface {
	Inc(key string, value int)
}

type queryMetricsDecorator[C, R any] struct {
	base   QueryHandler[C, R]
	client MetricsClient
}

func (q queryMetricsDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	startTime := time.Now()
	actionName := strings.ToLower(generateActionName(cmd))

	defer func() {
		endTime := time.Since(startTime)
		q.client.Inc(fmt.Sprintf("querys %s.duration", actionName), int(endTime.Seconds()))

		if err == nil {
			q.client.Inc(fmt.Sprintf("querys %s.success", actionName), 1)
		} else {
			q.client.Inc(fmt.Sprintf("querys %s.failure", actionName), 1)
		}
	}()
	return q.base.Handle(ctx, cmd)
}
