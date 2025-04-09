package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(ctx context.Context) (*Repository, error) {
	db, err := gorm.Open(sqlite.Open("./data/db.sqlite"), &gorm.Config{}) // TODO move to settings
	if err != nil {
		return nil, errors.Wrap(err, "failed to open gorm.db")
	}

	err = db.WithContext(ctx).Migrator().AutoMigrate(&Stats{}, &Run{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to auto migrate")
	}

	return &Repository{db: db}, nil
}
