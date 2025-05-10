package geo

import (
	"github.com/tidwall/geodesic"
	"math"
)

const earthRadiusKm = 6371 // radius of the earth in kilometers.

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

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * earthRadiusKm
}

func distanceGeodesic(lat1, lon1, lat2, lon2 float64) float64 {
	var dist float64
	geodesic.WGS84.Inverse(lat1, lon1, lat2, lon2, &dist, nil, nil)

	return dist
}

var distanceFunc = distanceHaversine

//var distanceFunc = distanceGeodesic

func Distance(a, b Point) float64 {
	return distanceFunc(a.Lat, a.Lon, b.Lat, b.Lon)
}

func (r Point) DistanceTo(other Point) float64 {
	return Distance(r, other)
}

func (r Point) DistanceToLatLng(lat, lon float64) float64 {
	return distanceHaversine(r.Lat, r.Lon, lat, lon)
}
