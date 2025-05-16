package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/spf13/cobra"

	"example.com/blobapi/internal/client"
	"example.com/blobapi/internal/models"
	"example.com/blobapi/internal/server"
)

var rootCmd = &cobra.Command{
	Use:   "blobapi",
	Short: "Blob API server and client",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runServer()
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runServer()
	},
}

var firstLastCmd = &cobra.Command{
	Use:   "first-last [baseURL]",
	Short: "Print the first and last blob",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL := "http://localhost:8080"
		if len(args) > 0 {
			baseURL = args[0]
		}
		return runClient("first-last", baseURL)
	},
}

var scrambledCmd = &cobra.Command{
	Use:   "scrambled [baseURL]",
	Short: "Print blobs with scrambled names",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL := "http://localhost:8080"
		if len(args) > 0 {
			baseURL = args[0]
		}
		return runClient("scrambled", baseURL)
	},
}

func runServer() error {
	server.Run(":8080", "blob.db")
	return nil
}

func runClient(op, baseURL string) error {
	c := client.New(baseURL)
	blobs, err := c.GetBlobs()
	if err != nil {
		return fmt.Errorf("failed to get blobs: %w", err)
	}
	if len(blobs) == 0 {
		fmt.Println("[]")
		return nil
	}

	if op == "scrambled" {
		rand.Seed(time.Now().UnixNano())
		scrambleNames(blobs)
		enc := json.NewEncoder(os.Stdout)
		if err := enc.Encode(blobs); err != nil {
			return fmt.Errorf("failed to encode result: %w", err)
		}
		return nil
	}

	sort.Slice(blobs, func(i, j int) bool { return blobs[i].Name < blobs[j].Name })
	result := struct {
		First interface{} `json:"first"`
		Last  interface{} `json:"last"`
	}{
		First: blobs[0],
		Last:  blobs[len(blobs)-1],
	}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(result); err != nil {
		return fmt.Errorf("failed to encode result: %w", err)
	}
	return nil
}

func scrambleNames(blobs []models.Blob) {
	for i := range blobs {
		r := []rune(blobs[i].Name)
		rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })
		blobs[i].Name = string(r)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd, firstLastCmd, scrambledCmd)
}
