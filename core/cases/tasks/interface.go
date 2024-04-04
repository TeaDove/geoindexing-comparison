package tasks

import (
	"geoindexing_comparison/core/addapter"
	"time"
)

type Task interface {
	Name() string
	Filename() string
	Description() string
	Run(col addapter.Collection, amount int) time.Duration
}

func All() []Task {
	return []Task{
		&KNNQuater{},
		&KNN90{},
		&RadiusSearch{},
		&Insert{},
		&KNN1{},
	}
}
