package presentation

import (
	"geoindexing_comparison/service/cases"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

var success = fiber.Map{"success": true}

func (r *Presentation) getIndexes(c *fiber.Ctx) error {
	return c.JSON(r.runner.Indexes)
}

func (r *Presentation) getTasks(c *fiber.Ctx) error {
	return c.JSON(r.runner.Tasks)
}

func (r *Presentation) runResume(c *fiber.Ctx) error {
	var runConfig cases.RunConfig

	err := c.BodyParser(&runConfig)
	if err != nil {
		return errors.Wrap(err, "failed to parse run config")
	}

	r.resultChannel = r.runner.Run(c.UserContext(), &runConfig)

	return c.JSON(success)
}

func (r *Presentation) runReset(c *fiber.Ctx) error {
	r.runner.Stop(c.UserContext())
	return c.JSON(success)
}
