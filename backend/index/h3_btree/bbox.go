package h3_btree

import (
	"geoindexing_comparison/backend/geo"
	"time"

	"github.com/pkg/errors"
	"github.com/uber/h3-go/v4"
)

func (r *CollectionGeohash) getMany(hashed []h3.Cell) geo.Points {
	var (
		points      geo.Points
		foundPoints geo.Points
	)

	for _, hash := range hashed {
		foundPoints, _ = r.btree.Get(hash)
		points = append(points, foundPoints...)
	}

	return points
}

func (r *CollectionGeohash) bbox(bottomLeft geo.Point, upperRight geo.Point) geo.Points {
	cells, err := h3.PolygonToCellsExperimental(
		h3.GeoPolygon{
			GeoLoop: h3.GeoLoop{
				h3.NewLatLng(bottomLeft.Lat, bottomLeft.Lon),
				h3.NewLatLng(upperRight.Lat, bottomLeft.Lon),
				h3.NewLatLng(upperRight.Lat, upperRight.Lon),
				h3.NewLatLng(bottomLeft.Lat, upperRight.Lon),
				h3.NewLatLng(bottomLeft.Lat, bottomLeft.Lon),
			},
		},
		r.resolution,
		h3.ContainmentOverlappingBbox,
	)
	if err != nil {
		panic(errors.Wrap(err, "failed to convert bbox to cells"))
	}

	var points geo.Points

	posiblePoints := r.getMany(cells)

	for _, point := range posiblePoints {
		if point.InsideBBox(bottomLeft, upperRight) {
			points = append(points, point)
		}
	}

	return points
}

func (r *CollectionGeohash) BBoxTimed(bottomLeft geo.Point, upperRight geo.Point) (geo.Points, time.Duration) {
	t0 := time.Now()

	return r.bbox(bottomLeft, upperRight), time.Since(t0)
}
