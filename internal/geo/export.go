package geo

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"time"

	"geoindexing_comparison/utils"

	"github.com/guregu/null"
	"github.com/rs/zerolog/log"
)

type ExportInput struct {
	Filename null.String
}

func (r PointsColored) MustExport(input *ExportInput) {
	log.Info().Str("status", "exporting.points").Int("count", len(r)).Send()
	if !input.Filename.Valid {
		input.Filename.SetValid("points.csv")
	}
	csvFile, err := os.Create(input.Filename.String)
	defer csvFile.Close()
	utils.Check(err)

	csvwriter := csv.NewWriter(csvFile)
	err = csvwriter.Write([]string{"id", "lat", "lon", "color", "description"})
	utils.Check(err)

	for _, point := range r {
		err = csvwriter.Write(
			[]string{
				point.ID.String(),
				fmt.Sprintf("%f", point.Lat),
				fmt.Sprintf("%f", point.Lon),
				string(point.Color),
				point.Description,
			},
		)
		utils.Check(err)
	}
	csvwriter.Flush()
	log.Info().Str("status", "points.exported").Int("count", len(r)).Send()
}

type DrawInput struct {
	URL           null.String
	OperationType null.String
}

func (r PointsColored) MustDraw(input *DrawInput) {
	log.Info().Str("status", "sending.request.for.draw").Int("count", len(r)).Send()
	if !input.URL.Valid {
		input.URL.SetValid("http://127.0.0.1:8000/draw-points")
	}
	if !input.OperationType.Valid {
		input.OperationType.SetValid("file")
	}

	pointsBytes, err := json.Marshal(r)
	utils.Check(err)
	res, err := http.Post(input.URL.String, "Application/json", bytes.NewBuffer(pointsBytes))
	utils.Check(err)
	if res.StatusCode != http.StatusOK {
		utils.FancyPanic(errors.Errorf("Bad status code: %d", res.StatusCode))
	}
	body, err := io.ReadAll(res.Body)
	utils.Check(err)

	filename := fmt.Sprintf(
		"%s-%s-draw.png",
		input.OperationType.String,
		time.Now().Format("2006-01-02T03:04:05"),
	)
	file, err := os.Create(filename)
	utils.Check(err)
	defer file.Close()

	_, err = file.Write(body)
	utils.Check(err)
	log.Info().Str("status", "draw.completed").Int("count", len(r)).Send()
}
