# Scoreboard

A microservices-based scoreboard system for card games, mahjong, and other casual gaming scenarios where you want to track wins/losses without using real cash.

[中文文档](README_CN.md)

## Overview

Scoreboard helps you keep track of game scores and balances when playing card games, mahjong, or any other games with friends. Instead of dealing with real money during the game, players can settle up afterwards based on the recorded balances.

## Features

- 🎮 Game room creation and management
- 👥 Player join/leave functionality
- 💰 Real-time balance tracking for each player
- 📊 Score recording and settlement
- 📡 RESTful API and gRPC support
- 💾 Redis-based data persistence
- 🔧 Auto-generated API clients from OpenAPI and Protocol Buffers

## Architecture

This project follows a Domain-Driven Design (DDD) microservices architecture with clear bounded contexts:

- **Room Service**: Manages the room bounded context - room creation, player joining, and game state
- **User Service**: Manages the user bounded context - user profiles and authentication (planned)
- **Common Module**: Shared configurations, clients, and utilities
- **API Definitions**: OpenAPI specs and Protocol Buffer definitions

## Prerequisites

- Go 1.25.6 or higher
- Redis server
- Protocol Buffers compiler (`protoc`)
- Required Go tools:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`
  - `oapi-codegen`

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Crows-Storm/scoreboard.git
cd scoreboard
```

2. Install dependencies:
```bash
cd internal/room
go mod download
```

3. Configure Redis connection:

Edit `internal/common/config/global.yaml` with your Redis configuration.

## Code Generation

Generate API clients and server stubs from definitions:

```bash
# Generate all code (protobuf + OpenAPI)
make gen

# Generate protobuf code only
make genproto

# Generate OpenAPI code only
make genopenapi

# Clean generated code
make clean
```

## Running the Service

1. Start Redis server:
```bash
redis-server
```

2. Run the room service:
```bash
cd internal/room
go run main.go http.go
```

The service will start on the configured port (default: 8081).

## API Documentation

### Room Service API

Base URL: `http://localhost:8081/api/v1`

#### Create Room
Create a new game room. The creator becomes the room master.

```http
POST /rooms/create
Content-Type: application/json

{
  "name": "John Doe"
}
```

Response:
```json
{
  "id": "room-123456",
  "users": [
    {
      "id": "user-789",
      "name": "John Doe",
      "avatar": "https://example.com/avatar.jpg",
      "balance": 0
    }
  ]
}
```

#### Join Room
Join an existing game room. Each player starts with a balance of 0.

```http
POST /rooms/{roomId}/join
Content-Type: application/json

{
  "name": "Jane Smith"
}
```

Response:
```json
{
  "id": "room-123456",
  "timestamp": 1709164800,
  "users": [
    {
      "id": "user-789",
      "name": "John Doe",
      "balance": 100
    },
    {
      "id": "user-790",
      "name": "Jane Smith",
      "balance": 0
    }
  ]
}
```

## Use Cases

Perfect for:
- 🀄 Mahjong games with friends
- 🃏 Poker nights
- 🎲 Board game sessions
- 🎯 Any game where you want to track scores without handling cash

## How It Works

1. One player creates a room and shares the room ID
2. Other players join using the room ID
3. During the game, balances are updated in real-time
4. After the game, players settle up based on their final balances

## Project Structure

```
.
├── api/                    # API definitions
│   ├── openapi/           # OpenAPI specifications
│   └── roompb/            # Protocol Buffer definitions
├── internal/              # Internal packages
│   ├── common/            # Shared utilities and configurations
│   │   ├── client/        # Generated API clients
│   │   ├── config/        # Configuration management
│   │   ├── genproto/      # Generated protobuf code
│   │   └── server/        # HTTP server utilities
│   ├── room/              # Room service implementation
│   └── users/             # User service (future)
├── pkg/                   # Public packages
│   └── client/            # Public API clients
├── scripts/               # Build and generation scripts
└── Makefile              # Build automation
```

## Development

### Adding New APIs

1. Define your API in `api/openapi/*.yml` or `api/*pb/*.proto`
2. Run `make gen` to generate code
3. Implement the handlers in the respective service

### Testing

```bash
cd internal/room
go test ./...
```

## Configuration

Configuration is managed through `internal/common/config/global.yaml`:

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

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Project Link: [https://github.com/Crows-Storm/scoreboard](https://github.com/Crows-Storm/scoreboard)
