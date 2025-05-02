package geo

import (
	rstartreego "github.com/anayks/go-rstar-tree"
	"github.com/dhconnelly/rtreego"
)

func (r Point) Bounds() rtreego.Rect {
	rect, err := rtreego.NewRectFromPoints(rtreego.Point{r.Lat, r.Lon}, rtreego.Point{r.Lat, r.Lon})
	if err != nil {
		panic(err)
	}

	return rect
}

func (r Point) ToPointForRStarTree() PointForRStarTree {
	return PointForRStarTree(r)
}

func (r PointForRStarTree) ToPoint() Point {
	return Point(r)
}

type PointForRStarTree struct {
	ID string `json:"id"`

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
