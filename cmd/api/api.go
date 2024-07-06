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
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "https://localhost", "http://localhost:3000", "https://localhost:3000", "https://abyan.dev"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Post("/python", services.HandleExecutePython)

	return router
}

func (s *Server) NewServer() *http.Server {
	router := s.NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Port),
		Handler: router,
	}

	return server
}

func (s *Server) Run(server *http.Server) {
	slog.Info(fmt.Sprintf("The server is now live on port %s", s.Port))

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
		return
	}
}
