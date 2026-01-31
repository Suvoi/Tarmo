package middleware

import (
	"net/http"
	"time"

	"tarmo/internal/lib/logger"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		msg := "%s %s %d %s"
		args := []any{
			r.Method,
			r.URL.Path,
			rw.status,
			duration,
		}

		switch {
		case rw.status >= 500:
			logger.Error(msg, args...)
		case rw.status >= 400:
			logger.Warn(msg, args...)
		default:
			logger.Info(msg, args...)
		}
	})
}
