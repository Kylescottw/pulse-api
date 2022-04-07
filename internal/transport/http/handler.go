package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

 


type Handler struct {
	Router *mux.Router
	Service CommentService
	Server *http.Server
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service, 
	}

	h.Router = mux.NewRouter()

	h.mapRoutes()
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	h.Router.HandleFunc("/api/v1/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.DeleteComment).Methods("DELETE")
}

func (h *Handler) Serve() error {
	go func() {
		// make listen and serve a non-blocking operation
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1) // create a blocking channel
	signal.Notify(c, os.Interrupt) // wait for os.Interrupt signal to unblock the channel
	<-c // code will resume here once channel is unblocked

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // defer calling calcel function for 15 seconds
	h.Server.Shutdown(ctx) // shutdown gracefully, first close all open listeners, then all idle connections, waiting indefinelty for connections to return to idle, then shuts down the server 

	log.Println("shut down gracefully")

	return nil
}