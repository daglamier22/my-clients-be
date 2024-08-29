package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)
	r.Route("/auth", s.authRoutes)

	return r
}

func (s *Server) authRoutes(r chi.Router) {
	r.Post("/signup", s.SignupHandler)
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	// temp disable DB test and just return a basic response
	// jsonResp, _ := json.Marshal(s.db.Health())
	resp := make(map[string]string)
	resp["serverHealth"] = "up and running"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}
