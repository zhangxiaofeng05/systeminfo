## system info
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)
![golangci-lint](https://github.com/zhangxiaofeng05/systeminfo/actions/workflows/golangci-lint.yml/badge.svg?branch=main)
![license](https://img.shields.io/github/license/zhangxiaofeng05/systeminfo)
<!-- [![codecov](https://codecov.io/gh/zhangxiaofeng05/systeminfo/branch/main/graph/badge.svg?token=OAQ31EUR2N)](https://codecov.io/gh/zhangxiaofeng05/systeminfo) -->

## release
![GitHub release (latest by date)](https://img.shields.io/github/v/release/zhangxiaofeng05/systeminfo)
```
shasum -a 256 xxx.tar.gz
```

## go install
```bash
# install
go install github.com/zhangxiaofeng05/systeminfo@latest
# run
systeminfo
```
call: http://127.0.0.1:8080

## container
### docker
#### Docker Hub
![Docker Image Version](https://img.shields.io/docker/v/zhangxiaofeng05/systeminfo)
```
docker run -d -p 8080:8080 --name systeminfo zhangxiaofeng05/systeminfo:latest
```
#### GitHub Container Registry
```
docker run -d -p 8080:8080 --name ghcr.io/systeminfo zhangxiaofeng05/systeminfo:latest
```

Dockerfile reference: https://docs.docker.com/engine/reference/builder/  
reference: https://studygolang.com/articles/24854

### docker-compose
```
wget https://raw.githubusercontent.com/zhangxiaofeng05/systeminfo/main/docker-compose.yml

docker-compose up -d
```
### k8s
```
kubectl apply -f https://raw.githubusercontent.com/zhangxiaofeng05/systeminfo/main/k8s/deployment.yaml
```
`attention`: server port is 30000

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

