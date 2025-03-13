package presentation

import (
	"geoindexing_comparison/service/cases/tasks"
	"geoindexing_comparison/service/index/indexes"
	"github.com/gofiber/fiber/v2"
)

func (r *Presentation) formIndex(c *fiber.Ctx) error {
	return c.Render("index.html", map[string]any{"Indexes": indexes.IndexNames, "tasks": tasks.TaskNames})
}
