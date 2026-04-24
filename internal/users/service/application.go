package service

import (
	"context"

	"github.com/Crows-Storm/scoreboard/internal/users/app"
)

func NewApplication(ctx context.Context) app.Application {
	return app.Application{}
}
