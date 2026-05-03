package service

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/common/metrics"
	"github.com/Crows-Storm/scoreboard/internal/user/app"
	"github.com/Crows-Storm/scoreboard/internal/user/app/query"
	"github.com/Crows-Storm/scoreboard/internal/user/domain/adapters"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	userRepository := adapters.NewMemoryUserRepository() // can change repository
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetUser: query.NewGetUserHandler(userRepository, logger, metricsClient),
		},
	}
}
