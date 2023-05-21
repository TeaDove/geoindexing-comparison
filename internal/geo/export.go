package geo

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"geoindexing_comparison/utils"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
)

func (r Points) MustExport(filename string) {
	log.Info().Str("status", "exporting.points").Int("count", len(r)).Send()
	csvFile, err := os.Create(filename)
	defer csvFile.Close()
	utils.Check(err)

	csvwriter := csv.NewWriter(csvFile)
	err = csvwriter.Write([]string{"id", "lat", "lon", "color"})
	utils.Check(err)

	for _, point := range r {
		err = csvwriter.Write([]string{point.ID.String(), fmt.Sprintf("%f", point.Lat), fmt.Sprintf("%f", point.Lon), string(point.Color)})
		utils.Check(err)
	}
	csvwriter.Flush()
	log.Info().Str("status", "points.exported").Int("count", len(r)).Send()
}

func (r Points) MustDraw() {
	log.Info().Str("status", "sending.request.for.draw").Int("count", len(r)).Send()
	file, err := os.Create("file.png")
	utils.Check(err)
	defer file.Close()

	pointsBytes, err := json.Marshal(r)
	utils.Check(err)

	res, err := http.Post("http://localhost:8000/draw", "Application/json", bytes.NewBuffer(pointsBytes))
	utils.Check(err)

	body, err := io.ReadAll(res.Body)
	utils.Check(err)
	_, err = file.Write(body)
	utils.Check(err)
	log.Info().Str("status", "draw.completed").Int("count", len(r)).Send()
}
