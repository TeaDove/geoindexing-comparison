package geo

import (
	"encoding/json"
	"geoindexing_comparison/utils"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"math/rand"
)

// Point represents a geographic coordinate
type Point struct {
	ID uuid.UUID `json:"id"`

	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type PointExtended struct {
	Point
	Color       Color  `json:"color"`
	Description string `json:"description"`
}

type Points []Point

type PointsExtended []PointExtended

type Color string

const (
	Blue   Color = "Blue"
	Green        = "Green"
	Yellow       = "Yellow"
	Red          = "Red"
)

func (r Point) Dimensions() int {
	return 2
}

func (r Point) Dimension(i int) float64 {
	switch i {
	case 0:
		return r.Lat
	default:
		return r.Lon
	}
}

func (r Points) GetRandomPoint() Point {
	return r[rand.Intn(len(r))]
}

func (r PointsExtended) Paint(category Color) PointsExtended {
	for idx := range r {
		r[idx].Color = category
	}
	return r
}

func (r PointsExtended) PaintPartially(category Color, points Points) PointsExtended {
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

func (r Points) ToSet() mapset.Set[uuid.UUID] {
	result := mapset.NewSet[uuid.UUID]()
	for _, point := range r {
		result.Add(point.ID)
	}
	return result
}

func (r Points) ToPointExtended() PointsExtended {
	var result = make(PointsExtended, len(r))
	for idx := range r {
		result[idx] = PointExtended{
			Point: r[idx],
			Color: Blue,
		}
	}
	return result
}
