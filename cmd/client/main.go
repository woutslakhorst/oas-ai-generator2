package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"

	"example.com/blobapi/internal/client"
	"example.com/blobapi/internal/models"
)

func main() {
	op := "first-last"
	baseURL := "http://localhost:8080"
	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "first-last", "scrambled":
			op = args[0]
			if len(args) > 1 {
				baseURL = args[1]
			}
		default:
			baseURL = args[0]
		}
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

	if op == "scrambled" {
		rand.Seed(time.Now().UnixNano())
		scrambleNames(blobs)
		enc := json.NewEncoder(os.Stdout)
		if err := enc.Encode(blobs); err != nil {
			log.Fatalf("failed to encode result: %v", err)
		}
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

func scrambleNames(blobs []models.Blob) {
	for i := range blobs {
		r := []rune(blobs[i].Name)
		rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })
		blobs[i].Name = string(r)
	}
}
