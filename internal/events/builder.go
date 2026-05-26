package events

import (
	"io"
	"net"
	"net/http"
	"time"
)

func FromRequest(r *http.Request) (*Event, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	return &Event{
		Time:      time.Now().UTC(),
		SourceIP:  net.ParseIP(host),
		Method:    r.Method,
		Path:      r.URL.Path,
		UserAgent: r.UserAgent(),
		Headers:   r.Header.Clone(),
		Body:      body,
	}, nil
}
