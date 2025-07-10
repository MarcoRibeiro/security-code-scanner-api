# security-code-scanner-api

## Overview

This is the backend API for the Security Code Scanner project. It provides endpoints to scan source code for security vulnerabilities using pluggable analyzers. The API supports scanning local directories or remote Git repositories, and can be easily extended with new analyzers.

## Features

- Scan local folders or remote Git repositories for security issues
- Pluggable analyzer architecture
- Exclude files or directories from scans
- RESTful API (JSON)
- OpenAPI/Swagger documentation

## Getting Started

### Prerequisites

- Go 1.20+
- Docker (optional, for containerized deployment)

### Running Locally

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/security-code-scanner-api.git
   cd security-code-scanner-api/backend
   ```
2. Install dependencies:
   ```sh
   go mod download
   ```
3. Run the API:
   ```sh
   go run ./cmd/codeScanner
   ```

### Running with Docker

1. Build the Docker image:
   ```sh
   docker build -t security-code-scanner-api ..
   ```
2. Run the container:
   ```sh
   docker run -p 8080:8080 security-code-scanner-api
   ```

## API Usage

### Start a Scan

- **Endpoint:** `POST /v1/scans`
- **Request Body:**
  ```json
  {
    "path": "path/to/scan/or/git/repo",
    "configuration": {
      "exclude": ["node_modules", "testdata"]
    }
  }
  ```
- **Response:**
  ```json
  {
    "scan": {
      "path": "C:\\Users\\marco\\AppData\\Local\\Temp\\repo-2150606555",
      "findings": [
        {
          "rule": "SQL Injection Analyzer",
          "file": "main.go",
          "message": "SELECT * FROM users WHERE ...",
          "line": 42
        }
      ],
      "done": true
    }
  }
  ```

### API Documentation

- After running the API, visit `/swagger/index.html` for interactive docs (if enabled).

## Project Structure

- `cmd/codeScanner/` - Main entry point
- `internal/domain/` - Core domain models and interfaces
- `internal/analyzers/` - Analyzer implementations
- `internal/useCases/` - Scan orchestration logic
- `internal/api/v1/scans/` - HTTP handlers and DTOs
- `internal/utils/` - Utility functions

## Testing

Run all tests:

```sh
go test ./...
```
