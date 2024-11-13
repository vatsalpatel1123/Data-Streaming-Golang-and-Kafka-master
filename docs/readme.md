# Real-Time Data Streaming API with Golang and Redpanda (Kafka)

## Overview
This project is a high-performance API built in Golang, designed to stream data in real-time between clients and Redpanda (a Kafka-compatible streaming platform). It leverages WebSocket and Server-Sent Events (SSE) to achieve bi-directional communication, providing clients with immediate feedback on processed data. The API is optimized for handling at least 1000 concurrent streams with low latency, ensuring efficient, scalable data handling in real-time applications.

## Features
- **RESTful API Endpoints** for initiating, sending, and retrieving real-time stream data.
- **Kafka (Redpanda) Integration** for distributed message handling and real-time processing.
- **Concurrency and Performance Optimization** using goroutines, channels, and efficient Kafka consumers/producers.
- **Real-Time Processing** with WebSocket/SSE-based feedback for continuous data retrieval.
- **Robust Error Handling and Logging** for tracking requests, Kafka interactions, and errors.
- **Secure API Design** with API key authentication, stream isolation, and rate limiting.

---

## Table of Contents
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Endpoints](#endpoints)
- [Architecture](#architecture)
- [Performance Benchmarking](#performance-benchmarking)
- [Configuration](#configuration)
- [Future Improvements](#future-improvements)
- [Contact](#contact)

---

## Getting Started

### Prerequisites
- **Go** (>=1.16)
- **Redpanda** or **Kafka** for message streaming
- **Docker** (optional, for easy Redpanda setup)
- **Prometheus** (optional, for monitoring)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your_username/real-time-streaming-api.git
   cd real-time-streaming-api

2. Install dependencies:
   ```bash
    go mod tidy


3. Start Redpanda/Kafka (Docker example):
   ```bash
   docker run -d --name redpanda -p 9092:9092 vectorized/redpanda:latest redpanda start --overprovisioned --smp 1 --memory 1G --reserve-memory 0M --node-id 0 --check=false --kafka-addr PLAINTEXT://0.0.0.0:9092


4. Run the application:
   ```bash
   go run main.go

5. Optionally, run Prometheus metrics server:
   ```bash
   go run monitoring/prometheus.go

6. Run the application with final benchmark:
   ```bash
   docker ps

## Environment Variables
Configure environment variables in a `.env` file (sample below):

```env
KAFKA_BROKER=localhost:9092
API_KEY=test-api-key
PROMETHEUS_PORT=2112
```

## Project Structure
```text

├── handlers
│   └── stream-handler.go          # Stream request handlers
├── monitoring
│   └── prometheus.go              # Prometheus metrics setup
├── services
│   ├── kafka.go                   # Kafka producer and consumer
│   └── processing.go              # Data processing functions
├── utils
│   ├── auth.go                    # API key authentication middleware
│   ├── logger.go                  # Logging utilities
│   └── rate_limiter.go            # Rate limiting utilities
├── Dockerfile                     # Docker setup for containerization
├── main.go                        # Application entry point
├── config.go                      # Configuration and environment variable setup
├── router.go                      # API routing setup
├── benchmark.go                   # Benchmarking code
└── README.md                      # Project documentation

```

## Endpoints
- **POST** `/stream/start`: Initiate a new data stream.
- **POST** `/stream/{stream_id}/send`: Send data to the server for the specified stream.
- **GET** `/stream/{stream_id}/results`: Retrieve processed results in real-time for the specified stream.

**Example Requests:**

```bash
curl -X POST http://localhost:8080/stream/start -H "x-api-key: test-api-key"
curl -X POST http://localhost:8080/stream/{stream_id}/send -d '{"data": "sample data"}' -H "x-api-key: test-api-key"
curl -X GET http://localhost:8080/stream/{stream_id}/results -H "x-api-key: test-api-key"
```

## Architecture

### Core Components
- **Router** (`router.go`): Manages API endpoints and routes requests to appropriate handlers.
- **Handlers** (`handlers/stream-handler.go`): Handles streaming requests, data processing, and WebSocket/SSE connections.
- **Kafka Integration** (`services/kafka.go`): Manages Kafka producer and consumer operations.
- **Processing** (`services/processing.go`): Processes incoming data chunks in real-time and performs transformations.
- **Monitoring** (`monitoring/prometheus.go`): Collects metrics on requests and errors for Prometheus.
- **Utilities** (`utils/auth.go`, `utils/logger.go`, `utils/rate_limiter.go`): Provides authentication, logging, and rate-limiting utilities.

### Data Flow
1. Clients initiate streams via **POST /stream/start**.
2. Data is sent to Kafka through **POST /stream/{stream_id}/send**, where each `stream_id` represents a unique Kafka topic or partition.
3. **GET /stream/{stream_id}/results** returns processed data to the client via WebSocket/SSE for real-time feedback.

### Concurrency Management
- **Goroutines and Channels**: Employed for concurrent data handling, enabling high throughput and low latency.
- **Kafka Consumers and Producers**: Used to handle individual data streams, isolating them by `stream_id`.

---

## Performance Benchmarking

### Results Summary
- **Concurrent Streams**: 1000
- **Requests per Second**: 5000
- **Average Latency**: ~12ms per request

### Benchmarking Script
To run the benchmark test:
```bash
go run benchmark.go
```

## Configuration

### Prometheus Monitoring
Prometheus metrics are available at `http://localhost:2112/metrics`. Use Prometheus to monitor request counts, error counts, and system resource usage.

### Security
- **API Key Authentication**: Enforced with middleware (`utils/auth.go`).
- **Rate Limiting**: Rate-limits each client to prevent overloading the system (`utils/rate_limiter.go`).

## Future Improvements
- **Horizontal Scaling**: Expand with additional Kafka nodes for higher throughput.
- **Advanced Security**: Add features such as OAuth2 or JWT authentication.
- **Complex Data Processing**: Integrate real-time machine learning predictions.
- **Enhanced Error Handling**: Capture and log finer error details for improved troubleshooting.

## Contact
- For inquiries, please contact me 
- Author: **Ayushi Modi**
- Email: **ayushimodi818@gmail.com**.

