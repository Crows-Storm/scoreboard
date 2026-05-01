package app

import (
	"context"

	domain "github.com/Crows-Storm/scoreboard/internal/user/domain/user"
)

// Application is a simple the CQRS
type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Repo domain.Repository // write repo
}

type Queries struct {
	Repo domain.Repository // read repo
}

func (c Commands) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return c.Repo.Create(ctx, user)
}

func (c Commands) UpdateUser(ctx context.Context, user *domain.User, updateFn func(context.Context, *domain.User) (*domain.User, error)) error {
	return c.Repo.Update(ctx, user, updateFn)
}

func (q Queries) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return q.Repo.Get(id)
}
