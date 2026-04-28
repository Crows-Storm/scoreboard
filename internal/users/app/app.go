package app

import (
	"context"

	domain "github.com/Crows-Storm/scoreboard/internal/users/domain/users"
)

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

func (c Commands) CreateUser(ctx context.Context, user *domain.Users) (*domain.Users, error) {
	return c.Repo.Create(ctx, user)
}

func (c Commands) UpdateUser(ctx context.Context, user *domain.Users, updateFn func(context.Context, *domain.Users) (*domain.Users, error)) error {
	return c.Repo.Update(ctx, user, updateFn)
}

func (q Queries) GetUser(ctx context.Context, id string) (*domain.Users, error) {
	return q.Repo.Get(id)
}
