package geo

import (
	"encoding/csv"
	"fmt"
	"geoindexing_comparison/utils"
	"os"
)

// Point represents a geographic coordinate
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

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

func MustExport(points []Point) {
	csvFile, err := os.Create("points.csv")
	defer csvFile.Close()
	utils.Check(err)

	csvwriter := csv.NewWriter(csvFile)
	err = csvwriter.Write([]string{"lat", "lon"})
	utils.Check(err)

	for _, point := range points {
		err = csvwriter.Write([]string{fmt.Sprintf("%f", point.Lat), fmt.Sprintf("%f", point.Lon)})
		utils.Check(err)
	}
	csvwriter.Flush()
}
