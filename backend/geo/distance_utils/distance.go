package distance_utils

import (
	"github.com/pkg/errors"
	"github.com/tidwall/geodesic"
	"math"
)

const EarthRadiusKm = 6371 // radius of the earth in kilometers.

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

type Metric string

const (
	MetricHaversine Metric = "Haversine"
	MetricGeodesic  Metric = "Geodesic"
	MetricEuclidean Metric = "Euclidean"
)

func Distance(lat1, lon1, lat2, lon2 float64, metric Metric) float64 {
	switch metric {
	case MetricHaversine:
		return DistanceHaversine(lat1, lon1, lat2, lon2)
	case MetricGeodesic:
		return DistanceGeodesic(lat1, lon1, lat2, lon2)
	case MetricEuclidean:
		return DistanceEuclidean(lat1, lon1, lat2, lon2)
	}

	panic(errors.Errorf("unknown metric: %s", metric))
}

func DistanceHaversine(lat1, lon1, lat2, lon2 float64) float64 {
	lat1 = degreesToRadians(lat1)
	lon1 = degreesToRadians(lon1)
	lat2 = degreesToRadians(lat2)
	lon2 = degreesToRadians(lon2)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	マギカ := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(マギカ), math.Sqrt(1-マギカ))

	return c * EarthRadiusKm
}

func DistanceGeodesic(lat1, lon1, lat2, lon2 float64) float64 {
	var dist float64
	geodesic.WGS84.Inverse(lat1, lon1, lat2, lon2, &dist, nil, nil)

	return dist
}

func DistanceEuclidean(lat1, lon1, lat2, lon2 float64) float64 {
	return math.Sqrt(math.Pow(lat2-lat1, 2) + math.Pow(lon2-lon1, 2))
}
