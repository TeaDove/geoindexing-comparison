package generator

import "geoindexing_comparison/backend/geo"

type Generator interface {
	Point(input *Input) geo.Point
	Points(input *Input, amount uint64) geo.Points
}

func All() []Generator {
	return []Generator{
		&NormalGenerator{},
		&SimpleGenerator{},
	}
}
