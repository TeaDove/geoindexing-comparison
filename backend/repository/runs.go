package repository

import (
	"context"
	"github.com/guregu/null/v6"
	"github.com/pkg/errors"
	"time"
)

type Run struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   string    `json:"createdBy"`
	CompletedAt null.Time `json:"completedAt"`

	Status RunStatus

	Indexes []string `gorm:"serializer:json" json:"indexes"`
	Tasks   []string `gorm:"serializer:json" json:"tasks"`
	Start   uint64   `json:"start"`
	Stop    uint64   `json:"stop"`
	Step    uint64   `json:"step"`
}

type RunStatus string

const (
	RunStatusPending   RunStatus = "PENDING"
	RunStatusCompleted RunStatus = "COMPLETED"
	RunStatusCancelled RunStatus = "CANCELLED"
)

func (r *Repository) SaveRun(ctx context.Context, v *Run) error {
	err := r.db.WithContext(ctx).Save(v).Error
	if err != nil {
		return errors.Wrap(err, "failed to save run")
	}

	return nil
}

func (r *Repository) StopRuns(ctx context.Context) error {
	err := r.db.WithContext(ctx).
		Model(Run{}).
		Where("status = ?", RunStatusPending).
		Update("status", RunStatusCancelled).
		Error
	if err != nil {
		return errors.Wrap(err, "failed to save run")
	}

	return nil
}

func (r *Repository) GetPending(ctx context.Context) ([]Run, error) {
	var runs []Run
	err := r.db.
		WithContext(ctx).
		Where("status = ?", RunStatusPending).
		Order("created_at asc").
		Find(&runs).
		Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to get the pending runs")
	}

	return runs, nil
}

func (r *Repository) GetRun(ctx context.Context, id int) (*Run, error) {
	var run Run
	err := r.db.
		WithContext(ctx).
		Find(&run).
		Where("id = ?", id).
		Error

	if err != nil {
		return nil, errors.Wrap(err, "failed to get run")
	}

	return &run, nil
}

func (r *Repository) GetRuns(ctx context.Context) ([]Run, error) {
	var runs []Run
	err := r.db.WithContext(ctx).
		Order("created_at desc").
		Find(&runs).
		Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the runs")
	}

	return runs, nil
}
