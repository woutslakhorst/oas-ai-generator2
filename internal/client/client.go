package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/blobapi/internal/models"
)

// Client provides methods to call the blob API server.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// New returns a new Client with the given base URL.
func New(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// GetBlobs fetches all blobs from the server.
func (c *Client) GetBlobs() ([]models.Blob, error) {
	url := fmt.Sprintf("%s/blobs", c.BaseURL)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var blobs []models.Blob
	if err := json.NewDecoder(resp.Body).Decode(&blobs); err != nil {
		return nil, err
	}
	return blobs, nil
}
