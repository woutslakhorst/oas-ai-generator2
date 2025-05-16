package fixtures

import "example.com/blobapi/internal/models"

func NewBlob() *models.Blob {
	return &models.Blob{
		Name:   "fixture blob",
		Photo:  "photo.png",
		Status: "alive",
	}
}
