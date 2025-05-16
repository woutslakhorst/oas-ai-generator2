package migrations

import (
	"example.com/blobapi/internal/models"
	"gorm.io/gorm"
)

// Migrate runs database migrations.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Blob{})
}
