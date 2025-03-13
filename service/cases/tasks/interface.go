package tasks

import (
	"geoindexing_comparison/service/index"
	"maps"
	"slices"
	"time"
)

type TaskImpl interface {
	Name() string
	Filename() string
	Description() string
	Run(index index.Index, amount uint64) time.Duration
}

type TaskInfo struct {
	ShortName   string `json:"shortName"`
	LongName    string `json:"longName"`
	Description string `json:"description"`
}

type Task struct {
	TaskInfo
	Builder func() TaskImpl
}

func AllTasks() []Task {
	return []Task{
		{
			TaskInfo: TaskInfo{
				ShortName:   "knn_quarters",
				LongName:    "КНН",
				Description: "КНН на четверть точек",
			},
			Builder: func() TaskImpl { return &KNNQuater{} },
		},
		{
			TaskInfo: TaskInfo{ShortName: "knn_90",
				LongName:    "КНН",
				Description: "КНН на 90% точек из структуры",
			},
			Builder: func() TaskImpl { return &KNN90{} },
		},
		{
			TaskInfo: TaskInfo{ShortName: "knn_1",
				LongName:    "КНН",
				Description: "КНН на 1% точек из структуры",
			},
			Builder: func() TaskImpl { return &KNN1{} },
		},
		{
			TaskInfo: TaskInfo{ShortName: "radius_search",
				LongName:    "Поиск в радиусе",
				Description: "TDB",
			},
			Builder: func() TaskImpl { return &RadiusSearch{} },
		},
		{
			TaskInfo: TaskInfo{ShortName: "insert",
				LongName:    "Вставка",
				Description: "Вставка 10% точек",
			},
			Builder: func() TaskImpl { return &Insert{} },
		},
	}
}

var Tasks = []TaskImpl{
	&KNNQuater{},
	&KNN90{},
	&KNN1{},
	&RadiusSearch{},
	&Insert{},
}

var NameToTask = func() map[string]TaskImpl {
	mapping := make(map[string]TaskImpl)
	for _, task := range Tasks {
		mapping[task.Name()] = task
	}
	return mapping
}()

var TaskNames = slices.Sorted(maps.Keys(NameToTask))
