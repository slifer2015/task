GO               = go
GOBIN			 ?= $(PWD)/bin
TARGET_DIR       ?= $(PWD)/.build
GOOS 			 ?= linux
GOARCH			 ?= amd64
M                = $(shell printf "\033[34;1m>>\033[0m")

.PHONY: all
all: build

.PHONY: build-server
build: ## Build  binary
	$(info $(M) building server...)
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(TARGET_DIR)/server ./cmd/*.go

.PHONY: start
start:
	$(info $(M) starting application...)
	$(GO) run ./cmd/*.go

.PHONY: install-tools
install-tools: $(GOBIN) ## Install tools needed for development
	@GOBIN=$(GOBIN) $(GO) install github.com/swaggo/swag/cmd/swag@latest
	@GOBIN=$(GOBIN) $(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: lint
lint: install-tools## Run linters
	$(info $(M) running linters...)
	$(GOBIN)/golangci-lint run --config $(PWD)/.golangci.yml --timeout 5m0s $(PWD)/...

.PHONY: test
test: ## Run all tests
	$(info $(M) running tests...)
	@$(GO) test ./... -v -cover

.PHONY: generate
generate: install-tools## Run go generate
	$(info $(M) generating...)
	@$(GO) generate ./...

.PHONY: clean
clean: ## Clean build artifacts
	rm -rf $(TARGET_DIR)