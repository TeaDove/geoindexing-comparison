package indexes

import (
	"geoindexing_comparison/service/index"
	"geoindexing_comparison/service/index/bruteforce"
	"geoindexing_comparison/service/index/kdtree"
	"geoindexing_comparison/service/index/quadtree"
	"geoindexing_comparison/service/index/rstartree"
	"geoindexing_comparison/service/index/rtree"
)

var Indexes = []index.NewIndex{
	kdtree.New,
	rtree.New,
	rstartree.New,
	quadtree.New,
	bruteforce.New,
}

func AllIndexes() []index.Index {
	return []index.Index{
		{
			Builder: kdtree.New,
			Info: index.IndexInfo{
				ShortName:   "kdtree",
				LongName:    "K-d-дерево",
				Description: "Структура данных с разбиением пространства для упорядочивания точек в k-мерном пространстве",
			},
		},
		{
			Builder: rtree.New,
			Info: index.IndexInfo{
				ShortName:   "rtree",
				LongName:    "R-дерево",
				Description: "Древовидная структура данных (дерево), предложенная в 1984 году Антонином Гуттманом. Она подобна B-дереву, но используется для организации доступа к пространственным данным",
			},
		},
		{
			Builder: rstartree.New,
			Info: index.IndexInfo{
				ShortName:   "rstartree",
				LongName:    "R*-дерево",
				Description: "Вариант R-деревьев, используемый для индексирования пространственной информации. R*-деревья имеют слегка повышенные затраты на создание, чем стандартные R-деревья",
			},
		},
		{
			Builder: quadtree.New,
			Info: index.IndexInfo{
				ShortName:   "quadtree",
				LongName:    "Дерево квадрантов",
				Description: "Дерево, в котором у каждого внутреннего узла ровно 4 потомка",
			},
		},
		{
			Builder: bruteforce.New,
			Info: index.IndexInfo{
				ShortName:   "bruteforce",
				LongName:    "Перебор",
				Description: "Представляет собой обычный динамический массив, для которого все операции проводятся простым перебором",
			},
		},
	}
}
