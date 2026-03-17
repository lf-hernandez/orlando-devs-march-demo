BINARY  := app
VERSION ?= dev
COMMIT  := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)"

.PHONY: build run clean setup

build:
	go build $(LDFLAGS) -o $(BINARY) .

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)

setup:
	git config core.hooksPath .githooks
	@echo "hooks configured"
