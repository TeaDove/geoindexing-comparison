package manager_presentation

import (
	"geoindexing_comparison/backend/schemas"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (r *Presentation) getPendingJobs(c *fiber.Ctx) error {
	job, ok := r.managerService.GetPendingJob()
	if !ok {
		c.Status(fiber.StatusNotFound)
		return c.SendString("No pending jobs available")
	}

	return c.JSON(job)
}

func (r *Presentation) reportJob(c *fiber.Ctx) error {
	var jobResult schemas.JobResult
	err := c.BodyParser(&jobResult)
	if err != nil {
		return errors.Wrap(err, "failed to parse request body")
	}

	err = r.managerService.ReportJob(c.UserContext(), &jobResult)
	if err != nil {
		return errors.Wrap(err, "failed to report")
	}

	return nil
}
