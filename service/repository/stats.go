package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
)

type Stats struct {
	RunID uuid.UUID

	Index  string `gorm:"primaryKey"`
	Task   string `gorm:"primaryKey"`
	Amount uint64 `gorm:"primaryKey"`

	Durs datatypes.JSON
}

func (r *Repository) SaveStats(ctx context.Context, v *Stats) error {
	err := r.db.WithContext(ctx).Save(v).Error
	if err != nil {
		return errors.Wrap(err, "failed to save stats")
	}

	return nil
}
