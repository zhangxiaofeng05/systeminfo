# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

- **Build binary**: `make build` (produces `./bin/systeminfo`)
- **Run server**: `make run` (uses default port 8080, optional flags `-port <num>` and `-pprof`)
- **Live‑reloading**: `make run_live_reloading` (requires `air`)
- **Run tests**: `make test` (includes race detector, generates `coverage.txt`)
- **Lint**: `make lint` (via `golangci-lint`)
- **Dependency tidy**: `make mod_tidy`
- **Docker image**: `make docker_build` then `make docker_run`
- **Docker cleanup**: `make docker_stop && make docker_rm`
- **Kubernetes deployment**: `kubectl apply -f https://github.com/zhangxiaofeng05/systeminfo/main/k8s/deployment.yaml`
- **Install binary via Go**: `go install github.com/zhangxiaofeng05/systeminfo@latest`
- **Check version**: `systeminfo -version` (or run `make run` and hit `/version` endpoint)

## High‑Level Architecture

- **Entry point**: `main` in `systeminfo.go` sets up a Gin HTTP server.
- **Command‑line flags**: `-port` (default 8080) and `-pprof` (adds pprof middleware).
- **Endpoints**:
  - `GET /ping` – health check, returns `"pong"`.
  - `GET /version` – JSON with build metadata (version, commit, etc.).
  - `GET /system` – returns a JSON map containing Go runtime info, host info, CPU info, and memory stats. Supports `?all=true` for full structs.
  - `GET /` – provides a map of all available URLs for quick discovery.
  - `GET /metrics` – Prometheus metrics (`http_requests_total`, `http_request_duration_seconds`).
  - `GET /debug/pprof/*` – pprof data when the `-pprof` flag is enabled.
- **Prometheus integration**: defined in `prometheus.go` – registers a request counter and a histogram, wrapped by a Gin middleware.
- **System info gathering**: uses `gopsutil` packages for host, CPU, and memory details.
- **Graceful shutdown**: listens for `SIGINT`/`SIGTERM`, shuts down the HTTP server with a 5‑second timeout.

## CI / CD

- GitHub Actions workflows in `.github/workflows/` handle linting, testing, and publishing Docker images to Docker Hub and GitHub Container Registry.
- Release workflow tags (`v*.*.*`) trigger Docker builds for multi‑arch images.

## Development Tips

- The server can be run locally with `make run`. Use `-pprof` for profiling while developing performance‑critical code.
- Run `make test -run TestXYZ` to execute a single test.
- Use `golangci-lint run` to catch static analysis issues before committing.
- The `Makefile` variable `LDFLAGS` injects version information from Git; ensure the repository is clean for reproducible builds.
