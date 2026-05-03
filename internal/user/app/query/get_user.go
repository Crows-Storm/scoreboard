package query

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/common/config/decorator"
	domain "github.com/Crows-Storm/scoreboard/internal/user/domain/user"
	"github.com/sirupsen/logrus"
)

type GetUser struct {
	Id string
}

// GetUserHandler == type UserId int, so GetUserHandler interface is a QueryHandler type
type GetUserHandler decorator.QueryHandler[GetUser, *domain.User]

type getUserHandler struct {
	userRepo domain.Repository
}

func NewGetUserHandler(
	userRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) GetUserHandler {
	if userRepo == nil {
		panic("userRepo is nil !!!")
	}
	return decorator.ApplyQueryDecorator[GetUser, *domain.User](
		getUserHandler{userRepo: userRepo},
		logger,
		metricsClient)
}

func (g getUserHandler) Handle(ctx context.Context, query GetUser) (*domain.User, error) {
	v, err := g.userRepo.Get(ctx, query.Id)
	if err != nil {
		return nil, err
	}
	return v, nil
}
