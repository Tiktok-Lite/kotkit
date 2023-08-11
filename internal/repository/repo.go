package repository

import (
	z "github.com/Tiktok-Lite/kotkit/pkg/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Repository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:     db,
		logger: z.Logger(),
	}
}
