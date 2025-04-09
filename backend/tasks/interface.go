package tasks

import (
	"geoindexing_comparison/backend/index"
	"time"
)

type TaskImpl interface {
	Name() string
	Filename() string
	Description() string
	Run(index index.IndexImpl, amount uint64) time.Duration
}

type TaskInfo struct {
	ShortName   string `json:"shortName"`
	LongName    string `json:"longName"`
	Description string `json:"description"`
}

type Task struct {
	Info    TaskInfo        `json:"info"`
	Builder func() TaskImpl `json:"-"`
}

func AllTasks() []Task {
	return []Task{
		{
			Info: TaskInfo{
				ShortName:   "knn_quarters",
				LongName:    "КНН 25%",
				Description: "КНН на четверть точек",
			},
			Builder: func() TaskImpl { return &KNNQuater{} },
		},
		{
			Info: TaskInfo{
				ShortName:   "knn_90",
				LongName:    "КНН 90%",
				Description: "КНН на 90% точек из структуры",
			},
			Builder: func() TaskImpl { return &KNN90{} },
		},
		{
			Info: TaskInfo{
				ShortName:   "knn_1",
				LongName:    "КНН",
				Description: "КНН на 1% точек из структуры",
			},
			Builder: func() TaskImpl { return &KNN1{} },
		},
		{
			Info: TaskInfo{
				ShortName:   "radius_search",
				LongName:    "Поиск в радиусе",
				Description: "TDB",
			},
			Builder: func() TaskImpl { return &RadiusSearch{} },
		},
		{
			Info: TaskInfo{
				ShortName:   "insert",
				LongName:    "Вставка",
				Description: "Вставка 10% точек",
			},
			Builder: func() TaskImpl { return &Insert{} },
		},
	}
}
