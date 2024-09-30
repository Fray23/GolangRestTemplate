## Overview

This is a skeleton application built with Go, utilizing the following libraries:

- **GORM**: An ORM library for Go, used for database interactions.
- **Gin**: A web framework for building web applications and APIs.
- **Goose**: A database migration tool for managing schema changes.


## Project Structure
```
├── api                     # api
│   ├── app.go              # register router group
│   └── auth                # router for auth/signup
│       ├── login.go
│       ├── routers.go
│       └── signup.go
├── cmd
│   └── main.go             # Entry point of the application.
├── core
│   ├── db
│   │   └── db_config.go
│   ├── dto
│   │   ├── api             # DTOs for the API.
│   │   │   └── auth.go
│   │   └── repository.go   # Common DTOs for repository interactions.
│   ├── enum
│   │   └── auth_errors.go
│   ├── error
│   │   └── error.go
│   ├── middleware
│   │   └── error_wrapper.go
│   └── security
│       └── hash.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── migrations
├── models                  # Folder for database models.
│   └── user.go
├── repository              # Folder for data access logic.
│   └── user.go
└── service                 # contains the business logic of the application
    └── auth.go

```

### RUN migrations
```bash
goose postgres "postgres://postgres_username:password@localhost:5432/your_database_name?sslmode=disable" up
```

### RUN project
```bash
docker compose up -d --build
docker compose exec backend bash
go run cmd/main.go
```
