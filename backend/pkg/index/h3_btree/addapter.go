package h3_btree

import (
	"geoindexing_comparison/pkg/geo"
	"geoindexing_comparison/pkg/geo/distance_utils"
	"geoindexing_comparison/pkg/index"
	"time"

	"github.com/pkg/errors"
	"github.com/tidwall/btree"
	"github.com/uber/h3-go/v4"
)

type CollectionGeohash struct {
	btree      btree.Map[h3.Cell, geo.Points]
	resolution int
	metric     distance_utils.Metric
}

func Factory(resolution int) func() index.Impl {
	return func() index.Impl {
		return &CollectionGeohash{
			btree:      *btree.NewMap[h3.Cell, geo.Points](2),
			resolution: resolution,
			metric:     distance_utils.MetricHaversine,
		}
	}
}

func (r *CollectionGeohash) hash(point geo.Point) h3.Cell {
	cell, err := h3.LatLngToCell(h3.NewLatLng(point.Lat, point.Lon), r.resolution)
	if err != nil {
		panic(errors.New("failed to build cell"))
	}

	return cell
}

func (r *CollectionGeohash) FromArray(points geo.Points) {
	for _, point := range points {
		r.Insert(point)
	}
}

func (r *CollectionGeohash) ToArray() geo.Points {
	var points geo.Points
	for _, arr := range r.btree.Values() {
		points = append(points, arr...)
	}

	return points
}

func (r *CollectionGeohash) Insert(point geo.Point) {
	v := r.hash(point)
	points, _ := r.btree.Get(v)
	points = append(points, point)

	r.btree.Set(r.hash(point), points)
}

func (r *CollectionGeohash) InsertTimed(point geo.Point) time.Duration {
	t0 := time.Now()

	r.Insert(point)

	return time.Since(t0)
}

func (r *CollectionGeohash) String() string {
	points := r.ToArray()
	return points.String()
}
