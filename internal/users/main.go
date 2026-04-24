package main

import (
	"context"
	"log"

	"github.com/Crows-Storm/scoreboard/internal/common/config"
	"github.com/Crows-Storm/scoreboard/internal/common/genproto/userspb"
	"github.com/Crows-Storm/scoreboard/internal/common/server"
	"github.com/Crows-Storm/scoreboard/internal/users/protos"
	"github.com/Crows-Storm/scoreboard/internal/users/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}

	err := config.NewRedisInstance()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("users.service-name")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)

	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		userServiceServer := protos.NewGRPCServer(application)
		userspb.RegisterUserServiceServer(server, userServiceServer)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		protos.RegisterHandlersWithOptions(router, HTTPServer{
			app: application, // inject application
		}, protos.GinServerOptions{
			BaseURL:      "/api/v1",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})

	log.Println(serviceName)
	defer func() {
		if err := config.RedisClient.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
