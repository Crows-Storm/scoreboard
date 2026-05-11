package command

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/common/config/decorator"
	domain "github.com/Crows-Storm/scoreboard/internal/user/domain/user"
	"github.com/sirupsen/logrus"
)

type RegisterUserCommand struct {
	Email    string
	Name     string
	Password string
	Avatar   string
}

type RegisterUserHandler decorator.CommandHandler[RegisterUserCommand, *domain.User]

type registerUserHandler struct {
	userRepo domain.Repository
}

func NewRegisterUserHandler(
	userRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) RegisterUserHandler {
	if userRepo == nil {
		panic("userRepo is nil !!!")
	}
	return decorator.ApplyCommandDecorator(
		registerUserHandler{userRepo: userRepo},
		logger,
		metricsClient,
	)
}

func (r registerUserHandler) Handle(ctx context.Context, cmd RegisterUserCommand) (user *domain.User, err error) {
	v, err := r.userRepo.Create(ctx, &domain.User{
		Email:    cmd.Email,
		Name:     cmd.Name,
		Password: cmd.Password,
		Avatar:   cmd.Avatar,
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}
