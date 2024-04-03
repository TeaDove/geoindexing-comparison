package quadtree

import (
	"geoindexing_comparison/core/addapter"
	"geoindexing_comparison/core/geo"
	qtree "github.com/TeaDove/go-quad-tree"
	"github.com/google/uuid"
	"time"
)

type CollectionQuadTree struct {
	impl qtree.Qtree[uuid.UUID]
}

func New() addapter.Collection {
	r := CollectionQuadTree{}

	r.impl = *qtree.NewQtree[uuid.UUID](0, 0, 180, 180, 10)

	return &r
}

func (r *CollectionQuadTree) Name() string {
	return "QuadTree"
}

func (r *CollectionQuadTree) FromArray(points geo.Points) {
	for _, point_ := range points {
		r.impl.Insert(qtree.NewPoint(point_.Lat, point_.Lon, point_.ID))
	}
}

func (r *CollectionQuadTree) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()
	r.impl.Insert(qtree.NewPoint(point.Lat, point.Lon, point.ID))

	return time.Now().Sub(t0)
}

func toConcrete(qtreePoints []qtree.Point[uuid.UUID]) geo.Points {
	geoPoints := make(geo.Points, 0, len(qtreePoints))
	for _, point := range qtreePoints {
		geoPoints = append(geoPoints, geo.Point{
			ID:  point.Val,
			Lat: point.X,
			Lon: point.Y,
		})
	}

	return geoPoints
}

func (r *CollectionQuadTree) RangeSearchTimed(
	point geo.Point,
	radius float64,
) (geo.Points, time.Duration) {
	t0 := time.Now()

	res := r.impl.QueryRange(
		qtree.NewBounds[uuid.UUID](
			point.Lat-radius/2,
			point.Lon-radius/2,
			point.Lat+radius/2,
			point.Lon+radius/2,
		),
	)
	dur := time.Now().Sub(t0)

	return toConcrete(res), dur
}

func (r *CollectionQuadTree) KNNTimed(point geo.Point, n int) (geo.Points, time.Duration) {
	t0 := time.Now()
	res := r.impl.KNN(qtree.NewPoint[uuid.UUID](point.Lat, point.Lon, uuid.Nil), n)
	dur := time.Now().Sub(t0)

	return toConcrete(res), dur
}

func (r *CollectionQuadTree) String() string {
	return r.impl.String()
}
