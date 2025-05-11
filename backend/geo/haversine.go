// Package geo
//
// Distance functions
package geo

import (
	"github.com/tidwall/geodesic"
	"math"
)

const earthRadiusKm = 6371 // radius of the earth in kilometers.

func distanceEuclidean(lat1, lon1, lat2, lon2 float64) float64 {
	return math.Sqrt(math.Pow(lat2-lat1, 2) + math.Pow(lon2-lon1, 2))
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func distanceHaversine(lat1, lon1, lat2, lon2 float64) float64 {
	lat1 = degreesToRadians(lat1)
	lon1 = degreesToRadians(lon1)
	lat2 = degreesToRadians(lat2)
	lon2 = degreesToRadians(lon2)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	マギカ := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(マギカ), math.Sqrt(1-マギカ))

	return c * earthRadiusKm
}

func distanceGeodesic(lat1, lon1, lat2, lon2 float64) float64 {
	var dist float64
	geodesic.WGS84.Inverse(lat1, lon1, lat2, lon2, &dist, nil, nil)

	return dist
}

func (r Point) DistanceHaversine(other Point) float64 {
	return distanceHaversine(r.Lat, r.Lon, other.Lat, other.Lon)
}

func (r Point) DistanceGeodesic(other Point) float64 {
	return distanceGeodesic(r.Lat, r.Lon, other.Lat, other.Lon)
}

func (r Point) DistanceEuclidean(other Point) float64 {
	return distanceEuclidean(r.Lat, r.Lon, other.Lat, other.Lon)
}
