package task

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"time"
)

type Input struct {
	Index       index.Impl
	Amount      uint64
	Points      geo.Points
	RandomPoint geo.Point
}

type Impl interface {
	Run(input *Input) time.Duration
}

type Info struct {
	ShortName   string `json:"shortName"`
	LongName    string `json:"longName"`
	Description string `json:"description"`
}

type Task struct {
	Info    Info        `json:"info"`
	Builder func() Impl `json:"-"`
}

func AllTasks() []Task {
	return []Task{
		{
			Info: Info{
				ShortName:   "knn_25_p",
				LongName:    "КНН 25%",
				Description: "КНН на 25% точек",
			},
			Builder: func() Impl { return &KNN25P{} },
		},
		{
			Info: Info{
				ShortName:   "knn_90_p",
				LongName:    "КНН 90%",
				Description: "КНН на 90% точек из структуры",
			},
			Builder: func() Impl { return &KNN90P{} },
		},
		{
			Info: Info{
				ShortName:   "knn_1_p",
				LongName:    "КНН 1%",
				Description: "КНН на 1% точек из структуры",
			},
			Builder: func() Impl { return &KNN1P{} },
		},
		{
			Info: Info{
				ShortName:   "knn_10",
				LongName:    "КНН 10",
				Description: "КНН на 10 точек из структуры",
			},
			Builder: func() Impl { return &KNN10{} },
		},
		{
			Info: Info{
				ShortName:   "knn_100",
				LongName:    "КНН 100",
				Description: "КНН на 100 точек из структуры",
			},
			Builder: func() Impl { return &KNN100{} },
		},
		{
			Info: Info{
				ShortName:   "bbox_all",
				LongName:    "BBox all",
				Description: "Поиск всех точек",
			},
			Builder: func() Impl { return &BBoxAll{} },
		},
		{
			Info: Info{
				ShortName:   "bbox_2km",
				LongName:    "BBox 2km",
				Description: "Поиск точек в прямоугольнике 2 на 2 км.",
			},
			Builder: func() Impl { return &BBox2km{} },
		},
		//{
		//	Info: TaskInfo{
		//		ShortName:   "insert",
		//		LongName:    "Вставка",
		//		Description: "Вставка 10% точек",
		//	},
		//	Builder: func() TaskImpl { return &Insert{} },
		// },
	}
}
