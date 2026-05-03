package app

import "github.com/Crows-Storm/scoreboard/internal/user/app/query"

// Application is a simple the CQRS
type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	GetUser query.GetUserHandler
}
