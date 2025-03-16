package cases

import (
	"context"
	"geoindexing_comparison/service/cases/tasks"
	"geoindexing_comparison/service/index"
	"geoindexing_comparison/service/index/indexes"
	"github.com/rs/zerolog"
	"maps"
	"slices"
)

type Runner struct {
	NameToIndex map[string]index.Index
	Indexes     []index.Index

	NameToTask map[string]tasks.Task
	Tasks      []tasks.Task

	stop context.CancelFunc
}

func NewRunner(ctx context.Context) *Runner {
	r := Runner{
		NameToIndex: make(map[string]index.Index),
		Indexes:     make([]index.Index, 0),
		NameToTask:  make(map[string]tasks.Task),
		Tasks:       make([]tasks.Task, 0),
	}

	for _, v := range tasks.AllTasks() {
		r.NameToTask[v.Info.ShortName] = v
		r.Tasks = append(r.Tasks, v)
	}

	for _, v := range indexes.AllIndexes() {
		r.NameToIndex[v.Info.ShortName] = v
		r.Indexes = append(r.Indexes, v)
	}

	zerolog.Ctx(ctx).
		Info().
		Strs("indexes", slices.Collect(maps.Keys(r.NameToIndex))).
		Strs("tasks", slices.Collect(maps.Keys(r.NameToTask))).
		Msg("runner.initialized")
	return &r
}
