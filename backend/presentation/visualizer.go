package presentation

import (
	"geoindexing_comparison/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (r *Presentation) NewVisualizer(c *fiber.Ctx) error {
	var req service.NewVisualizerInput

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	visualizer, err := r.service.SetVisualizer(&req)
	if err != nil {
		return errors.Wrap(err, "failed to set visualizer")
	}

	return c.JSON(visualizer)
}

func (r *Presentation) GetPoints(c *fiber.Ctx) error {
	return c.JSON(r.service.Visualizer.GetPoints())
}

func (r *Presentation) KNN(c *fiber.Ctx) error {
	var req service.KNNInput

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	points, _ := r.service.Visualizer.KNN(&req)
	return c.JSON(points)
}

func (r *Presentation) RangeSearch(c *fiber.Ctx) error {
	var req service.RangeSearchInput

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	points, _ := r.service.Visualizer.RangeSearch(&req)
	return c.JSON(points)
}
