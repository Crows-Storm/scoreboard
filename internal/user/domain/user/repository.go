package user

import (
	"context"
	"fmt"
)

type Repository interface {
	Create(context.Context, *User) (*User, error)
	Update(
		ctx context.Context,
		u *User,
		updateFun func(context.Context, *User) (*User, error),
	) error

	Get(ctx context.Context, id string) (*User, error)
}

type NotFoundError struct {
	UserId string
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("User Not Found: %s", n.UserId)
}
