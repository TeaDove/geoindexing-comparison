package generator

import (
	"geoindexing_comparison/service/geo"
)

type SimpleGenerator struct{}

var DefaultGenerator = SimpleGenerator{}

func (r *SimpleGenerator) Point(input *Input) geo.Point {
	return geo.NewPoint(
		randFloat(input.LatLowerBound, input.LatUpperBound),
		randFloat(input.LonLowerBound, input.LonUpperBound),
	)
}

func (r *SimpleGenerator) Points(input *Input, amount int) geo.Points {
	points := make([]geo.Point, amount)
	for i := 0; i < amount; i++ {
		points[i] = r.Point(input)
	}

	return points
}
