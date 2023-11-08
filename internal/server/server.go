package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/offluck/ilove2rest/internal/repository"
	"go.uber.org/zap"
)

type Server struct {
	http.Server
	DBClient repository.Client
	logger   *zap.Logger
}

func New(port uint16, dbClient repository.Client, logger *zap.Logger) *Server {
	s := new(Server)
	s.Addr = fmt.Sprintf(":%d", port)
	s.Handler = s.setUpRouter()
	s.DBClient = dbClient
	s.logger = logger
	return s
}

func (s *Server) setUpRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", s.healthHandler)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v0", func(r chi.Router) {
			r.Route("/user", func(r chi.Router) {
				r.Get("/", s.getUsersHandler)
				r.Post("/", s.postUserHandler)
				r.Get("/{username}", s.getUserHandler)
				r.Put("/{username}", s.putUserHandler)
				r.Delete("/{username}", s.deleteUserHandler)
			})
		})
	})

	return r
}

func (s *Server) Start() error {
	return s.ListenAndServe()
}
