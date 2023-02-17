## system info
![ci](https://github.com/zhangxiaofeng05/systeminfo/actions/workflows/ci.yml/badge.svg?branch=main)
![license](https://img.shields.io/github/license/zhangxiaofeng05/systeminfo)
<!-- [![codecov](https://codecov.io/gh/zhangxiaofeng05/systeminfo/branch/main/graph/badge.svg?token=OAQ31EUR2N)](https://codecov.io/gh/zhangxiaofeng05/systeminfo) -->

## go install
```bash
# install
go install github.com/zhangxiaofeng05/systeminfo@latest
# run
systeminfo
```

simple: http://localhost:8080/system  
complex: http://localhost:8080/system?all=true  

## container
### docker
```
docker run -d -p 8080:8080 --name systeminfo zhangxiaofeng05/systeminfo:latest
```

Dockerfile reference: https://docs.docker.com/engine/reference/builder/  
reference: https://studygolang.com/articles/24854

### docker-compose
```
wget https://raw.githubusercontent.com/zhangxiaofeng05/systeminfo/main/docker-compose.yml

docker-compose up -d
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

