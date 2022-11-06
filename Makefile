.PHONY: default help

default: help
help: ## help: display make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m make %-20s -> %s\n\033[0m", $$1, $$2}'

# make: app info
APP_NAME     := nanoid
APP_WORKDIR  := $(shell pwd)
APP_PACKAGES := $(shell go list -f '{{.Dir}}' ./...)
APP_LOG_FMT  := `/bin/date "+%Y-%m-%d %H:%M:%S %z [$(APP_NAME)]"`
APP_VERSION  ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "0.0.1")

# --------------------------------------------------
# Test Targets
# --------------------------------------------------
COVERAGE_DIR := $(APP_WORKDIR)/coverage

# unit coverage
UNIT_DIR      := $(COVERAGE_DIR)/unit
UNIT_WEBPAGE  := $(UNIT_DIR)/index.html
UNIT_COVERAGE := $(UNIT_DIR)/coverage.txt

.PHONY: test-clean
test-clean: ## test: clean test workspace
	@echo $(APP_LOG_FMT) "cleaning workspace"
	@rm -rf $(COVERAGE_DIR)

.PHONY: test-lint
test-lint: ## test: check for lint failures
	@echo $(APP_LOG_FMT) "checking for lint failures"
	@golangci-lint run --fix -v

.PHONY: test-unit
test-unit: ## test: execute unit test suite
	@echo $(APP_LOG_FMT) "executing unit test suite"
	@mkdir -p $(UNIT_DIR)
	@go test \
		$(APP_PACKAGES) \
		-v \
		-race \
		-covermode=atomic \
		-coverprofile=$(UNIT_COVERAGE)
	@go tool cover -func=$(UNIT_COVERAGE)
	@go tool cover -html=$(UNIT_COVERAGE) -o $(UNIT_WEBPAGE)
