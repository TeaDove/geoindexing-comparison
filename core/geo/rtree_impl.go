package geo

import (
	"geoindexing_comparison/core/utils"
	"github.com/dhconnelly/rtreego"
)

func (r Point) Bounds() rtreego.Rect {
	rect, err := rtreego.NewRectFromPoints(rtreego.Point{r.Lat, r.Lon}, rtreego.Point{r.Lat, r.Lon})
	utils.Check(err)
	return rect
}
