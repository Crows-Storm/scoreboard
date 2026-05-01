package service

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/user/app"
	"github.com/Crows-Storm/scoreboard/internal/user/domain/adapters"
	domain "github.com/Crows-Storm/scoreboard/internal/user/domain/user"
)

func NewApplication(ctx context.Context) app.Application {
	userRepository := adapters.NewMemoryUserRepository()
	return app.Application{
		Commands: struct{ Repo domain.Repository }{Repo: userRepository}, // injection write repo
		Queries:  struct{ Repo domain.Repository }{Repo: userRepository}, // injection read repo
	}
}
