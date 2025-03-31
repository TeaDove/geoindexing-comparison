package presentation

import (
	"geoindexing_comparison/service/cases"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var success = fiber.Map{"success": true}

func (r *Presentation) getIndexes(c *fiber.Ctx) error {
	return c.JSON(r.runner.Indexes)
}

func (r *Presentation) getTasks(c *fiber.Ctx) error {
	return c.JSON(r.runner.Tasks)
}

func (r *Presentation) runResume(c *fiber.Ctx) error {
	r.runner.Stop(c.UserContext())
	runConfig := cases.RunConfig{AmountStart: 100, AmountEnd: 10_000, AmountStep: 100}

	err := c.BodyParser(&runConfig)
	if err != nil {
		return errors.Wrap(err, "failed to parse run config")
	}
	zerolog.Ctx(c.UserContext()).
		Info().
		Interface("run_config", runConfig).
		Msg("run.resuming")

	channel := r.runner.Run(c.UserContext(), &runConfig)
	go r.reader(channel)

	return c.JSON(success)
}

func (r *Presentation) reader(results <-chan cases.Result) {
	for item := range results {
		r.results = append(r.results, item)
	}
}

func (r *Presentation) runReset(c *fiber.Ctx) error {
	r.runner.Stop(c.UserContext())
	r.results = make([]cases.Result, 0, 10_000)
	return c.JSON(success)
}
