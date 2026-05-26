package server

import "net/http"

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /healthz", s.handleHealth)
}

func (s *Server) handleHealth(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func (s *Server) Start(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}

	return server.ListenAndServe()
}
