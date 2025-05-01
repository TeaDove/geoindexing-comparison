package indexes

import (
	"geoindexing_comparison/backend/index"
	"geoindexing_comparison/backend/index/bruteforce"
	"geoindexing_comparison/backend/index/kdtree"
	"geoindexing_comparison/backend/index/quadtree"
	"geoindexing_comparison/backend/index/rstartree"
	"geoindexing_comparison/backend/index/rtree"
)

func AllIndexes() []index.Index {
	return []index.Index{
		{
			Builder: kdtree.New,
			Info: index.Info{
				ShortName: "kdtree",
				LongName:  "K-d-дерево",
				Description: "Структура данных с разбиением пространства " +
					"для упорядочивания точек в k-мерном пространстве",
			},
		},
		{
			Builder: rtree.New,
			Info: index.Info{
				ShortName: "rtree",
				LongName:  "R-дерево",
				Description: "Древовидная структура данных (дерево), предложенная в 1984 году Антонином Гуттманом. " +
					"Она подобна B-дереву, но используется для организации доступа к пространственным данным",
			},
		},
		{
			Builder: rstartree.New,
			Info: index.Info{
				ShortName: "rstartree",
				LongName:  "R*-дерево",
				Description: "Вариант R-деревьев, используемый для индексирования пространственной информации. " +
					"R*-деревья имеют слегка повышенные затраты на создание, чем стандартные R-деревья",
			},
		},
		{
			Builder: quadtree.New,
			Info: index.Info{
				ShortName:   "quadtree",
				LongName:    "Дерево квадрантов",
				Description: "Дерево, в котором у каждого внутреннего узла ровно 4 потомка",
			},
		},
		{
			Builder: bruteforce.New,
			Info: index.Info{
				ShortName: "bruteforce",
				LongName:  "Перебор",
				Description: "Представляет собой обычный динамический массив, " +
					"для которого все операции проводятся простым перебором",
			},
		},
	}
}
