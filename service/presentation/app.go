package presentation

import (
	"context"
	"geoindexing_comparison/service/cases"
	"geoindexing_comparison/service/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/teadove/teasutils/fiber_utils"
)

type Presentation struct {
	fiberApp   *fiber.App
	repository *repository.Repository

	runner  *cases.Runner
	results []cases.Result
}

func NewPresentation(runner *cases.Runner, repository *repository.Repository) *Presentation {
	app := fiber.New(fiber.Config{ErrorHandler: fiber_utils.ErrHandler()})
	r := Presentation{fiberApp: app, runner: runner, results: make([]cases.Result, 0, 10_000), repository: repository}

	app.Use(fiber_utils.MiddlewareLogger())
	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/points", r.getPoints)
	app.Get("/tasks", r.getTasks)
	app.Get("/indexes", r.getIndexes)
	app.Post("/runs/resume", r.runResume)
	app.Post("/runs/reset", r.runReset)

	return &r
}

func (r *Presentation) Run(_ context.Context, url string) error {
	return r.fiberApp.Listen(url)
}
