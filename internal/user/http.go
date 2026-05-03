package main

import (
	"net/http"

	"github.com/Crows-Storm/scoreboard/internal/user/app"
	"github.com/Crows-Storm/scoreboard/internal/user/app/query"
	"github.com/gin-gonic/gin"
)

// HTTPServer to implement HTTPServer interface
type HTTPServer struct {
	app app.Application
}

func (H HTTPServer) LoginUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (H HTTPServer) RegisterUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (H HTTPServer) GetUser(c *gin.Context, userId string) {
	v, err := H.app.Queries.GetUser.Handle(c, query.GetUser{
		//Id: userId,
		Id: "apple",
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Failed!!", "error": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully!!", "data": v})
}

func (H HTTPServer) UpdateUser(c *gin.Context, userId string) {
	//TODO implement me
	panic("implement me")
}
