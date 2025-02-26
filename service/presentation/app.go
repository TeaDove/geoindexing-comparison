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
	app.Get("/runs", func(c *fiber.Ctx) error { return c.Render("index.html", nil) })
	app.Get("/plots/ws", websocket.New(r.wsHandle))
	app.Get("/favicon.ico", func(c *fiber.Ctx) error { return filesystem.SendFile(c, httpStaticFS, "favicon.ico") })
	app.Get("/*", func(c *fiber.Ctx) error { return c.Redirect("/") })

	return &Presentation{fiberApp: app}
}

func (r *Presentation) Run(_ context.Context, url string) error {
	return r.fiberApp.Listen(url)
}
