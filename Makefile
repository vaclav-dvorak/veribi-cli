GIT_REV?=$$(git rev-parse --short HEAD)
VERSION?=$$(git describe --tags --abbrev=0)
LDFLAGS="-s -w -X 'github.com/vaclav-dvorak/veribi-cli/cmd/veribi.version=$(VERSION)+$(GIT_REV)'"
goos?=$$(go env GOOS)
goarch?=$$(go env GOARCH)
file:=veribi
package:=$(file)_$(goos)_$(goarch)

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: default
default: help

## Build:
prepare: ## Download depencies and prepare dev env
	@pre-commit install
	@go mod download
	@go mod tidy
	@go mod vendor

build:  ## Builds the cli binary
	go build -trimpath -ldflags=$(LDFLAGS) -o ./bin/$(file) main.go

build-ci: ## Optimized build for CI
	@echo $(goos)/$(goarch)
	go build -trimpath -ldflags=$(LDFLAGS) -o ./$(file) main.go

## Test:
coverage:  ## Run test coverage suite
	@go test ./... -coverprofile=cov.out
	@go tool cover -html=cov.out
	@rm cov.out

## Help:
.PHONE: help
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
