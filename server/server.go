package server

import "net/http"

// New http handler
func New(s *http.ServeMux) *Handler {
	h := Handler{s}
	h.registerRoutes()

	return &h
}

// Register routes for all http endpoints
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/", h.HelloWorld)
}

func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world !!"))
}
