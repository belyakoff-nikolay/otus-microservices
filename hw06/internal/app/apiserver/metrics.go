package apiserver

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var requestDurationHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "request_duration_histogram_seconds",
		Help:    "Request latency distribution",
		Buckets: prometheus.ExponentialBuckets(0.005, 1.20, 30),
	}, []string{
		"method", "path", "status_code",
	})

func init() {
	prometheus.MustRegister(
		requestDurationHistogram,
	)
}

func (s *APIServer) injectMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		startedAt := time.Now()
		w := &responseWriter{
			ResponseWriter: writer,
			code:           http.StatusOK,
		}

		const sep = "/"
		a := strings.Split(request.URL.Path, sep)
		if len(a) > 2 {
			a[len(a)-1] = ""
		}
		path := strings.Join(a, sep)

		defer func() {
			statusCode := strconv.Itoa(w.StatusCode())
			durationSeconds := time.Since(startedAt).Seconds()
			requestDurationHistogram.WithLabelValues(request.Method, path, statusCode).Observe(durationSeconds)
		}()

		next.ServeHTTP(w, request)
	})
}

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) StatusCode() int {
	return w.code
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
