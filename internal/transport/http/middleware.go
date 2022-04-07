package http

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)


func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		next.ServeHTTP(w,r)
	})
}



func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"method": r.Method,
				"path": r.URL.Path,
			}).Info("handled request")

		next.ServeHTTP(w,r)
	})
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	// Create new context that calls cancel after 15 seconds have passed
	// every request coming through the service only has 15 seconds to return a response 
	// or else an error will be returned.

	// flow overview
	// context gets passed into the
	// transport layer function then
	// through to the service layers, which then goes into the storage layer
	// where the storage layer uses context to make the db query.
	// now if the the cancel func is called then original query will be cancled, any processing, 
	// and propigated back up to the handler function

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

