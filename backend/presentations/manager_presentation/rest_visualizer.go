package manager_presentation

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/services/visualizer_service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func sendPoints(c *fiber.Ctx, points geo.Points, dur time.Duration) error {
	c.Set("X-Duration-Microseconds", strconv.Itoa(int(dur.Microseconds())))
	return c.JSON(points.GeoJSON())
}

func (r *Presentation) NewVisualizer(c *fiber.Ctx) error {
	var req visualizer_service.NewVisualizerInput

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	err = r.visualizerService.SetVisualizer(&req)
	if err != nil {
		return errors.Wrap(err, "failed to set visualizer")
	}

	points := r.visualizerService.GetPoints()

	return c.JSON(points.GeoJSON())
}

func (r *Presentation) GetPoints(c *fiber.Ctx) error {
	points := r.visualizerService.GetPoints()
	return c.JSON(points.GeoJSON())
}

func (r *Presentation) KNN(c *fiber.Ctx) error {
	var req visualizer_service.KNNInput

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	points, dur := r.visualizerService.KNN(&req)

	return sendPoints(c, points, dur)
}

func (r *Presentation) BBox(c *fiber.Ctx) error {
	var req visualizer_service.BBoxInput

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	points, dur := r.visualizerService.BBox(&req)

	return sendPoints(c, points, dur)
}
