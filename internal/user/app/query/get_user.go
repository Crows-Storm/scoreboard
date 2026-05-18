package query

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/common/config/decorator"
	domain "github.com/Crows-Storm/scoreboard/internal/user/domain/user"
	"github.com/sirupsen/logrus"
)

type GetUser struct {
	Id       string
	Username string
}

// GetUserHandler == type UserId int, so GetUserHandler interface is a QueryHandler type
type GetUserHandler decorator.QueryHandler[GetUser, *domain.User]

type getUserHandler struct {
	userRepo domain.Repository
	roomGRPC RoomService
}

func NewGetUserHandler(
	userRepo domain.Repository,
	roomGRPC RoomService,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) GetUserHandler {
	if userRepo == nil {
		panic("userRepo is nil !!!")
	}
	return decorator.ApplyQueryDecorator[GetUser, *domain.User](
		getUserHandler{
			userRepo: userRepo,
			roomGRPC: roomGRPC,
		},
		logger,
		metricsClient)
}

func (g getUserHandler) Handle(ctx context.Context, query GetUser) (*domain.User, error) {
	switch {
	case query.Id != "":
		v, err := g.roomGRPC.InTheRoom(ctx, query.Id)
		if err != nil {
			return nil, err
		}
		v = v
		return g.userRepo.Get(ctx, query.Id)
	case query.Username != "":
		return g.userRepo.GetByUsername(ctx, query.Username)
	default:
		return nil, domain.NotFoundError{UserId: query.Id}
	}
}
