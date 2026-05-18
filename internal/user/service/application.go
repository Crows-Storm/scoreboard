package service

import (
	"context"

	grpcClient "github.com/Crows-Storm/scoreboard/internal/common/client"
	"github.com/Crows-Storm/scoreboard/internal/common/metrics"
	"github.com/Crows-Storm/scoreboard/internal/user/app"
	"github.com/Crows-Storm/scoreboard/internal/user/app/command"
	"github.com/Crows-Storm/scoreboard/internal/user/app/query"
	"github.com/Crows-Storm/scoreboard/internal/user/domain/adapters"
	"github.com/Crows-Storm/scoreboard/internal/user/domain/adapters/grpc"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	client, closeRoomGRPCClient, err := grpcClient.NewRoomGRPCClient(ctx)
	if err != nil {
		panic("RoomGRPC Init ERROR!!!")
	}
	roomGRPC := grpc.NewRoomGRPC(client)
	return newApplication(ctx, roomGRPC), func() {
		_ = closeRoomGRPCClient()
	}
}

func newApplication(ctx context.Context, roomGRPC *grpc.RoomGRPC) app.Application {
	userRepository := adapters.NewMemoryUserRepository() // can change repository
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TodoMetrics{}

	return app.Application{
		Commands: app.Commands{
			RegisterUser: command.NewRegisterUserHandler(userRepository, logger, metricsClient),
			UpdateUser:   command.NewUpdateUserHandler(userRepository, logger, metricsClient),
		},
		Queries: app.Queries{
			GetUser: query.NewGetUserHandler(userRepository, roomGRPC, logger, metricsClient),
		},
	}
}
