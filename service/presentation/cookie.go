package presentation

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const cookieIDKey = "cookie_id"

func (r *Presentation) withCookieID(c *fiber.Ctx) error {
	cookieID := c.Cookies(cookieIDKey)
	if cookieID == "" {
		cookieID = uuid.New().String()
	}

	c.Locals(cookieIDKey, cookieID)
	c.Cookie(&fiber.Cookie{Name: cookieIDKey, Value: cookieID})

	return c.Next()
}

func (r *Presentation) mustGetCookieID(c *fiber.Ctx) string {
	return c.Locals(cookieIDKey).(string)
}

func (r *Presentation) mustGetCookieIDFromWS(c *websocket.Conn) string {
	return c.Locals(cookieIDKey).(string)
}
