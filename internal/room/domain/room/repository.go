package room

import (
	"context"
	"fmt"
)

type Repository interface {
	Create(context.Context, *Room) error // return nothing & error

	Update(
		ctx context.Context,
		r *Room,
		updateFun func(context.Context, *Room) (*Room, error), // Anonymous function
	) error

	Get(ctx context.Context, id, roomId string) (*Room, error) // detail
	List(ctx context.Context) ([]Room, error)                  // room list

}

// NotFoundError is a error
type NotFoundError struct {
	RoomId string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("room not found by: %s", e.RoomId)
}
