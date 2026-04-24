package main

import (
	"github.com/Crows-Storm/scoreboard/internal/room/app"
	"github.com/gin-gonic/gin"
)

// HTTPServer to implement HTTPServer interface
type HTTPServer struct {
	app app.Application
}

func (H HTTPServer) CreateRoom(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (H HTTPServer) JoinRoom(c *gin.Context, roomId string) {
	//TODO implement me
	panic("implement me")
}
