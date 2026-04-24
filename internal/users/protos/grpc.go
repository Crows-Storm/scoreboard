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

func (G GRPCServer) CreateUser(ctx context.Context, request *userspb.CreateUserRequest) (*userspb.CreateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) GetUser(ctx context.Context, request *userspb.GetUserRequest) (*userspb.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) UpdateUser(ctx context.Context, request *userspb.UpdateUserRequest) (*userspb.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) DeleteUser(ctx context.Context, request *userspb.DeleteUserRequest) (*userspb.DeleteUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
