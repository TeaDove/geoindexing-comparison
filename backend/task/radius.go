package task

import (
	"geoindexing_comparison/backend/geo"
	"time"
)

type BBoxAll struct{}

func (r *BBoxAll) Run(input *Input) time.Duration {
	leftBottom, rightUpper := input.Points.FindCorners()

	_, t := input.Index.BBoxTimed(leftBottom, rightUpper)

	return t
}

type BBox2km struct{}

func (r *BBox2km) Run(input *Input) time.Duration {
	centerLat, centerLon := input.Points.Center()

	_, t := input.Index.BBoxTimed(
		geo.NewPoint(centerLat, centerLon).AddLatitude(-2).AddLongitude(-2),
		geo.NewPoint(centerLat, centerLon).AddLatitude(2).AddLongitude(2),
	)

	return t
}
