package generator

import (
	"geoindexing_comparison/geo"

	"github.com/google/uuid"
)

func (r *Generator) GeneratePoint() geo.Point {
	coord := geo.Point{
		Lat: r.randLat(),
		Lon: r.randLon(),
		ID:  uuid.New(),
	}
	return coord
}

func (r *Generator) GeneratePoints(amount int) geo.Points {
	points := make([]geo.Point, amount)
	for i := 0; i < amount; i++ {
		points[i] = r.GeneratePoint()
	}
	return points
}

func (r *Generator) GeneratePointsDefaultAmount() geo.Points {
	return r.GeneratePoints(r.PointsAmount)
}
