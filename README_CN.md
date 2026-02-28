# Scoreboard

一个基于微服务架构的记分板系统，专为打牌、麻将等休闲游戏场景设计，无需使用现金即可记录游戏盈亏。

[English Documentation](README.md)

## 项目简介

Scoreboard 帮助你在和朋友玩牌、打麻将或其他游戏时记录分数和余额。游戏过程中无需使用真实货币，游戏结束后根据记录的余额进行结算即可。

## 特性

- � 游戏房间创建和管理
- 👥 玩家加入/离开功能
- � 实时追踪每位玩家的余额
- 📊 分数记录和结算
- 📡 RESTful API 和 gRPC 支持
- 💾 基于 Redis 的数据持久化
- 🔧 从 OpenAPI 和 Protocol Buffers 自动生成 API 客户端

## 架构

本项目采用领域驱动设计（DDD）的微服务架构，具有清晰的限界上下文：

- **Room Service（房间服务）**: 管理房间限界上下文 - 房间创建、玩家加入和游戏状态
- **User Service（用户服务）**: 管理用户限界上下文 - 用户资料和认证（计划中）
- **Common Module（公共模块）**: 共享配置、客户端和工具
- **API Definitions（API 定义）**: OpenAPI 规范和 Protocol Buffer 定义

## 环境要求

- Go 1.25.6 或更高版本
- Redis 服务器
- Protocol Buffers 编译器 (`protoc`)
- 必需的 Go 工具：
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`
  - `oapi-codegen`

## 安装

1. 克隆仓库：
```bash
git clone https://github.com/Crows-Storm/scoreboard.git
cd scoreboard
```

2. 安装依赖：
```bash
cd internal/room
go mod download
```

3. 配置 Redis 连接：
编辑 `internal/common/config/global.yaml` 文件，配置你的 Redis 连接信息。

## 代码生成

从定义文件生成 API 客户端和服务端代码：

```bash
# 生成所有代码（protobuf + OpenAPI）
make gen

# 仅生成 protobuf 代码
make genproto

# 仅生成 OpenAPI 代码
make genopenapi

# 清理生成的代码
make clean
```

## 运行服务

1. 启动 Redis 服务器：
```bash
redis-server
```

2. 运行房间服务：
```bash
cd internal/room
go run main.go http.go
```

服务将在配置的端口上启动（默认：8081）。

## API 文档

### 房间服务 API

基础 URL: `http://localhost:8081/api/v1`

#### 创建房间
创建一个新的游戏房间，创建者成为房主。

```http
POST /rooms/create
Content-Type: application/json

{
  "name": "张三"
}
```

响应：
```json
{
  "id": "room-123456",
  "users": [
    {
      "id": "user-789",
      "name": "张三",
      "avatar": "https://example.com/avatar.jpg",
      "balance": 0
    }
  ]
}
```

#### 加入房间
加入一个已存在的游戏房间，每位玩家初始余额为 0。

```http
POST /rooms/{roomId}/join
Content-Type: application/json

{
  "name": "李四"
}
```

响应：
```json
{
  "id": "room-123456",
  "timestamp": 1709164800,
  "users": [
    {
      "id": "user-789",
      "name": "张三",
      "balance": 100
    },
    {
      "id": "user-790",
      "name": "李四",
      "balance": 0
    }
  ]
}
```

## 使用场景

非常适合：
- 🀄 和朋友打麻将
- 🃏 扑克之夜
- 🎲 桌游聚会
- 🎯 任何想要记录分数而不想处理现金的游戏

## 工作原理

1. 一位玩家创建房间并分享房间 ID
2. 其他玩家使用房间 ID 加入
3. 游戏过程中，余额实时更新
4. 游戏结束后，玩家根据最终余额进行结算

## 项目结构

```
.
├── api/                    # API 定义
│   ├── openapi/           # OpenAPI 规范
│   └── roompb/            # Protocol Buffer 定义
├── internal/              # 内部包
│   ├── common/            # 共享工具和配置
│   │   ├── client/        # 生成的 API 客户端
│   │   ├── config/        # 配置管理
│   │   ├── genproto/      # 生成的 protobuf 代码
│   │   └── server/        # HTTP 服务器工具
│   ├── room/              # 房间服务实现
│   └── users/             # 用户服务（未来）
├── pkg/                   # 公共包
│   └── client/            # 公共 API 客户端
├── scripts/               # 构建和生成脚本
└── Makefile              # 构建自动化
```

## 开发

### 添加新的 API

1. 在 `api/openapi/*.yml` 或 `api/*pb/*.proto` 中定义你的 API
2. 运行 `make gen` 生成代码
3. 在相应的服务中实现处理器

### 测试

```bash
cd internal/room
go test ./...
```

## 配置

配置通过 `internal/common/config/global.yaml` 管理：

```yaml
room:
  service-name: "room-service"
  port: 8081

redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0
```

## 贡献

1. Fork 本仓库
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m '添加某个很棒的特性'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。

## 联系方式

项目链接: [https://github.com/Crows-Storm/scoreboard](https://github.com/Crows-Storm/scoreboard)
