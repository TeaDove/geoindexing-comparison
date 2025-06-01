package indexes

import (
	"geoindexing_comparison/pkg/index"
	"geoindexing_comparison/pkg/index/bruteforce"
	"geoindexing_comparison/pkg/index/geobrtree"
	"geoindexing_comparison/pkg/index/geobtree"
	"geoindexing_comparison/pkg/index/h3btree"
	"geoindexing_comparison/pkg/index/kdtree"
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
			Builder: rtree_tidwall.New,
			Info: index.Info{
				ShortName:   "rtree_tidwal",
				LongName:    "R-дерево Tidwal",
				Description: "Реализация R-tree от разработчика `Tidwal`",
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
			Builder: geobtree.Factory(6),
			Info: index.Info{
				ShortName:   "geobtree_6",
				LongName:    "GeoB-дерево 6",
				Description: "Аналогично geohash_btree, но с точностью геохеша 6",
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
				ShortName: "geobrtree_7",
				LongName:  "GeoBR-дерево 7",
				Description: "Смесь Geohash, btree и rtree, geohash и btree используются для кластеризации, " +
					"rtree для поиска в кластерах",
			},
		},
	}
}
