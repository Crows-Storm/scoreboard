package main

import (
	"github.com/Crows-Storm/scoreboard/internal/common/server"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	serviceName := viper.GetString("service-name")
	server.RunGRPCServer(serviceName, func(server *grpc.Server) {

	})
}
