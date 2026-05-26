package events

import (
	"net"
	"net/http"
	"time"
)

// Event represents a single event captured from the network;
// every ingress source eventually becomes this.
type Event struct {
	Time      time.Time
	SourceIP  net.IP
	Method    string
	Path      string
	UserAgent string
	Headers   http.Header
	Body      []byte
}
