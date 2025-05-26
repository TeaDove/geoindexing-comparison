package builder_service

import (
	"geoindexing_comparison/pkg/index"
	"geoindexing_comparison/pkg/index/indexes"
	"geoindexing_comparison/pkg/task"
)

type Service struct {
	NameToIndex map[string]index.Index
	Indexes     []index.Index

	NameToTask map[string]task.Task
	Tasks      []task.Task
}

func NewService() *Service {
	r := Service{
		NameToIndex: make(map[string]index.Index),
		Indexes:     make([]index.Index, 0),
		NameToTask:  make(map[string]task.Task),
		Tasks:       make([]task.Task, 0),
	}

	for _, v := range task.AllTasks() {
		r.NameToTask[v.Info.ShortName] = v
		r.Tasks = append(r.Tasks, v)
	}

	for _, v := range indexes.AllIndexes() {
		r.NameToIndex[v.Info.ShortName] = v
		r.Indexes = append(r.Indexes, v)
	}

	return &r
}
