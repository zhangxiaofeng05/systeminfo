## system info
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)
![golangci-lint](https://github.com/zhangxiaofeng05/systeminfo/actions/workflows/golangci-lint.yml/badge.svg?branch=main)
![license](https://img.shields.io/github/license/zhangxiaofeng05/systeminfo)
<!-- [![codecov](https://codecov.io/gh/zhangxiaofeng05/systeminfo/branch/main/graph/badge.svg?token=OAQ31EUR2N)](https://codecov.io/gh/zhangxiaofeng05/systeminfo) -->

## Installation

### Binary (recommended)

Download the pre-built binary for your platform from the [GitHub Releases](https://github.com/zhangxiaofeng05/systeminfo/releases) page.

```bash
# Verify checksum after download
shasum -a 256 systeminfo_*.tar.gz
# Compare with checksums.txt from the same release
```

Supported platforms: Linux, macOS, Windows — amd64, arm64, arm, 386, ppc64.

### go install

Requires Go 1.24+.

```bash
go install github.com/zhangxiaofeng05/systeminfo@latest
```

### Docker

#### Docker Hub
![Docker Image Version](https://img.shields.io/docker/v/zhangxiaofeng05/systeminfo)
```bash
docker run -d -p 8080:8080 --name systeminfo zhangxiaofeng05/systeminfo:latest
```

#### GitHub Container Registry
```bash
docker run -d -p 8080:8080 --name systeminfo ghcr.io/zhangxiaofeng05/systeminfo:latest
```

### docker-compose

```bash
wget https://raw.githubusercontent.com/zhangxiaofeng05/systeminfo/main/docker-compose.yml
docker-compose up -d
```

### Kubernetes

```bash
kubectl apply -f https://raw.githubusercontent.com/zhangxiaofeng05/systeminfo/main/k8s/deployment.yaml
```

> Note: the NodePort is 30000.

## Usage

```bash
# run (default port 8080)
systeminfo

# run with custom port
systeminfo -port 9090

# run with pprof enabled
systeminfo -pprof
```

Access: http://127.0.0.1:8080

## API endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /` | List all available endpoints |
| `GET /ping` | Health check, returns `"pong"` |
| `GET /version` | Show current version |
| `GET /system` | Basic system info (CPU, memory, host, Go runtime) |
| `GET /system?all=true` | Full system info |


## dependabot

https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file

## badge
1. workflow  
https://docs.github.com/en/actions/monitoring-and-troubleshooting-workflows/adding-a-workflow-status-badge
2. shields  
https://shields.io/
3. Codecov  
https://docs.codecov.com/docs/quick-start  

## development
1. vscode  
2. https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-ssh  
3. https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers  

reference: https://github.com/devcontainers/images

