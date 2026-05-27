package pipeline

import (
	"netEMP/internal/events"
	"netEMP/internal/filters"
)

type Pipeline struct {
	filters []filters.Filter
}

func New(filters ...filters.Filter) *Pipeline {
	return &Pipeline{
		filters: filters,
	}
}

func (p *Pipeline) Process(
	e *events.Event,
) bool {

	for _, f := range p.filters {
		if !f.Match(e) {
			return false
		}
	}

	return true
}
