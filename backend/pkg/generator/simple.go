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
	return RandomPoint(r.rng, input)
}

func (r *SimpleGenerator) Points(input *Input, amount int) geo.Points {
	points := make([]geo.Point, amount)
	for i := range amount {
		points[i] = r.Point(input)
	}

	return points
}

func RandomPoint(rng *rand.Rand, input *Input) geo.Point {
	return geo.NewPoint(
		randFloat(rng, input.LatLowerBound, input.LatUpperBound),
		randFloat(rng, input.LonLowerBound, input.LonUpperBound),
	)
}
