package manager_presentation

import (
	"geoindexing_comparison/backend/services/builder_service"
	"geoindexing_comparison/backend/services/manager_service"
	"geoindexing_comparison/backend/services/visualizer_service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/teadove/teasutils/fiber_utils"
)

type Presentation struct {
	fiberApp *fiber.App

	managerService    *manager_service.Service
	visualizerService *visualizer_service.Service
	builderService    *builder_service.Service
}

func NewPresentation(
	managerService *manager_service.Service,
	visualizerService *visualizer_service.Service,
	builderService *builder_service.Service,
) *Presentation {
	app := fiber.New(fiber.Config{ErrorHandler: fiber_utils.ErrHandler()})
	r := Presentation{
		fiberApp:          app,
		managerService:    managerService,
		builderService:    builderService,
		visualizerService: visualizerService,
	}

	app.Use(recover2.New())
	app.Use(fiber_utils.MiddlewareLogger())
	app.Use(cors.New(cors.ConfigDefault))

	api := app.Group("/api")

	api.Get("/tasks", r.getTasks)
	api.Get("/indexes", r.getIndexes)
	api.Get("/runs", r.runs)
	api.Post("/runs/stats", r.getStats)
	api.Post("/runs/resume", r.runResume)
	api.Post("/runs/reset", r.runReset)

	api.Post("/visualizer", r.NewVisualizer)
	api.Get("/visualizer/points", r.GetPoints)
	api.Post("/visualizer/knn", r.KNN)
	api.Post("/visualizer/bbox", r.BBox)

	api.Get("/jobs/pending", r.getPendingJobs)
	api.Post("/jobs/report", r.reportJob)

	return &r
}

func (r *Presentation) Run(url string) error {
	return r.fiberApp.Listen(url)
}
