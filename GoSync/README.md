# GoSync

A user synchronization and management service built with Go and the Gin framework. This project demonstrates a production-ready approach to building REST APIs by implementing a **Layered Architecture** (Controller -> Service -> Data Layer).

## Features

*   **RESTful API**: Comprehensive endpoints to Create, Get, Update, and Delete user records.
*   **Layered Architecture**: Decouples HTTP request handling from business logic and database operations for improved testability and maintainability.
*   **MongoDB Integration**: Uses the official MongoDB Go driver to manage persistent user data.
*   **Dependency Injection**: Components are injected into one another, making the codebase easier to unit test and scale.

## Project Structure

```text
GoSync/
├── controllers/    # HTTP handlers and request validation
├── models/         # Data structures and user schemas
├── services/       # Business logic and database operations
└── main.go         # Entry point, dependency wiring, and server initialization
```

## Getting Started

### Prerequisites

*   [Go 1.20+](https://golang.org/dl/)
*   [MongoDB](https://www.mongodb.com/)

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd GoSync
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure your MongoDB connection string in `main.go`.

4. Run the application:
   ```bash
   go run main.go
   ```

## Development Approach

This project focuses on **Clean Architecture**. By separating the Controller from the Service layer:
*   Controllers remain thin, focusing solely on HTTP request/response.
*   Business logic is contained within the Service layer.
*   Dependency injection allows for easier integration of mock databases during testing.

## Contributing

Contributions are welcome! If you have suggestions for improvements, bug fixes, or feature additions, please feel free to open an issue or submit a pull request.

## License

[MIT](https://choosealicense.com/licenses/mit/)
