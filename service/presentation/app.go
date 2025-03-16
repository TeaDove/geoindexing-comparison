package presentation

import (
	"context"
	"geoindexing_comparison/service/cases"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
)

type Presentation struct {
	fiberApp      *fiber.App
	runner        *cases.Runner
	resultChannel <-chan cases.Result
}

func NewPresentation(runner *cases.Runner) *Presentation {
	app := fiber.New(fiber.Config{ErrorHandler: errHandler})
	r := Presentation{fiberApp: app, runner: runner}

	app.Use(r.withCookieID)
	app.Use(r.logCtxMiddleware)
	app.Use(logger.New())

	app.Get("/plots/ws", websocket.New(r.wsHandle))
	app.Post("/tasks", r.getTasks)
	app.Post("/indexes", r.getIndexes)
	app.Post("/runs/resume", r.runResume)
	app.Post("/runs/reset", r.runReset)
	app.Use("/*", filesystem.New(filesystem.Config{
		Root: http.Dir("./frontend"),
	}))

	return &Presentation{fiberApp: app}
}

func (r *Presentation) Run(_ context.Context, url string) error {
	return r.fiberApp.Listen(url)
}
