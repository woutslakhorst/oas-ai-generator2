package models

import "gorm.io/gorm"

// Blob represents the blob model as defined in the OpenAPI spec.
type Blob struct {
	gorm.Model
	Name   string `json:"name"`
	Photo  string `json:"photo"`
	Status string `json:"status"`
}
