package security

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/httprate"
)

const (
	MaxBodySize = 64 * 1024 // 64 KB
)

// MaxBodySizeMiddleware enforces a maximum request body size.
// Returns 413 Payload Too Large if exceeded.
func MaxBodySizeMiddleware(maxBytes int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
			next.ServeHTTP(w, r)
		})
	}
}

// RateLimitMiddleware applies per-IP rate limiting.
// Default: 10 requests per minute.
func RateLimitMiddleware(requestsPerMin int) func(http.Handler) http.Handler {
	return httprate.LimitByIP(requestsPerMin, time.Minute)
}

// LoggingMiddleware logs request metadata (method, path, IP, status, duration).
// Does not log request/response bodies.
func LoggingMiddleware(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap response writer to capture status code
			wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(wrapped, r)

			duration := time.Since(start)
			log.Info("http_request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("ip", r.RemoteAddr),
				slog.Int("status", wrapped.statusCode),
				slog.Duration("duration_ms", duration),
			)
		})
	}
}

// responseWriter wraps http.ResponseWriter to capture status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.statusCode = code
		rw.written = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		rw.written = true
	}
	return rw.ResponseWriter.Write(b)
}
