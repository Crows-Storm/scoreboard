package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type queryLoggerDecorator[C, R any] struct {
	logger *logrus.Entry
	base   QueryHandler[C, R]
}

func (q queryLoggerDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	logger := q.logger.WithFields(logrus.Fields{
		"query":      generateActionName(cmd),
		"query_body": fmt.Sprintf("%#v", cmd),
	})
	logger.Debug("Executing: query")

	defer func() {
		if err == nil {
			logger.Info("Executing query successfully!!!")
		} else {
			logrus.Error("Failed to Executing query ERROR for: ", err)
		}
	}()
	return q.base.Handle(ctx, cmd)
}

type mutationsLoggerDecorator[C, R any] struct {
	logger *logrus.Entry
	base   CommandHandler[C, R]
}

func (q mutationsLoggerDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	logger := q.logger.WithFields(logrus.Fields{
		"command":      generateActionName(cmd),
		"command_body": fmt.Sprintf("%#v", cmd),
	})
	logger.Debug("Executing: mutations command")

	defer func() {
		if err == nil {
			logger.Info("Executing mutations command successfully!!!")
		} else {
			logrus.Error("Failed to Executing mutations command ERROR for: ", err)
		}
	}()
	return q.base.Handle(ctx, cmd)
}

func generateActionName(cmd any) string {
	return strings.Split(fmt.Sprintf("%T", cmd), ".")[1]
}
