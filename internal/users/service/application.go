package service

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/users/app"
	"github.com/Crows-Storm/scoreboard/internal/users/domain/adapters"
	domain "github.com/Crows-Storm/scoreboard/internal/users/domain/users"
)

func NewApplication(ctx context.Context) app.Application {
	userRepository := adapters.NewMemoryUserRepository()
	return app.Application{
		Commands: struct{ Repo domain.Repository }{Repo: userRepository}, // injection write repo
		Queries:  struct{ Repo domain.Repository }{Repo: userRepository}, // injection read repo
	}
}
