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
- Go 1.23 or later

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
├── build/
│   ├── ci/
│   │   └── .keep                # Placeholder for CI-related configurations
│   └── package/
│       └── .keep               # Placeholder for packaging configurations
├── cmd/                        # is the standard location for entry points.
│   └── app/
│       └── main.go             # Entry point for the application
├── configs/
│   └── .env                    # Environment default variables
├── deployments/
│   └── compose.yml             # Docker Compose configuration
├── internal/                   # for private code
│   ├── config/
│   │   └── config.go           # Configuration loader
│   └── logger/
│       ├── composite_logger.go # Composite logger implementation
│       ├── logger.go           # Logger interface and log level definitions
│       ├── sentry_logger.go    # Sentry logger implementation
│       └── standard_logger.go  # Standard console logger implementation
├── Dockerfile                  # Dockerfile for building the application
├── Makefile                    # Makefile for automation tasks
└── README.md                   # Project documentation
```
