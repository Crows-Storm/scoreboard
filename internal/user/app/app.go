package app

import (
	"github.com/Crows-Storm/scoreboard/internal/user/app/command"
	"github.com/Crows-Storm/scoreboard/internal/user/app/query"
)

// Application is a simple the CQRS
type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	RegisterUser command.RegisterUserHandler
	UpdateUser   command.UpdateUserHandler
}

type Queries struct {
	GetUser query.GetUserHandler
}
