package indexes

import (
	"geoindexing_comparison/pkg/index"
	"geoindexing_comparison/pkg/index/bruteforce"
	"geoindexing_comparison/pkg/index/geobrtree"
	"geoindexing_comparison/pkg/index/geobtree"
	"geoindexing_comparison/pkg/index/h3btree"
	"geoindexing_comparison/pkg/index/kdtree"
	"geoindexing_comparison/pkg/index/quadtree"
	"geoindexing_comparison/pkg/index/rstartree"
	"geoindexing_comparison/pkg/index/rtree"
	"geoindexing_comparison/pkg/index/rtree_tidwall"
)

func AllIndexes() []index.Index { //nolint: funlen // Combiner
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
			Builder: rtree_tidwall.New,
			Info: index.Info{
				ShortName:   "rtree_tidwal",
				LongName:    "R-дерево Tidwal",
				Description: "Реализация R-tree от разработчика `Tidwal`",
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
			Builder: geobtree.Factory(7),
			Info: index.Info{
				ShortName: "geobtree_7",
				LongName:  "GeoB-дерево 7",
				Description: "Смесь geohash, btree и перебора. Geohash кластеризирует точки, " +
					"которые кладутся в btree и через него производится поиск кластеров. " +
					"Далее в кластера перебором производится поиск указанных точек",
			},
		},
		{
			Builder: geobtree.Factory(6),
			Info: index.Info{
				ShortName:   "geobtree_6",
				LongName:    "GeoB-дерево 6",
				Description: "Аналогично geohash_btree, но с точностью геохеша 6",
			},
		},
		{
			Builder: geobtree.Factory(5),
			Info: index.Info{
				ShortName:   "geobtree_5",
				LongName:    "GeoB-дерево 5",
				Description: "Аналогично geohash_btree, но с точностью геохеша 5",
			},
		},
		{
			Builder: h3btree.Factory(5),
			Info: index.Info{
				ShortName:   "h3btree_5",
				LongName:    "H3B-дерево 5",
				Description: "Смесь h3, btree и перебора",
			},
		},
		{
			Builder: h3btree.Factory(6),
			Info: index.Info{
				ShortName:   "h3btree_6",
				LongName:    "H3B-дерево 6",
				Description: "Смесь h3, btree и перебора",
			},
		},
		{
			Builder: h3btree.Factory(7),
			Info: index.Info{
				ShortName:   "h3btree_7",
				LongName:    "H3B-дерево 7",
				Description: "Смесь h3, btree и перебора",
			},
		},
		{
			Builder: h3btree.Factory(8),
			Info: index.Info{
				ShortName:   "h3btree_8",
				LongName:    "H3B-дерево 8",
				Description: "Смесь h3, btree и перебора",
			},
		},
		{
			Builder: geobrtree.Factory(7),
			Info: index.Info{
				ShortName: "geobrtree",
				LongName:  "GeoBR-дерево",
				Description: "Смесь Geohash, btree и rtree, geohash и btree используются для кластеризации, " +
					"rtree для поиска в кластерах",
			},
		},
	}
}
