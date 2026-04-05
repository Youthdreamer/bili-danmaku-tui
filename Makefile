APP := bili-danmaku-tui

MODULE_NAME := $(shell go list -m)
VERSION     := $(shell git describe --tags --always 2>/dev/null || echo "dev")
COMMIT      := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE    		:= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

# 允许外部传入平台参数，默认使用当前系统环境
GOOS   ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

LDFLAGS := -ldflags "\
-X '$(MODULE_NAME)/cmd.Version=$(VERSION)' \
-X '$(MODULE_NAME)/cmd.GitCommit=$(COMMIT)' \
-X '$(MODULE_NAME)/cmd.BuildTime=$(DATE)'"

.PHONY: all build run clean

all: build

build:
				@echo "Building $(APP) $(VERSION) for $(GOOS)/$(GOARCH)..."
				@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $(APP)
				@echo "Build complete: ./$(APP)"

# 专门给 GitHub Action 使用的发布命令
release:
				@mkdir -p dist
				@echo "Releasing $(APP) for $(GOOS)/$(GOARCH)..."
				GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o dist/$(APP)
				cd dist && tar -czf $(APP)-linux-$(GOARCH).tar.gz $(APP)
				@rm dist/$(APP)

run: build
				@echo "Running $(APP)..."
				@./$(APP)

clean:
				@echo "Cleaning..."
				@rm -f $(APP)
				@rm -rf dist
				@echo "Done."
