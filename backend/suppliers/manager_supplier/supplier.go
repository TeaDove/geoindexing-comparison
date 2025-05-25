package manager_supplier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"geoindexing_comparison/backend/helpers"
	"geoindexing_comparison/backend/schemas"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Supplier struct {
	client http.Client
}

func NewSupplier() *Supplier {
	return &Supplier{client: http.Client{Timeout: 5 * time.Second}}
}

var ErrNotFound = errors.New("job not found")

func (r *Supplier) GetPendingJobs(ctx context.Context) (schemas.Job, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s:%s", helpers.Settings.ManagerURL, "/api/jobs/pending"),
		nil,
	)
	if err != nil {
		return schemas.Job{}, errors.Wrap(err, "failed to create request")
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return schemas.Job{}, errors.Wrap(err, "failed to send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return schemas.Job{}, errors.WithStack(ErrNotFound)
		}

		return schemas.Job{}, errors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var job schemas.Job

	err = json.NewDecoder(resp.Body).Decode(&job)
	if err != nil {
		return schemas.Job{}, errors.Wrap(err, "failed to decode response")
	}

	return job, nil
}

func (r *Supplier) ReportJob(ctx context.Context, jobResult schemas.JobResult) error {
	payload, err := json.Marshal(jobResult)
	if err != nil {
		return errors.Wrap(err, "failed to marshal job result")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s:%s", helpers.Settings.ManagerURL, "/api/jobs/report"),
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
