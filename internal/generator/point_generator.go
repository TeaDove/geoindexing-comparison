package generator

import "geoindexing_comparison/geo"

func (r *Generator) GeneratePoint() geo.Point {
	coord := geo.Point{
		Lat: r.randLat(),
		Lon: r.randLon(),
	}
	return coord
}

func (r *Generator) GeneratePoints(amount int) []geo.Point {
	var points = make([]geo.Point, amount)
	for i := 0; i < amount; i++ {
		points[i] = r.GeneratePoint()
	}
	return points
}

func (r *Generator) GeneratePointsDefaultAmount() []geo.Point {
	return r.GeneratePoints(r.PointsAmount)
}
