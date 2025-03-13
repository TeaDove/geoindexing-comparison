package presentation

import (
	"context"
	"geoindexing_comparison/service/presentation/static"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/teadove/teasutils/utils/settings_utils"
	"net/http"
)

type Presentation struct {
	fiberApp *fiber.App
}

func NewPresentation() *Presentation {
	httpStaticFS := http.FS(static.FS)
	sendFile := func(name string) func(c *fiber.Ctx) error {
		return func(c *fiber.Ctx) error { return filesystem.SendFile(c, httpStaticFS, name) }
	}

	renderEngine := html.NewFileSystem(httpStaticFS, "")
	if !settings_utils.BaseSettings.Release {
		renderEngine.Reload(true)
	}

	app := fiber.New(fiber.Config{Views: renderEngine, ErrorHandler: errHandler})
	r := Presentation{fiberApp: app}

	app.Use(r.withCookieID)
	app.Use(r.logCtxMiddleware)
	app.Use(logger.New())

	app.Get("/", r.formIndex)
	app.Get("/index.css", sendFile("index.css"))
	app.Get("/index.js", sendFile("index.js"))
	app.Get("/favicon.ico", sendFile("favicon.ico"))
	app.Get("/plots/ws", websocket.New(r.wsHandle))
	app.Post("/runs/resume", r.runResume)
	app.Post("/runs/reset", r.runReset)
	app.Get("/*", func(c *fiber.Ctx) error { return c.Redirect("/") })

	return &Presentation{fiberApp: app}
}

func (r *Presentation) Run(_ context.Context, url string) error {
	return r.fiberApp.Listen(url)
}
