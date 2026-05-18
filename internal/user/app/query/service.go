package query

import "context"

type RoomService interface {
	InTheRoom(ctx context.Context, id string) (bool, error)
}
