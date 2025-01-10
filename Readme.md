# Go-composite-logger (with Sentry as example)

## Description
This project demonstrates a composite logger implementation in Go. The logger supports standard console logging and integrates with Sentry for error tracking.

## Features
- Composite logger (console + Sentry)
- Configurable log levels
- Easy to extend to another logger
- Automatic HTTP transaction tracing with Sentry

## Setup

### Prerequisites
- Docker and Docker Compose
- Go 1.24 or later

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/dionisvl/my.go-composite-logger.git
   cd go-composite-logger
   ```

2. Create an `.env.override` file and fill variables:
   ```
   USE_SENTRY=true
   SENTRY_DSN=https://your_sentry_dsn@sentry.io/123456
   ```

3. Build and run the application using Docker Compose:
   ```bash
   make up
   ```

4. Access the application at [http://localhost:8080/hello](http://localhost:8080/hello).

## Usage
- To stop the application:
  ```bash
  make down
  ```

## Project Structure
```
project-root/
├── logger/
│   ├── logger.go              # Interface for logger
│   ├── standard_logger.go     # Standard logger implementation
│   ├── sentry_logger.go       # Sentry logger implementation
│   ├── composite_logger.go    # Composite logger implementation
├── config/
│   ├── config.go              # Configuration loader
├── main.go                    # Main application entry point
├── go.mod                     # Go modules definition
├── Makefile                   # Makefile for automation
├── Dockerfile                 # Dockerfile for building the app
├── docker-compose.yml         # Docker Compose configuration
├── .env.override              # Environment variables for overriding
└── README.md                  # Project documentation
```