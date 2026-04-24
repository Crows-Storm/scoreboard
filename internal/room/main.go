package main

import (
	"context"
	"log"

	"github.com/Crows-Storm/scoreboard/internal/common/config"
	"github.com/Crows-Storm/scoreboard/internal/common/genproto/roompb"
	"github.com/Crows-Storm/scoreboard/internal/common/server"
	"github.com/Crows-Storm/scoreboard/internal/room/protos"
	"github.com/Crows-Storm/scoreboard/internal/room/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
	err := config.NewRedisInstance() // create a redis instance and connect it!
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	//log.Println(viper.Get("room"))

	// test connect and query a key
	//values := config.RedisClient.Get(config.RedisCtx, "TRADING_PAIR:INDEPENDENT:PATTERN:ACTIVE:8ec2506c-32a4-4471-b6a2-8d309ca5c582:D1:MOMENTUM")
	//
	//result, err := values.Result()
	//if err != nil {
	//	log.Fatal("[Redis Query ERROR] ", err)
	//}
	//log.Println("[Query Successfully] Value: ", result)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serviceName := viper.GetString("room.service-name")

	// new a application form service
	application := service.NewApplication(ctx)

	// TODO need init NewGRPCServer
	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := protos.NewGRPCServer(application)
		roompb.RegisterRoomServiceServer(server, svc)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		protos.RegisterHandlersWithOptions(router, HTTPServer{
			app: application,
		}, protos.GinServerOptions{
			BaseURL:      "/api/v1",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})

	defer func() {
		if err := config.RedisClient.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
