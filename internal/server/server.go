package server

import (
	"log"
	"net/http"
	"netemp/internal/events"
	"netemp/internal/filters"
	"netemp/internal/pipeline"
)

type Server struct {
	mux      *http.ServeMux
	pipeline *pipeline.Pipeline
}

func New() *Server {
	p := pipeline.New(
		filters.MethodFilter{
			Allowed: "POST",
		},
	)

	s := &Server{
		mux:      http.NewServeMux(),
		pipeline: p,
	}

	s.routes()

	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/ingest", s.handleIngest)
	s.mux.HandleFunc("/healthz", s.handleHealthz)
}

// handleHealthz responds with HTTP 200 OK to indicate the server is healthy.
func (s *Server) handleHealthz(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Write([]byte("ok\n"))
}

// handleIngest processes incoming events. It validates the request,
// applies filters, and responds with appropriate HTTP status codes
func (s *Server) handleIngest(
	w http.ResponseWriter,
	r *http.Request,
) {
	event, err := events.FromRequest(r)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if !s.pipeline.Process(event) {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	log.Printf(
		"accepted event from=%s path=%s",
		event.SourceIP,
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
