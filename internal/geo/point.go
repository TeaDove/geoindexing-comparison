package geo

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"geoindexing_comparison/utils"
	"github.com/google/uuid"
	"os"
)

// Point represents a geographic coordinate
type Point struct {
	ID       uuid.UUID `json:"id"`
	Category Category  `json:"category"`

	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Points []Point

type Category string

const (
	EMPTY Category = "EMPTY"
	FOUND          = "FOUND"
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

func (r Points) MustExport(filename string) {
	csvFile, err := os.Create(filename)
	defer csvFile.Close()
	utils.Check(err)

	csvwriter := csv.NewWriter(csvFile)
	err = csvwriter.Write([]string{"id", "lat", "lon", "category"})
	utils.Check(err)

	for _, point := range r {
		err = csvwriter.Write([]string{point.ID.String(), fmt.Sprintf("%f", point.Lat), fmt.Sprintf("%f", point.Lon), string(point.Category)})
		utils.Check(err)
	}
	csvwriter.Flush()
}

func (r Points) Paint(category Category) {
	for idx := range r {
		r[idx].Category = category
	}
}

func (r Points) String() string {
	var buffer bytes.Buffer
	for _, point := range r {
		result, _ := json.Marshal(point)
		buffer.WriteString(string(result))
	}
	return buffer.String()
}
