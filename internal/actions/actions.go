package actions

import (
	"context"

	"netemp/internal/events"
)

type Action interface {
	Execute(context.Context, *events.Event) error
}
