package indexes

import (
	"geoindexing_comparison/backend/index"
	"geoindexing_comparison/backend/index/bruteforce"
	"geoindexing_comparison/backend/index/geohash_btree"
	"geoindexing_comparison/backend/index/h3_btree"
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
		{
			Builder: geohash_btree.Factory(7),
			Info: index.Info{
				ShortName: "geohash_btree_7",
				LongName:  "Геохэш + Б-дерево 7",
				Description: "Смесь geohash, btree и перебора. Geohash кластеризирует точки, " +
					"которые кладутся в btree и через него производится поиск кластеров. " +
					"Далее в кластера перебором производится поиск указанных точек",
			},
		},
		{
			Builder: geohash_btree.Factory(6),
			Info: index.Info{
				ShortName:   "geohash_btree_6",
				LongName:    "Геохэш + Б-дерево 6",
				Description: "Аналогично geohash_btree, но с точностью геохеша 6",
			},
		},
		{
			Builder: geohash_btree.Factory(5),
			Info: index.Info{
				ShortName:   "geohash_btree_5",
				LongName:    "Геохэш + Б-дерево 5",
				Description: "Аналогично geohash_btree, но с точностью геохеша 5",
			},
		},
		{
			Builder: h3_btree.Factory(5),
			Info: index.Info{
				ShortName:   "h3_btree_5",
				LongName:    "H3 + Б-дерево 5",
				Description: "Смесь h3, btree и перебора",
			},
		},
		{
			Builder: h3_btree.Factory(6),
			Info: index.Info{
				ShortName:   "h3_btree_6",
				LongName:    "H3 + Б-дерево 6",
				Description: "Смесь h3, btree и перебора",
			},
		},
		{
			Builder: h3_btree.Factory(7),
			Info: index.Info{
				ShortName:   "h3_btree_7",
				LongName:    "H3 + Б-дерево 7",
				Description: "Смесь h3, btree и перебора",
			},
		},
		{
			Builder: h3_btree.Factory(8),
			Info: index.Info{
				ShortName:   "h3_btree_8",
				LongName:    "H3 + Б-дерево 8",
				Description: "Смесь h3, btree и перебора",
			},
		},
	}
}
