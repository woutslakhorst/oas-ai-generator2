# Blob API and Client

This project provides a small HTTP API server and a command line client.
The API is defined in `blobs.yaml` and implemented using Go and Gin.

## Prerequisites

- Go 1.20 or later installed

## Running the server

From the repository root run:

```bash
go run ./cmd/blob server
```

Running without a subcommand also starts the server:

```bash
go run ./cmd/blob
```

The server listens on `localhost:8080` and uses a SQLite database file named
`blob.db`. A new database file will be created automatically if it does not
exist.

## Running the client

The client fetches blob data from the running server and prints a JSON object
containing the first and last blob (sorted by name). To run the client, use:

```bash
go run ./cmd/blob first-last
```

To show all blobs with scrambled names, run:

```bash
go run ./cmd/blob scrambled
```

Both commands assume the server is running on `http://localhost:8080`. You can
specify a different base URL as the first argument:

```bash
go run ./cmd/blob first-last http://other-host:8080
```

## Testing

Unit tests are located under `internal/` and `cmd/`. Run all tests with:

```bash
go test ./...
```

