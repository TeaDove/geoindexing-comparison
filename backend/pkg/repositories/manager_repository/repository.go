package manager_repository

import (
	"context"
	"geoindexing_comparison/pkg/helpers"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(ctx context.Context) (*Repository, error) {
	db, err := gorm.Open(sqlite.Open(helpers.Settings.SqlitePath), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to open gorm.db")
	}

	err = db.WithContext(ctx).Migrator().AutoMigrate(&Stats{}, &Run{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to auto migrate")
	}

	return &Repository{db: db}, nil
}
