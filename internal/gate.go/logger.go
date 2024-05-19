package gate

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

func LoggerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		span := trace.SpanFromContext(r.Context())
		traceid := ""
		if span != nil {
			traceid = span.SpanContext().TraceID().String()
		}

		next.ServeHTTP(w, r)

		if r.URL.Path != "/ping" {
			logrus.WithFields(logrus.Fields{
				"context":   "logger",
				"component": "logger-mw",
				"duration":  time.Since(start).String(),
				"trace_id":  traceid,
				"method":    r.Method,
			}).Infof("%v", r.URL.Path)
		}
	})
}
