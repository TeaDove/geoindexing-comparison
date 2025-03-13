package presentation

import "github.com/gofiber/fiber/v2"

var success = fiber.Map{"success": true}

type ResumeRequest struct {
}

func (r *Presentation) runResume(c *fiber.Ctx) error {
	println(string(c.Body()))
	return c.JSON(success)
}

func (r *Presentation) runReset(c *fiber.Ctx) error {
	println(string(c.Body()))
	return c.JSON(success)
}
