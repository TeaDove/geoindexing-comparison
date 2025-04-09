package presentation

import (
	"context"
	"geoindexing_comparison/backend/repository"
	"geoindexing_comparison/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/teadove/teasutils/fiber_utils"
)

type Presentation struct {
	fiberApp *fiber.App

	repository *repository.Repository
	service    *service.Service
}

func NewPresentation(service *service.Service, repository *repository.Repository) *Presentation {
	app := fiber.New(fiber.Config{ErrorHandler: fiber_utils.ErrHandler()})
	r := Presentation{fiberApp: app, service: service, repository: repository}

	app.Use(recover2.New())
	app.Use(fiber_utils.MiddlewareLogger())
	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/points", r.getPoints)
	app.Get("/tasks", r.getTasks)
	app.Get("/indexes", r.getIndexes)
	app.Get("/runs", r.runs)
	app.Post("/runs/resume", r.runResume)
	app.Post("/runs/reset", r.runReset)

	return &r
}

func (r *Presentation) Run(_ context.Context, url string) error {
	return r.fiberApp.Listen(url)
}
