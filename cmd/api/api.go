package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/YanSystems/compiler/pkg/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	Port string
}

func (s *Server) NewRouter() http.Handler {
	slog.Info("Setting up new router")
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "https://localhost", "http://localhost:3000", "https://localhost:3000", "https://abyan.dev"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	slog.Info("CORS middleware configured")

	// Health check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("Health check endpoint hit")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	slog.Info("Health check route configured")

	// Compiler services
	router.Post("/python", services.HandleExecutePython)
	slog.Info("Compiler service routes configured")

	return router
}

func (s *Server) NewServer() *http.Server {
	slog.Info("Creating new server instance")
	router := s.NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Port),
		Handler: router,
	}
	slog.Info("New server instance created", "port", s.Port)

	return server
}

func (s *Server) Run(server *http.Server) {
	slog.Info(fmt.Sprintf("The server is now live on port %s", s.Port))

	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Server encountered an error", "error", err)
		log.Panic(err)
		return
	}
}
