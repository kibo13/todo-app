## Todo App

This is a RESTful API application for managing a todo list, built using the Go programming language. The application follows a clean architecture approach and interacts with a PostgreSQL database for data storage. It also includes features for user registration, authentication, and comprehensive API documentation using Swagger.

### Features

- **REST API**: A fully functional REST API for managing todo lists and items, including CRUD (Create, Read, Update, Delete) operations.
- **PostgreSQL Database**: Integrated with PostgreSQL for persistent data storage. The app utilizes SQL queries for managing data interactions.
- **SQL Queries**: The app demonstrates the use of SQL queries to interact with the database (select, insert, update, delete operations).
- **Clean Architecture**: Implements the principles of clean architecture, separating concerns into different layers for better scalability and maintainability.
- **User Registration & Authentication**: Provides endpoints for user registration and login, using JWT (JSON Web Tokens) for authentication and session management.
- **Swagger Documentation**: API documentation is auto-generated and accessible via Swagger UI for easy exploration and testing of the API.

### Getting Started

#### Prerequisites

- [Docker](https://www.docker.com/get-started) for running the application in containers.
- [Go](https://golang.org/dl/) installed on your local machine (if you prefer to run the app without Docker).
- A PostgreSQL database (Dockerized or locally installed) to store data.

#### Setup and Installation

1. Clone the repository:

```bash
git clone https://github.com/kibo13/todo-app.git
cd todo-app
```

2. **Using Docker**: build and start the containers using Docker Compose.

```sh
docker-compose up --build
```

This will:

- Build the Go application.
- Set up a PostgreSQL container.
- Run the application inside a Docker container.

**Note**: Make sure to adjust the database settings in the `.env` file (if using local DB) or `docker-compose.yml` for Dockerized PostgreSQL.

3. **Without Docker**: If you want to run the application directly without Docker, follow these steps:

```sh
go mod download
```

- Set up the PostgreSQL database:
  - Create a database for the app (e.g., `todo_db`).
  - Update the `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, and `DB_PASSWORD` in the `.env` file or in the code to match your local PostgreSQL configuration.
- Run the application:

```sh
go run cmd/app/main.go
```

4. **Swagger Documentation**: After the app is running, you can access the Swagger documentation UI at:

```sh
http://localhost:8080/api/documentation
```

This interface allows you to interact with the API and test the endpoints.
