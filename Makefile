BINARY_NAME=togo
OUTPUT_DIR=bin

GO=go
GOFLAGS=-ldflags="-s -w"

build: build-windows build-linux build-macos

build-windows: build-windows-amd64 build-windows-arm64
build-linux: build-linux-amd64 build-linux-arm64
build-macos: build-macos-amd64 build-macos-arm64

build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-windows-amd64.exe

build-windows-arm64:
	GOOS=windows GOARCH=arm64 $(GO) build $(GOFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-windows-arm64.exe

build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64

build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GO) build $(GOFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux-arm64

build-macos-amd64:
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-macos-amd64

build-macos-arm64:
	GOOS=darwin GOARCH=arm64 $(GO) build $(GOFLAGS) -o $(OUTPUT_DIR)/$(BINARY_NAME)-macos-arm64

# Instalar binário em ambientes Linux
install: build-linux-amd64
	install -m 0755 $(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64 /usr/local/bin/$(BINARY_NAME)

.PHONY: build build-windows build-linux build-macos install