// Package geo
//
//	Point, Points method
package geo

import (
	"encoding/json"
	"geoindexing_comparison/backend/geo/distance_utils"
	"geoindexing_comparison/backend/helpers"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"golang.org/x/exp/slices"
	"math"
	"strings"

	"github.com/pkg/errors"

	"github.com/mmcloughlin/geohash"

	mapset "github.com/deckarep/golang-set/v2"
)

// Point represents a geographic coordinate.
type Point struct {
	ID string `json:"id"`

	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type Points []Point

func NewPoint(lat float64, lng float64) Point {
	return Point{
		ID:  helpers.ID(),
		Lat: lat,
		Lon: lng,
	}
}

func (r Point) Geohash(bits uint) uint64 {
	return geohash.EncodeIntWithPrecision(r.Lat, r.Lon, bits)
}

func (r Point) GeohashString(chars uint) string {
	return geohash.EncodeWithPrecision(r.Lat, r.Lon, chars)
}

func (r Point) GeoJSON() *geojson.Feature {
	feature := geojson.NewFeature(orb.Point{r.Lon, r.Lat})
	feature.Properties["ID"] = r.ID

	return feature
}

func (r Point) InsideBBox(bottomLeft Point, upperRight Point) bool {
	return bottomLeft.Lat < r.Lat && bottomLeft.Lon < r.Lon && r.Lat < upperRight.Lat && r.Lon < upperRight.Lon
}

func (r Point) AddLatitude(dvKM float64) Point {
	r.Lat = r.Lat + (dvKM/distance_utils.EarthRadiusKm)*(180/math.Pi)
	return r
}

func (r Point) AddLongitude(dvKM float64) Point {
	r.Lon = r.Lon + (dvKM/distance_utils.EarthRadiusKm)*(180/math.Pi)/math.Cos(r.Lat*math.Pi/180)
	return r
}

func (r *Points) GetRandomPoint() Point {
	return (*r)[helpers.RNG.IntN(len(*r))] //nolint: gosec // Allowed here
}

func (r *Points) FindCorners() (Point, Point) {
	bottomLeft, upperRight := (*r)[0], (*r)[0]
	for _, point := range *r {
		if point.Lat < bottomLeft.Lat && point.Lon < bottomLeft.Lon {
			bottomLeft = point
		}

		if point.Lat > upperRight.Lat && point.Lon > upperRight.Lon {
			upperRight = point
		}
	}

	return bottomLeft, upperRight
}

func (r *Points) Center() (float64, float64) {
	bottomLeft, upperRight := r.FindCorners()
	return (bottomLeft.Lat + upperRight.Lat) / 2.0, (bottomLeft.Lon + upperRight.Lon) / 2.0
}

func (r *Points) String() string {
	byteArray, err := json.Marshal(r)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal points"))
	}

	return string(byteArray)
}

func (r *Points) IDs() string {
	var ids []string
	for _, point := range *r {
		ids = append(ids, point.ID)
	}

	slices.Sort(ids)
	return strings.Join(ids, ",")
}

func (r *Points) Delete(pointID string) {
	for idx, point := range *r {
		if pointID == point.ID {
			*r = append((*r)[:idx], (*r)[idx+1:]...)
			return
		}
	}
}

func (r *Points) ToSet() mapset.Set[string] {
	result := mapset.NewSet[string]()
	for _, point := range *r {
		result.Add(point.ID)
	}

	return result
}

func (r *Points) GeoJSON() *geojson.FeatureCollection {
	featureCollection := geojson.NewFeatureCollection()
	for _, point := range *r {
		featureCollection.Append(point.GeoJSON())
	}

	return featureCollection
}

func (r *Points) SortByID() {
	slices.SortFunc(*r, func(a, b Point) int {
		return strings.Compare(a.ID, b.ID)
	})
}

func (r *Points) SortByDistance(origin Point) {
	slices.SortFunc(*r, func(a, b Point) int {
		if a.DistanceHaversine(origin) < b.DistanceHaversine(origin) {
			return -1
		}
		return 1
	})
}

func (r *Points) GetClosestViaSort(origin Point, n int) Points {
	if n > len(*r) {
		return *r
	}

	type dist struct {
		idx  int
		dist float64
	}

	knnMatrix := make([]dist, 0, len(*r))
	for idx, indexPoint := range *r {
		knnMatrix = append(knnMatrix, dist{idx: idx, dist: indexPoint.DistanceHaversine(origin)})
	}

	slices.SortFunc(knnMatrix, func(a, b dist) int {
		if a.dist < b.dist {
			return -1
		}

		return 0
	})

	result := make(Points, n)

	for idx := range n {
		result[idx] = (*r)[knnMatrix[idx].idx]
	}

	return result
}

func (r *Points) Equal(other Points) bool {
	return r.ToSet().Equal(other.ToSet())
}

func (r *Points) EqualMany(other []Points) bool {
	for _, otherPoint := range other {
		if !r.Equal(otherPoint) {
			return false
		}
	}

	return true
}
