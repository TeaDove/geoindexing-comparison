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

var All = []Task{
	&KNNQuater{},
	&KNN90{},
	&RadiusSearch{},
	&Insert{},
	&KNN1{},
}

var AllOnePerType = []Task{
	&RadiusSearch{},
	&Insert{},
	&KNN1{},
}

var KnnAndRadiusSearch = []Task{
	&RadiusSearch{},
	&KNN1{},
}

var OnlyRadiusSearch = []Task{
	&RadiusSearch{},
}

var OnlyKNN1 = []Task{
	&KNN1{},
}
