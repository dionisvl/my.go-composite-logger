# Go Logger

Composite logger implementation with Sentry integration and structured logging via `log/slog`.

## Architecture

```
.
├── Dockerfile              # Multi-stage build (test + binary)
├── Makefile                # Docker commands
├── cmd/app/main.go         # HTTP server + middleware
├── internal/
│   ├── config/             # .env loading
│   ├── logger/             # slog + Sentry handler
│   └── security/           # HTTP middleware
└── go.mod
```

## Logger

- Structured logging via `log/slog` (Go 1.21+)
- Auto Sentry integration (if `SENTRY_DSN` env set)
- Fanout to multiple handlers via `slog-multi`
- Removed 3 custom files (CompositeLogger, StandardLogger, SentryLogger)

## Middleware

| Middleware | Purpose |
|-----------|---------|
| `RateLimitMiddleware` | Per-IP rate limiting (10 req/min) |
| `MaxBodySizeMiddleware` | Limit request body (64 KB) |
| `LoggingMiddleware` | Structured request logging |

## Docker

Multi-stage Dockerfile:
- Stage 1: Run tests + build binary
- Stage 2: Alpine image (~10 MB)

## Testing

```bash
go test ./... -v
go test ./... -cover
```

Coverage: logger 75%, security 79%

