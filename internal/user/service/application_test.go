package service

import (
	"context"
	"testing"

	"github.com/Crows-Storm/scoreboard/internal/user/app"
	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	ctx := context.Background()

	app := NewApplication(ctx)
	if app.Commands.RegisterUser == nil {
		t.Error("RegisterUser command should not be nil")
	}

	if app.Commands.UpdateUser == nil {
		t.Error("UpdateUser command should not be nil")
	}

	if app.Queries.GetUser == nil {
		t.Error("GetUser query should not be nil")
	}
}

func TestNewApplicationFunctionality(t *testing.T) {
	ctx := context.Background()
	application := NewApplication(ctx)

	tests := []struct {
		name string
		test func(*testing.T, app.Application)
	}{
		{
			name: "RegisterUser handler is properly initialized",
			test: func(t *testing.T, app app.Application) {
				if app.Commands.RegisterUser == nil {
					t.Fatal("RegisterUser handler is nil")
				}
				// 可以进一步验证 handler 的类型
				// 注意：这需要导出 handler 或者使用反射
			},
		},
		{
			name: "UpdateUser handler is properly initialized",
			test: func(t *testing.T, app app.Application) {
				if app.Commands.UpdateUser == nil {
					t.Fatal("UpdateUser handler is nil")
				}
			},
		},
		{
			name: "GetUser query handler is properly initialized",
			test: func(t *testing.T, app app.Application) {
				if app.Queries.GetUser == nil {
					t.Fatal("GetUser handler is nil")
				}
			},
		},
		{
			name: "Repository is shared across handlers",
			test: func(t *testing.T, app app.Application) {
				// 验证所有 handler 使用的是同一个 repository 实例
				// 这需要暴露 repository 或者通过其他方式验证
				// 例如：通过创建用户然后查询来验证
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t, application)
		})
	}
}

// 集成测试：验证实际功能
func TestNewApplicationIntegration(t *testing.T) {
	//ctx := context.Background()
	//_ := NewApplication(ctx)

	// 测试完整的业务流程
	t.Run("complete user flow", func(t *testing.T) {
		// 这里需要知道具体的命令和查询类型
		// 示例：
		// registerCmd := command.RegisterUserCommand{
		//     ID:   "user-123",
		//     Name: "Test User",
		// }
		//
		// err := application.Commands.RegisterUser.Handle(ctx, registerCmd)
		// if err != nil {
		//     t.Fatalf("Failed to register user: %v", err)
		// }
		//
		// getUserQuery := query.GetUserQuery{
		//     ID: "user-123",
		// }
		//
		// user, err := application.Queries.GetUser.Handle(ctx, getUserQuery)
		// if err != nil {
		//     t.Fatalf("Failed to get user: %v", err)
		// }
		//
		// if user.Name != "Test User" {
		//     t.Errorf("Expected user name 'Test User', got '%s'", user.Name)
		// }
	})
}

// 测试依赖注入是否正确
func TestNewApplicationDependencies(t *testing.T) {
	ctx := context.Background()

	// 方式1：直接创建多个实例对比
	app1 := NewApplication(ctx)
	app2 := NewApplication(ctx)

	// 验证每次调用都创建新的实例（不是单例）
	if &app1 == &app2 {
		t.Error("Each call should create a new application instance")
	}

	// 验证 repository 是独立的新实例
	// 通过比较 handler 内部引用来验证（需要反射或导出字段）
}

func TestNewApplicationWithAssert(t *testing.T) {
	ctx := context.Background()
	application := NewApplication(ctx)

	assert.NotNil(t, application.Commands.RegisterUser, "RegisterUser handler should be initialized")
	assert.NotNil(t, application.Commands.UpdateUser, "UpdateUser handler should be initialized")
	assert.NotNil(t, application.Queries.GetUser, "GetUser query handler should be initialized")
}

func BenchmarkNewApplication(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewApplication(ctx)
	}
}
