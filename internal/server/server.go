package server

import (
	"log"
	"net/http"
	"netEMP/internal/events"
	"netEMP/internal/pipeline"
)

type Server struct {
	mux      *http.ServeMux
	pipeline *pipeline.Pipeline
}

func New() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("POST /ingest", s.handleIngest)
}

func (s *Server) handleIngest(
	w http.ResponseWriter,
	r *http.Request,
) {
	event, err := events.FromRequest(r)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	log.Printf(
		"event from=%s method=%s path=%s",
		event.SourceIP,
		event.Method,
		event.Path,
	)

	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) Start(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}

	return server.ListenAndServe()
}
