package generator

import (
	"geoindexing_comparison/backend/geo"
)

type SimpleGenerator struct{}

var DefaultGenerator = SimpleGenerator{}

func (r *SimpleGenerator) Point(input *Input) geo.Point {
	return geo.NewPoint(
		randFloat(input.LatLowerBound, input.LatUpperBound),
		randFloat(input.LonLowerBound, input.LonUpperBound),
	)
}

func (r *SimpleGenerator) Points(input *Input, amount uint64) geo.Points {
	points := make([]geo.Point, amount)
	for i := range amount {
		points[i] = r.Point(input)
	}

	return points
}
