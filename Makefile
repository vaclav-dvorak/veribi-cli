GIT_REV?=$$(git rev-parse --short HEAD)
DATE?=$$(date -u +"%Y-%m-%dT%H:%M:%S")
VERSION?=$$(git describe --tags --always)
LDFLAGS="-s -w -X main.version=$(VERSION) -X main.sha=$(GIT_REV) -X main.date=$(DATE)"
goos?=$$(go env GOOS)
goarch?=$$(go env GOARCH)
file:=veribi
package:=$(file)_$(goos)_$(goarch)

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: default help release

default: help

## Build:
prepare: ## Download depencies and prepare dev env
	@pre-commit install
	@go mod download
	@go mod vendor

build:  ## Builds the cli binary
	go build -trimpath -ldflags=$(LDFLAGS) -o ./bin/$(file) .

build-ci: ## Optimized build for CI
	@echo $(goos)/$(goarch)
	go build -trimpath -ldflags=$(LDFLAGS) -o ./bin/$(file) .
	@cp LICENSE bin/LICENSE
	cd ./bin && tar -czf $(package).tar.gz ./$(file) ./LICENSE && cd ./..
	@rm bin/LICENSE

## Test:
coverage:  ## Run test coverage suite
	@go test ./... -coverprofile=cov.out
	@go tool cover -html=cov.out
	@rm cov.out

## Help:
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
