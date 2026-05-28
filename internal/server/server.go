package server

import (
	"net/http"
	"netemp/internal/actions"
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
		[]filters.Filter{
			filters.MethodFilter{
				Allowed: "POST",
			},
		},
		[]actions.Action{
			actions.LogAction{},
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

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok\n"))
}

// Processes incoming events, validates request, applies filters,
// and responds with appropriate HTTP status codes
func (s *Server) handleIngest(w http.ResponseWriter, r *http.Request) {
	event, err := events.FromRequest(r)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := s.pipeline.Process(
		r.Context(),
		event,
	); err != nil {

		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) Start(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}
	return server.ListenAndServe()
}
