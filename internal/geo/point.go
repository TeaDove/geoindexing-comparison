package geo

import (
	"encoding/json"
	"math/rand"

	"geoindexing_comparison/utils"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

// Point represents a geographic coordinate
type Point struct {
	ID uuid.UUID `json:"id"`

	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type Points []Point

type PointColored struct {
	Point
	Color       Color  `json:"color"`
	Description string `json:"description"`
}

type PointsColored []PointColored

type Color string

const (
	Blue   Color = "Blue"
	Green  Color = "Green"
	Yellow Color = "Yellow"
	Red    Color = "Red"
)

func (r Points) GetRandomPoint() Point {
	return r[rand.Intn(len(r))]
}

func (r PointsColored) Paint(category Color) PointsColored {
	for idx := range r {
		r[idx].Color = category
	}
	return r
}

func (r PointsColored) PaintPartially(category Color, points Points) PointsColored {
	set := points.ToSet()
	for idx := range r {
		if set.Contains(r[idx].ID) {
			r[idx].Color = category
		}
	}
	return r
}

func (r Points) String() string {
	byteArray, err := json.Marshal(r)
	utils.Check(err)
	return string(byteArray)
}

func (r Points) Delete(pointID uuid.UUID) {
	for idx, point := range r {
		if pointID == point.ID {
			r = append(r[:idx], r[idx+1:]...)
			return
		}
	}
}

func (r Points) ToSet() mapset.Set[uuid.UUID] {
	result := mapset.NewSet[uuid.UUID]()
	for _, point := range r {
		result.Add(point.ID)
	}
	return result
}

func (r Points) ToPointExtended() PointsColored {
	result := make(PointsColored, len(r))
	for idx := range r {
		result[idx] = PointColored{
			Point: r[idx],
			Color: Blue,
		}
	}

	return result
}
