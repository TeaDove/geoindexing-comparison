package quadtree

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"time"

	qtree "github.com/TeaDove/go-quad-tree"
	"github.com/google/uuid"
)

type CollectionQuadTree struct {
	impl qtree.Qtree[uuid.UUID]
}

func New() index.IndexImpl {
	r := CollectionQuadTree{}

	r.impl = *qtree.NewQtree[uuid.UUID](0, 0, 180, 180, 10)

	return &r
}

func (r *CollectionQuadTree) FromArray(points geo.Points) {
	for _, point := range points {
		r.impl.Insert(qtree.NewPoint(point.Lat, point.Lon, point.ID))
	}
}

func (r *CollectionQuadTree) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()

	r.impl.Insert(qtree.NewPoint(point.Lat, point.Lon, point.ID))

	return time.Since(t0)
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
			point.Lat-radius,
			point.Lon-radius,
			point.Lat+radius,
			point.Lon+radius,
		),
	)
	dur := time.Since(t0)

	return toConcrete(res), dur
}

func (r *CollectionQuadTree) KNNTimed(point geo.Point, n uint64) (geo.Points, time.Duration) {
	t0 := time.Now()
	res := r.impl.KNN(qtree.NewPoint[uuid.UUID](point.Lat, point.Lon, uuid.Nil), int(n))
	dur := time.Since(t0)

	return toConcrete(res), dur
}

func (r *CollectionQuadTree) String() string {
	return r.impl.String()
}
