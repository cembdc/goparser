BIN_DIR := bin
BIN_NAME := app

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

.PHONY: all clean build run test crossbuild hello

hello:
	echo "Hello"

all: clean build

build: clean $(BIN_DIR)
	@go build -ldflags="-w -s" -o $(BIN_DIR)/$(BIN_NAME)

run: clean build
	@$(BIN_DIR)/$(BIN_NAME)

test:
	@go test -v ./...

clean:
	@rm -rf $(BIN_DIR)

crossbuild: clean $(BIN_DIR)
	@GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(BIN_DIR)/$(BIN_NAME)_linux_amd64
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o $(BIN_DIR)/$(BIN_NAME)_darwin_amd64
	@GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o $(BIN_DIR)/$(BIN_NAME)_windows_amd64.exe

$(BIN_DIR):
