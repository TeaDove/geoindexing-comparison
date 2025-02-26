package tasks

import (
	"geoindexing_comparison/service/index"
	"time"
)

type Task interface {
	Name() string
	Filename() string
	Description() string
	Run(index index.Index, amount int) time.Duration
}

var Tasks = []Task{
	&KNNQuater{},
	&KNN90{},
	&KNN1{},
	&RadiusSearch{},
	&Insert{},
}

var NameToTask = func() map[string]Task {
	mapping := make(map[string]Task)
	for _, task := range Tasks {
		mapping[task.Name()] = task
	}
	return mapping
}()
