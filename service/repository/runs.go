package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"time"
)

type Run struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	CreatedBy string

	Status RunStatus

	Indexes datatypes.JSON
	Tasks   datatypes.JSON
	Start   uint64
	Stop    uint64
	Step    uint64
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
