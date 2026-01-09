# Project Go API Tasks

This project is a simple Task Management API built with Go, utilizing the Gin framework for routing and SQLite for data persistence.

## Technologies Used

- **Go**: The core programming language (v1.25.3).
- **Gin**: A Web Framework written in Go (Gin-Gonic).
- **SQLite**: A C-language library that implements a small, fast, self-contained, high-reliability, full-featured, SQL database engine. (Driver: `modernc.org/sqlite`).

## Project Structure

```text
.
├── db.go           # Database initialization and table creation
├── go.mod          # Go module definitions and dependencies
├── go.sum          # Checksums for dependencies
├── main.go         # Entry point of the application
├── routes.go       # Route definitions
├── services.go     # Business logic and handler functions
├── tasks.db        # SQLite database file (not versioned/ignored)
└── README.md       # Project documentation
```

## API Endpoints

The API provides the following endpoints for managing tasks:

| Method | Endpoint     | Description          |
| :----- | :----------- | :------------------- |
| GET    | `/`          | Health check / Hello |
| GET    | `/tasks`     | Get all tasks        |
| POST   | `/tasks`     | Create a new task    |
| GET    | `/tasks/:id` | Get a task by ID     |
| PUT    | `/tasks/:id` | Update a task by ID  |
| DELETE | `/tasks/:id` | Delete a task by ID  |

## How to Run

To run this project locally, ensure you have Go installed, then execute the following command in the root directory:

```bash
go run .
```

The server will start on port `3000`. You can access it at `http://localhost:3000`.
