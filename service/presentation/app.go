package presentation

import (
	"context"
	"geoindexing_comparison/service/cases"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/teadove/teasutils/fiber_utils"
	"net/http"
)

type Presentation struct {
	fiberApp      *fiber.App
	runner        *cases.Runner
	resultChannel <-chan cases.Result
}

func NewPresentation(runner *cases.Runner, frontend http.FileSystem) *Presentation {
	app := fiber.New(fiber.Config{ErrorHandler: fiber_utils.ErrHandler()})
	r := Presentation{fiberApp: app, runner: runner}

	app.Use(fiber_utils.MiddlewareLogger(&fiber_utils.LogCtxConfig{}))

	app.Get("/plots/ws", websocket.New(r.wsHandle))
	app.Get("/tasks", r.getTasks)
	app.Get("/indexes", r.getIndexes)
	app.Post("/runs/resume", r.runResume)
	app.Post("/runs/reset", r.runReset)
	app.Use("/*", filesystem.New(filesystem.Config{Browse: true, Root: frontend}))

	return &Presentation{fiberApp: app}
}

func (r *Presentation) Run(_ context.Context, url string) error {
	return r.fiberApp.Listen(url)
}
