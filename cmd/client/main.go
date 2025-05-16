package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"example.com/blobapi/internal/client"
)

func main() {
	baseURL := "http://localhost:8080"
	if len(os.Args) > 1 {
		baseURL = os.Args[1]
	}
	c := client.New(baseURL)
	blobs, err := c.GetBlobs()
	if err != nil {
		log.Fatalf("failed to get blobs: %v", err)
	}
	if len(blobs) == 0 {
		fmt.Println("[]")
		return
	}
	sort.Slice(blobs, func(i, j int) bool {
		return blobs[i].Name < blobs[j].Name
	})
	result := struct {
		First interface{} `json:"first"`
		Last  interface{} `json:"last"`
	}{
		First: blobs[0],
		Last:  blobs[len(blobs)-1],
	}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(result); err != nil {
		log.Fatalf("failed to encode result: %v", err)
	}
}
