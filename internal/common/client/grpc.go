package client

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/common/genproto/roompb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewRoomGRPCClient(ctx context.Context) (client roompb.RoomServiceClient, close func() error, err error) {
	grpcAddr := viper.GetString("room.grpc-addr")
	v, err := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, nil
	}

	conn, err := grpc.NewClient(grpcAddr, v...)
	if err != nil {
		return nil, func() error { return nil }, nil
	}
	return roompb.NewRoomServiceClient(conn), conn.Close, nil

}

func grpcDialOpts(addr string) ([]grpc.DialOption, error) {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}, nil
}
