package service

import (
	"context"
	"testing"
	"time"

	"github.com/Crows-Storm/scoreboard/internal/user/app"
	"github.com/Crows-Storm/scoreboard/internal/user/app/command"
	"github.com/Crows-Storm/scoreboard/internal/user/app/query"
	domain "github.com/Crows-Storm/scoreboard/internal/user/domain/user"
	"github.com/bytedance/gopkg/util/logger"
)

func TestApplicationAllCommandSuccessfully(t *testing.T) {
	ctx := context.Background()
	application := NewApplication(ctx)

	tests := []struct {
		name string
		job  func(*testing.T, app.Application) error
	}{
		{
			name: "Test Register User Command Handler",
			job: func(t *testing.T, app app.Application) error {
				if app.Commands.RegisterUser == nil {
					t.Fatal("RegisterUser handler is nil")
				}
				_, err := app.Commands.RegisterUser.Handle(t.Context(), command.RegisterUserCommand{
					Email:    "testCommandAllSuccessfully@test.com",
					Name:     "TestCommandAllSuccessfully",
					Password: "123456",
					Avatar:   "https://baidu.com/TestCommandAllSuccessfully.jpg",
				})
				if err != nil {
					return err
				}
				return nil
			},
		}, {
			name: "Test Update User Command Handler",
			job: func(t *testing.T, app app.Application) error {
				if app.Commands.RegisterUser == nil {
					t.Fatal("RegisterUser handler is nil")
				}
				body, _ := app.Queries.GetUser.Handle(t.Context(), query.GetUser{
					Id: "apple",
				})

				// req is a mock the request data, do something change now
				req := body
				req.Email = "sanderQiu@test.com"
				req.Avatar = "https://google.com/TestApplicationAllCommandSuccessfully.jpg"

				_, err := app.Commands.UpdateUser.Handle(t.Context(), command.UpdateUserCommand{
					User: req, // request body
					UpdateFun: func(ctx context.Context, u *domain.User) (*domain.User, error) {
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
				})
				if err != nil {
					return err
				}
				return nil
			},
		}, {
			name: "Testing GetUser By Username",
			job: func(t *testing.T, app app.Application) error {
				if v, err := app.Queries.GetUser.Handle(ctx, query.GetUser{
					Id:       "",
					Username: "apple",
				}); err == nil {
					logger.Debugf("Query result user: %#v", v)
				} else {
					t.Fatal("Get User by username is fail.")
				}
				return nil
			},
		},
	}

	//for _, tt := range tests {
	//	logger.Debugf("Runing %#v Now", tt.name)
	//	t.Run(tt.name, func(t *testing.T) {
	//		if err := tt.job(t, application); err != nil {
	//			t.Fail()
	//		}
	//	})
	//}

	create := tests[0]
	t.Run(create.name, func(t *testing.T) {
		if err := create.job(t, application); err != nil {
			t.Fail()
		}
	})

	update := tests[1]
	t.Run(update.name, func(t *testing.T) {
		if err := update.job(t, application); err != nil {
			t.Fail()
		}
	})

}
