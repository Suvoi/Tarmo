# Tarmo - Engine

The core backend component for the **Tarmo** system, designed to optimize and control processes based on templates and resources.

## Overview

Tarmo Engine provides a robust REST API for managing:
- **Resources**: Individual components with tracking for quantity, unit, and pricing.
- **Templates**: Structured processes defining sequential steps and required resources.

## Disclaimer

This is a work in progress. The project is not yet ready for production use.

## Entrypoints

- **REST API**: `cmd/rest_api/main.go`

## Tech Stack

- **Language**: Go
- **Routing**: [chi](https://github.com/go-chi/chi)
- **Database Adapters**: SQLite3 for now, but designed to be pluggable.
- **API Specification**: OpenAPI 3.0 (`oas.yml`)

## Getting Started

### Prerequisites

- [Go](https://go.dev/) (version 1.25.5 or higher as specified in `go.mod`)

### Running the Server

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the application:
   ```bash
   go run cmd/rest_api/main.go
   ```

The server will start on the default port (configured via environment variables or defaults).

## API Documentation

The API is documented using OpenAPI. You can find the full specification in `oas.yml`.