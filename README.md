# netemp

Lightweight event-driven network monitoring and processing service written in Go.

## Goals

* Lean and efficient
* Minimal dependencies
* Idiomatic Go architecture
* Extensible event pipeline
* Protocol-agnostic internal event model
* Easy to deploy as a single binary

---

# Current Features

* HTTP ingest endpoint
* Request normalization into internal `Event` objects
* Filter pipeline
* Method filtering
* Standard library HTTP server

---

# Project Structure

```text
netemp/
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── events/
│   ├── filters/
│   ├── pipeline/
│   └── server/
│
├── go.mod
└── README.md
```

---

# Architecture

```text
HTTP Request
    ↓
Normalize → Event
    ↓
Pipeline
    ↓
Filters
    ↓
Actions (planned)
```

The core idea is:

* transport-specific ingestion
* protocol-independent internal events
* composable processing pipeline

This allows future expansion into:

* TCP listeners
* UDP listeners
* raw packet processing
* async workers
* outbound enrichment/webhooks
* replay/testing pipelines

---

# Current Endpoints

## Health Check

```http
GET /healthz
```

Response:

```json
{"status":"ok"}
```

---

## Ingest Endpoint

```http
POST /ingest
```

Currently:

* normalizes request into an internal `Event`
* applies configured filters
* accepts/rejects event

---

# Running

## Start Server

```bash
go run ./cmd/app
```

Default listener:

```text
:8080
```

---

# Example Requests

## Health Check

```bash
curl localhost:8080/healthz
```

---

## Ingest Request

```bash
curl \
  -X POST \
  http://localhost:8080/ingest \
  -d 'hello world'
```

---

# Design Philosophy

This project intentionally favors:

* standard library usage
* explicit behavior
* composition over inheritance
* small focused interfaces
* minimal abstractions
---

# TODO

## Core Pipeline

* [ ] Action pipeline
* [ ] Logging actions
* [ ] Webhook actions
* [ ] Async worker pool
* [ ] Buffered event queues
* [ ] Context-aware processing

---

## Filters

* [ ] CIDR allow/deny filters
* [ ] Header filters
* [ ] Path filters
* [ ] Regex filters
* [ ] Rate limiting
* [ ] Port-state filtering

---

## Networking

* [ ] TCP listener support
* [ ] UDP listener support
* [ ] Raw socket experimentation
* [ ] Packet inspection research
* [ ] Optional pcap integration

---

## Observability

* [ ] Structured logging (`slog`)
* [ ] Metrics
* [ ] Request tracing
* [ ] Profiling endpoints
* [ ] Health metrics

---

## Reliability

* [ ] Graceful shutdown
* [ ] Request timeouts
* [ ] Backpressure handling
* [ ] Configurable worker limits
* [ ] Retry strategies

---

## Storage

* [ ] In-memory event retention
* [ ] SQLite backend
* [ ] Event replay tooling
* [ ] Export pipeline

---

## Security

* [ ] Request validation
* [ ] Authentication
* [ ] TLS support
* [ ] IP reputation hooks
* [ ] Abuse protections

---

# Long-Term Ideas

* Honeypot integrations
* Event replay engine
* Distributed workers
* Streaming pipelines
* eBPF experimentation
* Threat intel enrichment
* Real-time dashboards
* Rule engine
* Plugin system
