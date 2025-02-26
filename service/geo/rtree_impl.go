package geo

import (
	rstartreego "github.com/anayks/go-rstar-tree"
	"github.com/dhconnelly/rtreego"
	"github.com/google/uuid"
)

func (r Point) Bounds() rtreego.Rect {
	rect, err := rtreego.NewRectFromPoints(rtreego.Point{r.Lat, r.Lon}, rtreego.Point{r.Lat, r.Lon})
	if err != nil {
		panic(err)
	}

	return rect
}

func (r Point) ToPointForRStarTree() PointForRStarTree {
	return PointForRStarTree{
		ID:  r.ID,
		Lat: r.Lat,
		Lon: r.Lon,
	}
}

func (r PointForRStarTree) ToPoint() Point {
	return Point{
		ID:  r.ID,
		Lat: r.Lat,
		Lon: r.Lon,
	}
}

type PointForRStarTree struct {
	ID uuid.UUID `json:"id"`

	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (r PointForRStarTree) Bounds() *rstartreego.Rect {
	rect, err := rstartreego.NewRectFromPoints(
		rstartreego.Point{r.Lat, r.Lon},
		rstartreego.Point{r.Lat, r.Lon},
	)
	if err != nil {
		panic(err)
	}

	return rect
}
