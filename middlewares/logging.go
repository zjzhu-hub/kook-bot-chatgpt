package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Pass request down the chain
		next.ServeHTTP(w, r)

		// Calculate request processing time
		elapsed := time.Since(startTime)

		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, elapsed)
	})
}