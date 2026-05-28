package filters

import "netemp/internal/events"

type Filter interface {
	Match(*events.Event) bool
}
