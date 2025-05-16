package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/blobapi/internal/models"
)

func TestGetBlobs(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/blobs", func(w http.ResponseWriter, r *http.Request) {
		blobs := []models.Blob{{Name: "b"}, {Name: "a"}}
		json.NewEncoder(w).Encode(blobs)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	c := New(srv.URL)
	blobs, err := c.GetBlobs()
	if err != nil {
		t.Fatalf("GetBlobs failed: %v", err)
	}
	if len(blobs) != 2 {
		t.Fatalf("expected 2 blobs got %d", len(blobs))
	}
}
