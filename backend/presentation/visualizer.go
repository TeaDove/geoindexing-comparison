package presentation

import (
	"geoindexing_comparison/backend/geo"
	"geoindexing_comparison/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func sendPoints(c *fiber.Ctx, points geo.Points, dur time.Duration) error {
	c.Set("X-Duration-Microseconds", strconv.Itoa(int(dur.Microseconds())))
	return c.JSON(points.GeoJSON())
}

func (r *Presentation) NewVisualizer(c *fiber.Ctx) error {
	var req service.NewVisualizerInput

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	visualizer, err := r.service.SetVisualizer(c.UserContext(), &req)
	if err != nil {
		return errors.Wrap(err, "failed to set visualizer")
	}

	points := visualizer.GetPoints()
	return c.JSON(points.GeoJSON())
}

func (r *Presentation) GetPoints(c *fiber.Ctx) error {
	points := r.service.Visualizer.GetPoints()
	return c.JSON(points.GeoJSON())
}

func (r *Presentation) KNN(c *fiber.Ctx) error {
	var req service.KNNInput

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	points, dur := r.service.Visualizer.KNN(&req)
	return sendPoints(c, points, dur)
}

func (r *Presentation) BBox(c *fiber.Ctx) error {
	var req service.BBoxInput

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	points, dur := r.service.Visualizer.BBox(&req)
	return sendPoints(c, points, dur)
}
