package presentation

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/logger_utils"
	"net/http"
)

func errHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if code >= http.StatusInternalServerError {
		zerolog.Ctx(c.UserContext()).
			Error().
			Stack().Err(err).
			Int("code", code).
			Msg("http.internal.error")
	} else {
		zerolog.Ctx(c.UserContext()).
			Warn().
			Err(err).
			Int("code", code).
			Msg("http.client.error")
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return c.Status(code).SendString(err.Error())
}

const logCtxKey = "logCtx"

func (r *Presentation) logCtxMiddleware(c *fiber.Ctx) error {
	ctx := logger_utils.AddLoggerToCtx(c.UserContext())
	ctx = logger_utils.WithValue(ctx,
		"app_method",
		fmt.Sprintf(
			"%s %s",
			c.Method(),
			c.Path(),
		),
	)

	ctx = logger_utils.WithValue(ctx, "ip", c.IP())
	ctx = logger_utils.WithValue(ctx, "user_agent", c.Get("User-Agent"))
	ctx = logger_utils.WithValue(ctx, "cookie_id", r.mustGetCookieID(c))

	c.SetUserContext(ctx)
	c.Locals(logCtxKey, ctx)

	return c.Next()
}

func (r *Presentation) mustGetLogContext(c interface {
	Locals(key string, value ...interface{}) interface{}
}) context.Context {
	return c.Locals(logCtxKey).(context.Context)
}
