package command

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/common/config/decorator"
	domain "github.com/Crows-Storm/scoreboard/internal/user/domain/user"
	"github.com/sirupsen/logrus"
)

type Void struct{}

type UpdateUserCommand struct {
	User      *domain.User
	UpdateFun func(context.Context, *domain.User) (*domain.User, error)
}

// UpdateUserHandler just update user the basic info, cannot update password
type UpdateUserHandler decorator.CommandHandler[UpdateUserCommand, Void]

type updateUserHandler struct {
	userRepo domain.Repository
}

func NewUpdateUserHandler(
	userRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) UpdateUserHandler {
	if userRepo == nil {
		panic("userRepo is nil !!!")
	}
	return decorator.ApplyCommandDecorator[UpdateUserCommand, Void](
		updateUserHandler{
			userRepo: userRepo,
		},
		logger,
		metricsClient,
	)
}

func (u updateUserHandler) Handle(ctx context.Context, cmd UpdateUserCommand) (Void, error) {
	if err := u.userRepo.Update(ctx, cmd.User, cmd.UpdateFun); err != nil {
		return Void{}, err
	}
	return Void{}, nil
}
