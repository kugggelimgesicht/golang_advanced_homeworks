package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		logrus.WithFields(logrus.Fields{
			"method":  r.Method,
			"request": r.URL.Path,
			"status":  wrapper.StatusCode,
			"time":    time.Since(start),
		}).Info("Request served")
	})
}
