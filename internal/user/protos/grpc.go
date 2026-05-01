package protos

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/common/genproto/userpb"
	"github.com/Crows-Storm/scoreboard/internal/user/app"
)

type GRPCServer struct {
	userpb.UnimplementedUserServiceServer
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) GetUser(ctx context.Context, request *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	v, err := G.app.Queries.GetUser(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	v = v
	return nil, nil
}
