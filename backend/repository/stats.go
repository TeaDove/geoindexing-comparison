package repository

import (
	"context"
	"geoindexing_comparison/backend/service/stats"
	"github.com/pkg/errors"
	"time"
)

type Stats struct {
	RunID     uint64
	CreatedAt time.Time

	Index  string `gorm:"primaryKey"`
	Task   string `gorm:"primaryKey"`
	Amount uint64 `gorm:"primaryKey"`

	Durs stats.Array[time.Duration] `gorm:"serializer:json"`
}

func (r *Repository) SaveStats(ctx context.Context, v *Stats) error {
	err := r.db.WithContext(ctx).Save(v).Error
	if err != nil {
		return errors.Wrap(err, "failed to save stats")
	}

	return nil
}
