package geo

import (
	"encoding/json"
	"geoindexing_comparison/backend/helpers"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"strings"

	"golang.org/x/exp/slices"

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

func (r Point) GeoJSON() *geojson.Feature {
	feature := geojson.NewFeature(orb.Point{r.Lon, r.Lat})
	feature.Properties["ID"] = r.ID

	return feature
}

func (r *Points) GetRandomPoint() Point {
	return (*r)[helpers.RNG.IntN(len(*r))] //nolint: gosec // Allowed here
}

func (r *Points) String() string {
	byteArray, err := json.Marshal(r)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal points"))
	}

	return string(byteArray)
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
