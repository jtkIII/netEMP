package pipeline

import (
	"context"
	"netemp/internal/actions"
	"netemp/internal/events"
	"netemp/internal/filters"
)

type Pipeline struct {
	filters []filters.Filter
	actions []actions.Action
}

func New(filters []filters.Filter, actions []actions.Action) *Pipeline {
	return &Pipeline{
		filters: filters,
		actions: actions,
	}
}

func (p *Pipeline) Process(ctx context.Context, e *events.Event) error {

	for _, f := range p.filters {
		if !f.Match(e) {
			return nil
		}
	}

	for _, a := range p.actions {
		if err := a.Execute(ctx, e); err != nil {
			return err
		}
	}

	return nil
}
