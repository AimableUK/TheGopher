# GoAuth

A secure authentication and user management service built with Go and the Gin framework. This project demonstrates best practices for building secure, scalable API services with structured error handling and input validation.

## Features

*   **User Management**: Endpoints for creating and managing user accounts.
*   **Secure Validation**: Built-in request validation using Gin's binding tags.
*   **Structured Architecture**: Organized into models, controllers, and services for better maintainability.
*   **In-Memory Storage**: Implements a thread-safe user store using Go's `sync.RWMutex` for high-concurrency safety.

## Project Structure

```text
GoAuth/
├── models.go       # Data structures and user model definitions
├── main.go         # Entry point and server initialization
└── ...             # (Extend as your project grows)
```

## Getting Started

### Prerequisites

*   [Go 1.20+](https://golang.org/dl/)

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd GoAuth
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

## Development Notes

This project utilizes a **Thread-Safe In-Memory Store**. By using `sync.RWMutex` within the `UserStore` struct, the application ensures that multiple concurrent HTTP requests (from the Gin framework) can safely read and write user data without encountering race conditions or panics.

## Contributing

Contributions are welcome! If you have suggestions for security improvements or feature additions, please feel free to open an issue or submit a pull request.

## License

[MIT](https://choosealicense.com/licenses/mit/)
