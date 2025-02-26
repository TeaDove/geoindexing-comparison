package presentation

import (
	"geoindexing_comparison/service/index/indexes"
	"github.com/gofiber/fiber/v2"
	"maps"
	"slices"
)

func (r *Presentation) formIndex(c *fiber.Ctx) error {
	return c.Render("index.html", map[string]any{"Indexes": slices.Collect(maps.Keys(indexes.NameToNewIndex))})
}
