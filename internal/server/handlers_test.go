package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"example.com/blobapi/internal/models"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	if err := db.AutoMigrate(&models.Blob{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestAddBlob(t *testing.T) {
	db := setupTestDB(t)
	h := &Handler{DB: db}
	r := gin.Default()
	h.RegisterRoutes(r)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/blobs", strings.NewReader(`{"name":"a","photo":"b"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 got %d", w.Code)
	}
}
