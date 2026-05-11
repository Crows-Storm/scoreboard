package main

import (
	"context"
	"net/http"
	"time"

	"github.com/Crows-Storm/scoreboard/internal/user/app"
	"github.com/Crows-Storm/scoreboard/internal/user/app/command"
	"github.com/Crows-Storm/scoreboard/internal/user/app/query"
	domain "github.com/Crows-Storm/scoreboard/internal/user/domain/user"
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
	var req command.RegisterUserCommand
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	v, err := H.app.Commands.RegisterUser.Handle(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed!!", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully!!", "data": v})
}

func (H HTTPServer) GetUser(c *gin.Context, userId string) {
	v, err := H.app.Queries.GetUser.Handle(c, query.GetUser{
		Id: userId,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Failed!!", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully!!", "data": v})
	}
}

func (H HTTPServer) UpdateUser(c *gin.Context, userId string) {
	var body domain.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}
	updateCmd := command.UpdateUserCommand{
		User: &body,
		UpdateFun: func(ctx context.Context, u *domain.User) (*domain.User, error) {
			u.Id = userId
			u.UpdatedAt = time.Now().UnixMilli()
			if &body != nil {
				if body.Name != "" {
					u.Name = body.Name
				}
				if body.Email != "" {
					u.Email = body.Email
				}
				if body.Avatar != "" {
					u.Avatar = body.Avatar
				}
			}
			return u, nil
		},
	}

	_, err := H.app.Commands.UpdateUser.Handle(c, updateCmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Update failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully!!"})
}
