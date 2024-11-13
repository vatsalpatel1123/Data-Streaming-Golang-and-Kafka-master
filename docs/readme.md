# Real-Time Data Streaming API with Golang and Redpanda

## Overview
This API streams data to/from Redpanda (Kafka) and allows clients to send and receive data in real-time.

## Endpoints
- **POST /stream/start**: Start a new data stream.
- **POST /stream/{stream_id}/send**: Send data to an active stream.
- **GET /stream/{stream_id}/results**: Retrieve processed results.

## Setup and Usage
1. Start Redpanda: `docker-compose up -d`
2. Run the server: `go run main.go`
3. Monitor with Prometheus at `localhost:2112/metrics`.

## Performance Benchmark
Run `go run benchmark.go` to simulate 1000 concurrent streams.
