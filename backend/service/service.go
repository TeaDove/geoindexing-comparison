package service

import (
	"context"
	"geoindexing_comparison/backend/index"
	"geoindexing_comparison/backend/index/indexes"
	"geoindexing_comparison/backend/repository"
	"geoindexing_comparison/backend/task"
	"github.com/teadove/teasutils/utils/must_utils"
	"maps"
	"slices"

	"github.com/rs/zerolog"
)

type Service struct {
	NameToIndex map[string]index.Index
	Indexes     []index.Index

	NameToTask map[string]task.Task
	Tasks      []task.Task

	repository *repository.Repository
	Visualizer Visualizer
}

func NewRunner(ctx context.Context, repository *repository.Repository) *Service {
	r := Service{
		NameToIndex: make(map[string]index.Index),
		Indexes:     make([]index.Index, 0),
		NameToTask:  make(map[string]task.Task),
		Tasks:       make([]task.Task, 0),
		repository:  repository,
	}

	for _, v := range task.AllTasks() {
		r.NameToTask[v.Info.ShortName] = v
		r.Tasks = append(r.Tasks, v)
	}

	for _, v := range indexes.AllIndexes() {
		r.NameToIndex[v.Info.ShortName] = v
		r.Indexes = append(r.Indexes, v)
	}
	must_utils.Must(r.SetVisualizer(ctx, &NewVisualizerInput{Amount: 10, Index: r.Indexes[0].Info.ShortName}))

	go r.initRunner()

	zerolog.Ctx(ctx).
		Info().
		Strs("indexes", slices.Collect(maps.Keys(r.NameToIndex))).
		Strs("tasks", slices.Collect(maps.Keys(r.NameToTask))).
		Msg("runner.initialized")

	return &r
}
