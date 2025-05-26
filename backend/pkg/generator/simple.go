package generator

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/helpers"
	"math/rand/v2"
)

type SimpleGenerator struct{ rng *rand.Rand }

func NewSimplerGenerator() Impl {
	return &SimpleGenerator{rng: helpers.RNG()}
}

func (r *SimpleGenerator) Point(input *Input) geo.Point {
	return geo.NewPoint(
		randFloat(r.rng, input.LatLowerBound, input.LatUpperBound),
		randFloat(r.rng, input.LonLowerBound, input.LonUpperBound),
	)
}

func (r *SimpleGenerator) Points(input *Input, amount int) geo.Points {
	points := make([]geo.Point, amount)
	for i := range amount {
		points[i] = r.Point(input)
	}

	return points
}
