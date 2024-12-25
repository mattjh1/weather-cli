# Detect platform and architecture
PLATFORM := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

# Map architecture names for Go builds
ifeq ($(ARCH), x86_64)
    ARCH := amd64
endif
ifeq ($(ARCH), arm64)
    ARCH := arm64
endif

APP_NAME := weather
BUILD_DIR := ./bin
INSTALL_DIR := $(HOME)/.local/bin

.PHONY: clean install uninstall help

build: ## Build the application binary for the current platform
	mkdir -p $(BUILD_DIR)
	GOOS=$(PLATFORM) GOARCH=$(ARCH) go build -o $(BUILD_DIR)/$(APP_NAME) weather.go

clean: ## Remove the build artifacts
	rm -rf $(BUILD_DIR)

install: build ## Install the binary to $(INSTALL_DIR)
	install $(BUILD_DIR)/$(APP_NAME) $(INSTALL_DIR)/$(APP_NAME)

uninstall: ## Remove the installed binary from $(INSTALL_DIR)
	rm -f $(INSTALL_DIR)/$(APP_NAME)

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'
