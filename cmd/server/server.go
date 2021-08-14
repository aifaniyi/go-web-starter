package main

import (
	"encoding/json"
	"net/http"

	"github.com/aifaniyi/sample/pkg/exception"
	"github.com/aifaniyi/sample/pkg/repository"
	"github.com/go-chi/chi"

	"github.com/aifaniyi/sample/pkg/logger"
)

type server struct {
	repo   repository.Service
	router *chi.Mux
}

func NewServer(repo repository.Service) *server {
	s := &server{
		repo:   repo,
		router: chi.NewRouter(),
	}

	s.routes()
	return s
}

func (s *server) routes() {
	s.router.Use(s.loggingMiddleWare)
	s.router.Get("/", s.handleIndex())

	s.router.Post("/api/v1/signup", s.handleSignup())
}

func (s *server) start(port string) {
	server := &http.Server{
		Addr:    port,
		Handler: s.router,
	}

	logger.Info.Printf("starting server on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		logger.Info.Printf("http server error: %s", err.Error())
	}
}

type response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func httpResponse(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response{Data: data})
}

func httpError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	switch err.(type) {
	case *exception.ClientError:
		w.WriteHeader(http.StatusBadRequest)

	default:
		w.WriteHeader(code)
	}
	json.NewEncoder(w).Encode(response{Error: err.Error()})
}
