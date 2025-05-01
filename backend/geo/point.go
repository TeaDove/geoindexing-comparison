package geo

import (
	"encoding/json"
	"math/rand"

	"github.com/pkg/errors"

	"github.com/mmcloughlin/geohash"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

// Point represents a geographic coordinate.
type Point struct {
	ID uuid.UUID `json:"id"`

	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type Points []Point

func NewPoint(lat float64, lng float64) Point {
	return Point{
		ID:  uuid.New(),
		Lat: lat,
		Lon: lng,
	}
}

func (r Point) Geohash() string {
	return geohash.Encode(r.Lat, r.Lon)
}

func (r *Points) GetRandomPoint() Point {
	return (*r)[rand.Intn(len(*r))] //nolint: gosec // Allowed here
}

func (r *Points) String() string {
	byteArray, err := json.Marshal(r)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal points"))
	}

	return string(byteArray)
}

func (r *Points) Delete(pointID uuid.UUID) {
	for idx, point := range *r {
		if pointID == point.ID {
			*r = append((*r)[:idx], (*r)[idx+1:]...)
			return
		}
	}
}

func (r *Points) ToSet() mapset.Set[uuid.UUID] {
	result := mapset.NewSet[uuid.UUID]()
	for _, point := range *r {
		result.Add(point.ID)
	}

	return result
}
