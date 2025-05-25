package manager_repository

import (
	"context"
	"geoindexing_comparison/backend/helpers/stats"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
)

type Stats struct {
	RunID int `gorm:"primaryKey"`
	Idx   int `gorm:"primaryKey"`

	Index  string
	Task   string
	Amount int

	Durs stats.Array[time.Duration] `gorm:"serializer:json"`
}

func (r *Repository) GetStats(ctx context.Context, runID int) ([]Stats, error) {
	var v []Stats

	err := r.db.WithContext(ctx).
		Where("run_id = ?", runID).
		Order("idx asc").
		Find(&v).
		Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stats")
	}

	return v, nil
}

func (r *Repository) GetLastStat(ctx context.Context, runID int) (int, error) {
	var v Stats

	err := r.db.WithContext(ctx).
		Where("run_id = ?", runID).
		Order("idx desc").
		Limit(1).
		First(&v).
		Error
	if err != nil {
		return 0, errors.Wrap(err, "failed to get stats")
	}

	return v.Idx, nil
}

func (r *Repository) SaveStats(ctx context.Context, v *Stats) error {
	err := r.db.
		WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "run_id"}, {Name: "idx"}},
		DoUpdates: clause.AssignmentColumns([]string{"index", "task", "amount", "durs"}),
	}).
		Save(v).
		Error
	if err != nil {
		return errors.Wrap(err, "failed to save stats")
	}

	return nil
}
