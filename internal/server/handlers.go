package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"example.com/blobapi/internal/models"
)

// Handler wraps dependencies for HTTP handlers.
type Handler struct {
	DB *gorm.DB
}

// RegisterRoutes registers all API routes.
func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.PUT("/blobs", h.updateBlob)
	r.POST("/blobs", h.addBlob)
	r.GET("/blobs", h.getBlobs)
	r.GET("/blobs:id", h.findBlobByID)
	r.DELETE("/blobs:id", h.deleteBlob)
}

func (h *Handler) updateBlob(c *gin.Context) {
	var blob models.Blob
	if err := c.ShouldBindJSON(&blob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Save(&blob).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blob)
}

func (h *Handler) addBlob(c *gin.Context) {
	var blob models.Blob
	if err := c.ShouldBindJSON(&blob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&blob).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blob)
}

func (h *Handler) getBlobs(c *gin.Context) {
	var blobs []models.Blob
	if err := h.DB.Find(&blobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blobs)
}

func (h *Handler) findBlobByID(c *gin.Context) {
	id := c.Param("id")
	var blob models.Blob
	if err := h.DB.First(&blob, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blob)
}

func (h *Handler) deleteBlob(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Blob{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
