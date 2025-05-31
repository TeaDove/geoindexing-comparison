package generator

import (
	"geoindexing_comparison/pkg/geo"
	"math/rand/v2"
)

type Impl interface {
	Point(input *Input) geo.Point
	Points(input *Input, amount int) geo.Points
}

type Info struct {
	ShortName   string `json:"shortName"`
	LongName    string `json:"longName"`
	Description string `json:"description"`
}

type Generator struct {
	Info    Info                      `json:"info"`
	Builder func(rng *rand.Rand) Impl `json:"-"`
}

func AllGenerators() []Generator {
	return []Generator{
		{
			Info: Info{
				ShortName:   "simple_generator",
				LongName:    "Простой генератор",
				Description: "Генерирует точки используя простые случайные точки",
			},
			Builder: func(rng *rand.Rand) Impl { return &SimpleGenerator{rng: rng} },
		},
		{
			Info: Info{
				ShortName:   "cluster_generator",
				LongName:    "Кластеный генератор",
				Description: "Генерирует точки в 6 кластерах",
			},
			Builder: func(rng *rand.Rand) Impl { return &ClusterGenerator{ClusterN: 6, rng: rng} },
		},
	}
}
