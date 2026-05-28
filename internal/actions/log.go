package actions

import (
	"context"
	"log"

	"netemp/internal/events"
)

type LogAction struct{}

func (a LogAction) Execute(ctx context.Context, e *events.Event,
) error {
	log.Printf(
		"event accepted ip=%s method=%s path=%s",
		e.SourceIP,
		e.Method,
		e.Path,
	)
	return nil
}
