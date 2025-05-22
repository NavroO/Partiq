package shared

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
	body   bytes.Buffer
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	if LogPayloads {
		lrw.body.Write(b)
	}
	return lrw.ResponseWriter.Write(b)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqBody string
		if LogPayloads && r.Body != nil && r.Method != http.MethodGet {
			buf, _ := io.ReadAll(r.Body)
			reqBody = string(buf)
			r.Body = io.NopCloser(bytes.NewBuffer(buf))
		}

		lrw := &loggingResponseWriter{ResponseWriter: w, status: 200}
		start := time.Now()
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)

		event := log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", lrw.status).
			Dur("duration", duration)

		if LogPayloads {
			event = event.
				Str("request_payload", reqBody).
				Str("response_payload", lrw.body.String())
		}

		event.Msg("📥 HTTP request")
	})
}
