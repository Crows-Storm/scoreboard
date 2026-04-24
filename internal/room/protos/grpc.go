package protos

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/common/genproto/roompb"
	"github.com/Crows-Storm/scoreboard/internal/room/app"
)

type GRPCServer struct {
	roompb.UnimplementedRoomServiceServer
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) CreatedRoom(ctx context.Context, request *roompb.CreatedRoomRequest) (*roompb.CreatedRoomResponse, error) {
	//TODO implement me
	// G.app
	panic("implement me")
}

func (G GRPCServer) JoinRoom(ctx context.Context, request *roompb.JoinRoomRequest) (*roompb.JoinedRoomResponse, error) {
	//TODO implement me
	panic("implement me")
}
