package pipeline

import (
	"netemp/internal/events"
	"netemp/internal/filters"
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
