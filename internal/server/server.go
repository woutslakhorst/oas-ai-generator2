package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"example.com/blobapi/internal/db/migrations"
)

// New creates and configures a Gin engine with database connection and
// migrations applied.
func New(dsn string) (*gin.Engine, *gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	if err := migrations.Migrate(db); err != nil {
		return nil, nil, err
	}

	r := gin.Default()
	h := &Handler{DB: db}
	h.RegisterRoutes(r)
	return r, db, nil
}

// Run starts the HTTP server on the given address.
func Run(addr, dsn string) {
	r, _, err := New(dsn)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	if err := r.Run(addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
