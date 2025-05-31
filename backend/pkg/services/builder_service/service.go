package builder_service

import (
	"geoindexing_comparison/pkg/generator"
	"geoindexing_comparison/pkg/index"
	"geoindexing_comparison/pkg/index/indexes"
	"geoindexing_comparison/pkg/task"
)

type Service struct {
	IndexMap map[string]index.Index
	Indexes  []index.Index

	TaskMap map[string]task.Task
	Tasks   []task.Task

	GeneratorMap map[string]generator.Generator
	Generators   []generator.Generator
}

func NewService() *Service {
	r := Service{
		IndexMap:     make(map[string]index.Index),
		TaskMap:      make(map[string]task.Task),
		GeneratorMap: make(map[string]generator.Generator),
	}

	for _, v := range task.AllTasks() {
		r.TaskMap[v.Info.ShortName] = v
		r.Tasks = append(r.Tasks, v)
	}

	for _, v := range indexes.AllIndexes() {
		r.IndexMap[v.Info.ShortName] = v
		r.Indexes = append(r.Indexes, v)
	}

	for _, v := range generator.AllGenerators() {
		r.GeneratorMap[v.Info.ShortName] = v
		r.Generators = append(r.Generators, v)
	}

	return &r
}
