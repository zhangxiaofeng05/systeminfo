BIN_DIR=./bin

help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## dev: run server
dev:
	go build -o $(BIN_DIR)/systeminfo -v .
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
	go build -o $(BIN_DIR)/systeminfo -v .

## clean: clean bin directory
clean:
	rm $(BIN_DIR)/*
