package cases

import (
	"context"
	"geoindexing_comparison/service/cases/tasks"
	"geoindexing_comparison/service/index"
	"github.com/rs/zerolog"
	"maps"
	"slices"
)

type Runner struct {
	NameToIndex map[string]index.NewIndex
	NameToTask  map[string]tasks.Task
}

func NewRunner(ctx context.Context) *Runner {
	r := Runner{}

	zerolog.Ctx(ctx).
		Info().
		Strs("indexes", slices.Collect(maps.Keys(r.NameToIndex))).
		Strs("tasks", slices.Collect(maps.Keys(r.NameToTask))).
		Msg("runner.initialized")
	return &r
}
