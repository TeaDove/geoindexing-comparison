package geo

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"geoindexing_comparison/utils"
	"github.com/guregu/null"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
	"time"
)

type ExportConfig struct {
	Filename null.String
}

func (r Points) MustExport(config ExportConfig) {
	log.Info().Str("status", "exporting.points").Int("count", len(r)).Send()
	if !config.Filename.Valid {
		config.Filename.SetValid("points.csv")
	}
	csvFile, err := os.Create(config.Filename.String)
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

type DrawConfig struct {
	URL           null.String
	OperationType null.String
}

func (r Points) MustDraw(config DrawConfig) {
	log.Info().Str("status", "sending.request.for.draw").Int("count", len(r)).Send()
	if !config.URL.Valid {
		config.URL.SetValid("http://127.0.0.1:8000/draw-points")
	}
	if !config.OperationType.Valid {
		config.OperationType.SetValid("file")
	}

	pointsBytes, err := json.Marshal(r)
	utils.Check(err)
	// TODO move to settings
	res, err := http.Post(config.URL.String, "Application/json", bytes.NewBuffer(pointsBytes))
	utils.Check(err)
	if res.StatusCode != http.StatusOK {
		utils.Check(errors.New(fmt.Sprintf("Bad status code: %d", res.StatusCode)))
	}
	body, err := io.ReadAll(res.Body)
	utils.Check(err)

	filename := fmt.Sprintf("%s-%s-draw.png", config.OperationType.String, time.Now().Format("2006-01-02T03:04:05"))
	file, err := os.Create(filename)
	utils.Check(err)
	defer file.Close()

	_, err = file.Write(body)
	utils.Check(err)
	log.Info().Str("status", "draw.completed").Int("count", len(r)).Send()
}
