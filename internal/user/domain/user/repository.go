package user

import (
	"context"
	"fmt"
)

type Repository interface {
	Create(context.Context, *Users) (*Users, error)
	Update(
		ctx context.Context,
		u *Users,
		updateFun func(context.Context, *Users) (*Users, error),
	) error

	Get(id string) (*Users, error)
}

type NotFoundError struct {
	UserId string
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("User Not Found: %s", n.UserId)
}
