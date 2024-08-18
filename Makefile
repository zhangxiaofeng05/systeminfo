BIN_DIR=./bin
VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
TAG_DATE := $(shell git show -s --format=%ci) # date中包含空格，引用时需要加单引号
LDFLAGS := -w -s -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.branch=$(BRANCH) -X 'main.tagDate=$(TAG_DATE)'

help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## run: run server
run:
	go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/systeminfo -v .
	$(BIN_DIR)/systeminfo

## run_live_reloading: run server with live reloading
run_live_reloading:
	# https://github.com/air-verse/air
	# 嵌套了", 转义引用
	air --build.cmd "go build -ldflags \"$(LDFLAGS)\" -o $(BIN_DIR)/systeminfo -v ." --build.bin "$(BIN_DIR)/systeminfo"

## lint: run golangci-lint
lint:
	golangci-lint run ./...

## test: run test. view result:$ go tool cover -html=coverage.txt
test:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

## mod_tidy: go mod tidy
mod_tidy:
	go mod tidy

## build: build executable
build:
	go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/systeminfo -v .

## clean: clean bin directory
clean:
	rm $(BIN_DIR)/*

## docker_build: local build image
docker_build:
	docker build -t systeminfo .
	# DOCKER_BUILDKIT=0 docker build -t systeminfo .

## docker_run: run for local build image
docker_run:
	
	docker run -d -p 8080:8080 --name systeminfo systeminfo:latest

## docker_stop: stop for local build image
docker_stop:
	docker stop systeminfo

## docker_rm: remove for local build image
docker_rm:
	docker rm systeminfo
	docker rmi systeminfo
