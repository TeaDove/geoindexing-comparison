package generator

import (
	"geoindexing_comparison/backend/geo"
)

type Impl interface {
	Point(input *Input) geo.Point
	Points(input *Input, amount uint64) geo.Points
}

type Info struct {
	ShortName   string `json:"shortName"`
	LongName    string `json:"longName"`
	Description string `json:"description"`
}

type Generator struct {
	Info    Info        `json:"info"`
	Builder func() Impl `json:"-"`
}

func AllGenerators() []Generator {
	return []Generator{
		{
			Info: Info{
				ShortName:   "simple_generator",
				LongName:    "Простой генератор",
				Description: "Генерирует точки используя простые случайные точки",
			},
			Builder: func() Impl { return &SimpleGenerator{} },
		},
		//{
		//	Info: Info{
		//		ShortName:   "normal_generator",
		//		LongName:    "Нормальный генератор",
		//		Description: "Генерирует точки нормально",
		//	},
		//	Builder: func() Impl { return &NormalGenerator{ClusterN: 6} },
		//},
	}
}
