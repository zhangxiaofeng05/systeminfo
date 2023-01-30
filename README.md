## system info
![ci](https://github.com/zhangxiaofeng05/systeminfo/actions/workflows/ci.yml/badge.svg?branch=main)
![license](https://img.shields.io/github/license/zhangxiaofeng05/systeminfo)
<!-- [![codecov](https://codecov.io/gh/zhangxiaofeng05/systeminfo/branch/main/graph/badge.svg?token=OAQ31EUR2N)](https://codecov.io/gh/zhangxiaofeng05/systeminfo) -->

## usage
```shell
make dev
```
simple: http://localhost:8080/system  
complex: http://localhost:8080/system?all=true  

## docker
local build
```shell
# Building the Image
docker build -t systeminfo .
# run image
docker run -p 8080:8080 systeminfo:latest
```

## dependabot

https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file

## badge
1. workflow  
https://docs.github.com/en/actions/monitoring-and-troubleshooting-workflows/adding-a-workflow-status-badge
2. shields  
https://shields.io/
3. Codecov  
https://docs.codecov.com/docs/quick-start  

