package filters

import "netEMP/internal/events"

type Filter interface {
	Match(*events.Event) bool
}
