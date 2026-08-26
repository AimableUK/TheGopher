# GoTaskit

A simple, efficient Task Management REST API built with Go and the Gin framework. This project serves as a practical implementation of CRUD operations using an in-memory data store.

## Features

*   **RESTful API**: Manage tasks with standard endpoints (GET, POST, PUT, DELETE).
*   **In-Memory Storage**: Uses a thread-safe implementation with `sync.RWMutex` to manage data efficiently without external database dependencies.
*   **Structured Design**: Organized following clean project conventions (Models, Controllers, Data/Service layers).
*   **Concise API**: Clear request/response flow for managing task lifecycle.

## Project Structure

```text
GoTaskit/
├── main.go               # Entry point and server initialization
├── controllers/          # HTTP handlers and request logic
├── models/               # Task data structures
├── data/                 # In-memory service logic
├── router/               # Route definitions
└── docs/                 # API documentation
```

## Getting Started

### Prerequisites

*   [Go 1.20+](https://golang.org/dl/)

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd GoTaskit
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

## Development Approach

This project emphasizes clean backend principles:
*   **Thread Safety**: Uses `sync.RWMutex` to ensure safe concurrent access to the in-memory map.
*   **Separation of Concerns**: Logic is divided between routes, controllers, and services for better code organization.

## API Documentation

For detailed information on request payloads and response formats, please refer to the [api_documentation.md](docs/api_documentation.md) file.

## Contributing

Contributions are welcome! If you have suggestions or improvements, please feel free to open an issue or submit a pull request.

## License

[MIT](https://choosealicense.com/licenses/mit/)
