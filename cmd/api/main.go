package main

import "example.com/blobapi/internal/server"

func main() {
	server.Run(":8080", "blob.db")
}
