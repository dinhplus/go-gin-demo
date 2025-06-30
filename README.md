# My Gin App

A simple REST API using Go and Gin.

## Getting Started

1. Install dependencies:
   ```sh
   go mod tidy
   ```

2. Run the server:
   ```sh
   go run ./cmd/main.go
   ```

3. Test endpoints:
   - `GET /ping` → `{ "message": "pong from handler" }`
   - `GET /users` → List of users
   - `GET /departments` → List of departments
   - `GET /stacks` → List of stacks
   - `GET /positions` → List of positions
   - `POST /users`, `POST /departments`, `POST /stacks`, `POST /positions` → Create resources

4. API Documentation (Swagger UI):
   - Visit [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Project Structure

- `cmd/main.go` - Entry point
- `internal/handler` - HTTP handlers & routes
- `internal/service` - Business logic
- `internal/repository` - Data access & DB seeding
- `internal/model` - Data models
- `internal/client` - 3rd-party/external service clients
- `docs/` - Auto-generated Swagger docs

    ```
    my-gin-app
    ├── .env
    ├── .gitignore
    ├── README.md
    ├── cmd
    │   └── main.go
    ├── docs
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── go.mod
    ├── go.sum
    └── internal
        ├── client
        │   └── thirdparty_client.go
        ├── handler
        │   ├── department_handler.go
        │   ├── handler.go
        │   ├── position_handler.go
        │   ├── routes.go
        │   ├── stack_handler.go
        │   └── user_handler.go
        ├── model
        │   ├── department.go
        │   ├── model.go
        │   ├── position.go
        │   ├── stack.go
        │   └── user.go
        ├── repository
        │   ├── db.go
        │   ├── department_repository.go
        │   ├── position_repository.go
        │   ├── repository.go
        │   ├── seed.go
        │   ├── stack_repository.go
        │   └── user_repository.go
        └── service
            ├── department_service.go
            ├── position_service.go
            ├── service.go
            ├── stack_service.go
            └── user_service.go

    ```
## 3rd-Party Service Integration

- Add your 3rd-party service logic in `internal/client/thirdparty_client.go`.
- Use this client in your services or handlers as needed.

## OpenAPI/Swagger

- Annotate your handlers with Swagger comments.
- Regenerate docs with:
  ```sh
  swag init --generalInfo cmd/main.go --output docs
  ```
- Access docs at `/swagger/index.html`.
