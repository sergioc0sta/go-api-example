# Go API Example

A RESTful API built with Go that provides product management with JWT authentication. This project demonstrates clean architecture principles, comprehensive testing, and API documentation with Swagger.

## Features

- **Product Management**: Full CRUD operations for products
- **User Authentication**: JWT-based authentication system
- **API Documentation**: Auto-generated Swagger documentation
- **Database Integration**: GORM with SQLite for development
- **Comprehensive Testing**: Unit tests for all major components
- **Clean Architecture**: Organized code structure with separation of concerns

## Tech Stack

- **Language**: Go 1.24
- **Web Framework**: Chi Router
- **Database ORM**: GORM
- **Database**: SQLite (development)
- **Authentication**: JWT with go-chi/jwtauth
- **Documentation**: Swagger/OpenAPI
- **Testing**: Testify
- **Configuration**: Viper

## Project Structure
go-api-example/
├── cmd/server/           # Application entrypoint
│   ├── main.go
│   └── .env
├── configs/              # Configuration management
├── docs/                 # Auto-generated Swagger docs
├── internal/
│   ├── dto/              # Data Transfer Objects
│   ├── entity/           # Business entities
│   └── infra/
│       ├── database/     # Database layer
│       └── webserver/    # HTTP handlers
└── pkg/entity/           # Shared utilities
