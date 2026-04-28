package protos

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/common/genproto/userspb"
	"github.com/Crows-Storm/scoreboard/internal/users/app"
)

type GRPCServer struct {
	userspb.UnimplementedUserServiceServer
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) GetUser(ctx context.Context, request *userspb.GetUserRequest) (*userspb.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
