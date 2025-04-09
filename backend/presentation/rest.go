package presentation

import (
	"fmt"
	"geoindexing_comparison/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

var success = fiber.Map{"success": true}

func (r *Presentation) getIndexes(c *fiber.Ctx) error {
	return c.JSON(r.service.Indexes)
}

func (r *Presentation) getTasks(c *fiber.Ctx) error {
	return c.JSON(r.service.Tasks)
}

func (r *Presentation) runs(c *fiber.Ctx) error {
	runs, err := r.service.GetRuns(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get runs")
	}

	return c.JSON(runs)
}

func (r *Presentation) runResume(c *fiber.Ctx) error {
	var req service.RunRequest

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	run, err := r.service.AddRun(c.UserContext(), &req, fmt.Sprintf("%s:%s", c.IP(), c.Get(fiber.HeaderUserAgent)))
	if err != nil {
		return errors.Wrap(err, "could not save run")
	}

	return c.JSON(fiber.Map{"runId": run.ID})
}

func (r *Presentation) runReset(c *fiber.Ctx) error {
	var req struct {
		RunId int `json:"runId"`
	}

	err := c.BodyParser(&req)
	if err != nil {
		return errors.Wrap(err, "failed to parse request")
	}

	_, err = r.service.StopRun(c.UserContext(), req.RunId)
	if err != nil {
		return errors.Wrap(err, "failed to stop run")
	}

	return c.JSON(success)
}

type Point struct {
	Chart   string  `json:"chart"`
	Dataset string  `json:"dataset"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
}

func (r *Presentation) getPoints(c *fiber.Ctx) error {
	//offset := c.QueryInt("offset")
	//resultsLen := len(r.results)
	//if offset >= resultsLen {
	//	return c.JSON(fiber.Map{})
	//}
	//
	//points := make([]Point, 0, resultsLen)
	//for _, res := range r.results[offset:resultsLen] {
	//	points = append(points, Point{Chart: fmt.Sprintf("%s %s", "", res.Task), Dataset: res.Index, X: float64(res.Amount), Y: float64(res.Durs.Avg())})
	//}

	return c.JSON(nil)
}
