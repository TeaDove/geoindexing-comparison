package task

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/index/rtree"
	"time"
)

type BBox100 struct{}

func (r *BBox100) Run(input *Input) time.Duration {
	idx := rtree.New()
	idx.FromArray(input.Points)
	points, _ := idx.KNNTimed(input.RandomPoint, 100)

	leftBottom, rightUpper := points.FindCorners()

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
