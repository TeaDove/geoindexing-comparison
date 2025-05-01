package presentation

import (
	"fmt"
	"geoindexing_comparison/backend/service"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

var success = fiber.Map{"success": true} //nolint: gochecknoglobals // Allowed here

func (r *Presentation) getIndexes(c *fiber.Ctx) error {
	return c.JSON(r.service.Indexes)
}

func (r *Presentation) getTasks(c *fiber.Ctx) error {
	return c.JSON(r.service.Tasks)
}

func (r *Presentation) runs(c *fiber.Ctx) error {
	runs, err := r.service.GetRuns(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get runs")
	}

	return c.JSON(runs)
}

func (r *Presentation) runResume(c *fiber.Ctx) error {
	var req service.RunRequest

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	run, err := r.service.AddRun(c.UserContext(), &req, fmt.Sprintf("%s:%s", c.IP(), c.Get(fiber.HeaderUserAgent)))
	if err != nil {
		return errors.Wrap(err, "could not save run")
	}

	return c.JSON(fiber.Map{"runId": run.ID})
}

func (r *Presentation) runReset(c *fiber.Ctx) error {
	err := r.service.StopRuns(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to stop run")
	}

	return c.JSON(success)
}

func (r *Presentation) getStats(c *fiber.Ctx) error {
	var req struct {
		RunID uint64 `json:"runId"`
	}

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	points, err := r.service.GetPoints(c.UserContext(), req.RunID)
	if err != nil {
		return errors.Wrap(err, "failed to get stats")
	}

	if points == nil {
		points = make([]service.Point, 0)
	}

	return c.JSON(points)
}
