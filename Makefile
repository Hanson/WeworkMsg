.PHONY: build clean tool lint help build-cli build-cli-all

prod:
	GOOS=linux go build .

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -s -w -X github.com/hanson/WeworkMsg/internal/cli.version=$(VERSION)

build-cli:
	go build -ldflags "$(LDFLAGS)" -o wework-cli ./cmd/wework-cli/

build-cli-all:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/wework-cli-linux-amd64 ./cmd/wework-cli/
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/wework-cli-darwin-amd64 ./cmd/wework-cli/
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o dist/wework-cli-darwin-arm64 ./cmd/wework-cli/
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/wework-cli-windows-amd64.exe ./cmd/wework-cli/
