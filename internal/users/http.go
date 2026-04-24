package main

import (
	"github.com/Crows-Storm/scoreboard/internal/users/app"
	"github.com/gin-gonic/gin"
)

// HTTPServer to implement HTTPServer interface
type HTTPServer struct {
	app app.Application
}

func (H HTTPServer) CreateUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (H HTTPServer) DeleteUser(c *gin.Context, userId string) {
	//TODO implement me
	panic("implement me")
}

func (H HTTPServer) GetUser(c *gin.Context, userId string) {
	//TODO implement me
	panic("implement me")
}

func (H HTTPServer) UpdateUser(c *gin.Context, userId string) {
	//TODO implement me
	panic("implement me")
}
