package geo

import (
	"bytes"
	"encoding/json"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"math/rand"
)

// Point represents a geographic coordinate
type Point struct {
	ID    uuid.UUID `json:"id"`
	Color Color     `json:"color"`

	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Points []Point

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

func (r Points) Paint(category Color) Points {
	for idx := range r {
		r[idx].Color = category
	}
	return r
}

func (r Points) PaintPartially(category Color, points Points) Points {
	set := points.ToSet()
	for idx := range r {
		if set.Contains(r[idx].ID) {
			r[idx].Color = category
		}
	}
	return r
}

func (r Points) String() string {
	var buffer bytes.Buffer
	for _, point := range r {
		result, _ := json.Marshal(point)
		buffer.WriteString(string(result))
	}
	return buffer.String()
}

func (r Points) ToSet() mapset.Set[uuid.UUID] {
	result := mapset.NewSet[uuid.UUID]()
	for _, point := range r {
		result.Add(point.ID)
	}
	return result
}
