package presentation

import (
	"fmt"
	"geoindexing_comparison/service/cases"
	"geoindexing_comparison/service/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"time"
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
	runConfig := cases.RunConfig{Start: 100, Stop: 10_000, Step: 100}

	err := c.BodyParser(&runConfig)
	if err != nil {
		return errors.Wrap(err, "failed to parse run config")
	}

	run := &repository.Run{
		ID:        uuid.UUID{},
		CreatedAt: time.Time{},
		CreatedBy: fmt.Sprintf("%s:%s", c.IP(), c.Get(fiber.HeaderUserAgent)),
		Status:    repository.RunStatusPending,
		Start:     runConfig.Start,
		Stop:      runConfig.Stop,
		Step:      runConfig.Step,
	}
	err = run.Indexes.Scan(runConfig.Indexes)
	if err != nil {
		return errors.Wrap(err, "failed to scan run indexes")
	}
	err = run.Tasks.Scan(runConfig.Tasks)
	if err != nil {
		return errors.Wrap(err, "failed to scan run tasks")
	}

	err = r.repository.SaveRun(c.UserContext(), run)
	if err != nil {
		return errors.Wrap(err, "could not save run")
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
