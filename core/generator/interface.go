package generator

import "geoindexing_comparison/core/geo"

type Generator interface {
	Point(input *Input) geo.Point
	Points(input *Input, amount int) geo.Points
}

func All() []Generator {
	return []Generator{
		&NormalGenerator{},
		&SimpleGenerator{},
	}
}
