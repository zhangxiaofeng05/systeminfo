## usage
get system info
```shell
make dev
```
simple: http://localhost:8080/system  
complex: http://localhost:8080/system?all=true  

## docker
learn how to write Dockerfile
```shell
# Building the Image
docker build -t systeminfo .
# run image
docker run -p 8080:8080 systeminfo:latest
```

## dependabot

https://docs.github.com/cn/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file
