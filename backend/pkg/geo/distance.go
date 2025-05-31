// Package geo
//
// Distance functions
package geo

import "geoindexing_comparison/pkg/geo/distance_utils"

func (r Point) Distance(other Point, metric distance_utils.Metric) float64 {
	return metric(r.Lat, r.Lon, other.Lat, other.Lon)
}

func (r Point) DistanceHaversine(other Point) float64 {
	return distance_utils.DistanceHaversine(r.Lat, r.Lon, other.Lat, other.Lon)
}

func (r Point) DistanceGeodesic(other Point) float64 {
	return distance_utils.DistanceGeodesic(r.Lat, r.Lon, other.Lat, other.Lon)
}

func (r Point) DistanceEuclidean(other Point) float64 {
	return distance_utils.DistanceEuclidean(r.Lat, r.Lon, other.Lat, other.Lon)
}

func (r *Points) FindMostDistant(origin Point, metric distance_utils.Metric) float64 {
	var (
		mostDistance float64
		distance     float64
	)

	for _, point := range *r {
		distance = point.Distance(origin, metric)
		if distance > mostDistance {
			mostDistance = distance
		}
	}

	return mostDistance
}
