package grpc

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/common/genproto/roompb"
	"github.com/sirupsen/logrus"
)

type RoomGRPC struct {
	client roompb.RoomServiceClient
}

func NewRoomGRPC(client roompb.RoomServiceClient) *RoomGRPC {
	return &RoomGRPC{client: client}
}

// InTheRoom is a GRPC caller
func (r RoomGRPC) InTheRoom(ctx context.Context, id string) (bool, error) {
	resp, err := r.client.InTheRoom(
		ctx,
		&roompb.InTheRoomRequest{
			UserId: id,
		},
	)

	if err != nil {
		logrus.Debugf("InTheRoom GRPC call fila: %#v", id)
		panic(err)
	}
	logrus.Infof("InTheRoom GRPC call Successfully: %#v", id)
	return resp.InTheRoom, nil
}
