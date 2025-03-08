[![Go](https://github.com/osesantos/hermes/actions/workflows/go.yml/badge.svg)](https://github.com/osesantos/hermes/actions/workflows/go.yml)

# Hermes 📬
> Portfolio project

A lightweight, high-performance message queue system built in Go.

## Features (Planned)

- In-memory message routing
- Persistent storage
- gRPC/Network interface
- Concurrent processing
- Simple client library

## Tech Stack

- Go
- gRPC
- Badger/Bolt (Storage)
- Docker
- Prometheus Metrics

## Development Stages

1. Basic In-Memory Queue
2. Persistent Storage
3. Network Server
4. Client Libraries

---

## Development

### Build

```bash
go build -o hermes ./src/main.go
```

### Run

```bash
./hermes
```

### Test

```bash
go test -v -cover ./...
```

### Docker
> TODO: Add Docker

```bash
docker build -t hermes .
docker run -p 8080:8080 hermes
```

### Docker Compose
> TODO: Add Docker Compose

```bash
docker-compose up
```

