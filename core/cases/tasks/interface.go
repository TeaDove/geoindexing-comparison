package tasks

import (
	"geoindexing_comparison/core/addapter/addapter_all"
	"time"
)

type Task interface {
	Name() string
	Description() string
	Run(col addapter_all.CollectionInit, amount int) time.Duration
}

func All() []Task {
	return []Task{
		&KNN{},
		&RadiusSearch{},
		&Insert{},
	}
}
