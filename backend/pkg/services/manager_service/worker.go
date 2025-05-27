package manager_service

import (
	"context"
	"geoindexing_comparison/pkg/repositories/manager_repository"
	"geoindexing_comparison/pkg/schemas"
	"github.com/guregu/null/v6"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/time_utils"
	"time"

	"github.com/pkg/errors"
)

func (r *Service) GetPendingJob() (schemas.Job, bool) {
	r.jobMu.Lock()
	defer r.jobMu.Unlock()

	if len(r.jobs) == 0 || r.jobIdx >= len(r.jobs) {
		return schemas.Job{}, false
	}

	job := r.jobs[r.jobIdx]

	return job, true
}

func (r *Service) ReportJob(ctx context.Context, jobResult *schemas.JobResult) error {
	r.jobMu.Lock()
	defer r.jobMu.Unlock()

	err := r.repository.SaveStats(ctx, &manager_repository.Stats{
		RunID:  r.currentRun.ID,
		Idx:    r.jobIdx,
		Index:  jobResult.Index,
		Task:   jobResult.Task,
		Amount: jobResult.Amount,
		Durs:   jobResult.Durs,
	})
	if err != nil {
		return errors.Wrap(err, "failed to save stats")
	}

	r.jobIdx++
	if r.jobIdx >= len(r.jobs) {
		r.currentRun.CompletedAt = null.NewTime(time.Now(), true)
		r.currentRun.Status = manager_repository.RunStatusCompleted

		err = r.repository.SaveRun(ctx, r.currentRun)
		if err != nil {
			return errors.Wrap(err, "failed to save")
		}

		zerolog.Ctx(ctx).Info().
			Str("elapsed", time_utils.RoundDuration(time.Since(r.currentRun.CreatedAt))).
			Msg("run.done")
	}

	return nil
}
