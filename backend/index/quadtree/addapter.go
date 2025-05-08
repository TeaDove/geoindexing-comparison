package quadtree

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/index"
	"time"

	qtree "github.com/TeaDove/go-quad-tree"
)

type CollectionQuadTree struct {
	impl qtree.Qtree[string]
}

func New() index.Impl {
	r := CollectionQuadTree{}

	r.impl = *qtree.NewQtree[string](0, 0, 180, 180, 10)

	return &r
}

func (r *CollectionQuadTree) FromArray(points geo.Points) {
	for _, point := range points {
		r.impl.Insert(qtree.NewPoint(point.Lat, point.Lon, point.ID))
	}
}

func (r *CollectionQuadTree) ToArray() geo.Points {
	var res geo.Points
	for _, point := range r.impl.Points() {
		res = append(res, geo.Point{Lat: point.X, Lon: point.Y, ID: point.Val})
	}

	return res
}

func (r *CollectionQuadTree) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()

	r.impl.Insert(qtree.NewPoint(point.Lat, point.Lon, point.ID))

	return time.Since(t0)
}

func toConcrete(qtreePoints []qtree.Point[string]) geo.Points {
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

func (r *CollectionQuadTree) BBoxTimed(bottomLeft geo.Point, upperRight geo.Point) (geo.Points, time.Duration) {
	t0 := time.Now()

	res := r.impl.QueryRange(
		qtree.NewBounds[string](
			bottomLeft.Lat,
			bottomLeft.Lon,
			upperRight.Lat-bottomLeft.Lat,
			upperRight.Lon-bottomLeft.Lon,
		),
	)
	dur := time.Since(t0)

	return toConcrete(res), dur
}

func (r *CollectionQuadTree) KNNTimed(origin geo.Point, n uint64) (geo.Points, time.Duration) {
	if n == 0 {
		return geo.Points{}, 0
	}

	t0 := time.Now()
	res := r.impl.KNN(qtree.NewPoint[string](origin.Lat, origin.Lon, ""), int(n))
	dur := time.Since(t0)

	return toConcrete(res), dur
}

func (r *CollectionQuadTree) String() string {
	return r.impl.String()
}
