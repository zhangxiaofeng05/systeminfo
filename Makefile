BIN_DIR=./bin
VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse HEAD)
LDFLAGS := -w -s -X main.version=$(VERSION) -X main.commit=$(COMMIT)

help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## dev: run server
dev:
	go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/systeminfo -v .
	$(BIN_DIR)/systeminfo

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
