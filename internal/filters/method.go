package filters

import "netemp/internal/events"

type MethodFilter struct {
	Allowed string
}

func (f MethodFilter) Match(e *events.Event) bool {
	return e.Method == f.Allowed
}
