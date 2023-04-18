ifneq (, $(shell which tput))
	GREEN  := $(shell tput -Txterm setaf 2)
	YELLOW := $(shell tput -Txterm setaf 3)
	WHITE  := $(shell tput -Txterm setaf 7)
	CYAN   := $(shell tput -Txterm setaf 6)
	RESET  := $(shell tput -Txterm sgr0)
endif

GO=go

## Git version
GIT_VERSION ?= $(shell git describe --abbrev=6 --always --tags)


BINARY_NAME=relayers
VERSION?=$(GIT_VERSION)
CI?=false

OUTPUT_DIR?="./out"
LOGS_DIR?="./logs"

CONF_PATH?=$(shell pwd)
CONF_FILE?="config.json"


.DEFAULT_GOAL := all

.PHONY: all test build

all: clean build test


## Bootstrap:
bootstrap:  ## Bootstrap project (sub-tasks: conf)
bootstrap: conf

config.json:
	@cp ./config.json.local ./$(CONF_FILE)

conf:	## Generate a local development configuration
conf: config.json

deps:	## Install dependencies
	$(GO) get ./...
	$(GO) get -t ./...

## Format:
format:	## Format source files (alias: fmt)
	$(GO) fmt ./...

fmt: format


## Test:
test:	## Run all tests (unit and integration tests)
ifeq ($(CI), false)
test: test-unit test-integration
else
# If we are in a CI environment, we only run unit tests
test: test-unit
endif

test-unit:	## Run unit tests
	$(GO) test -v -race ./relayer-srv/... $(OUTPUT_OPTIONS)

test-integration:	## Run integration tests
	FILE_PATH=$(CONF_PATH) FILE_NAME=$(CONF_FILE) $(GO) test -v -race ./test/... $(OUTPUT_OPTIONS)


## Build:
build:	## Build main (output dir: ./out/bin/relayers)
	@mkdir -p $(OUTPUT_DIR)
	CGO_ENABLED=1 $(GO) build -ldflags="-w -s" -o $(OUTPUT_DIR)/bin/$(BINARY_NAME) ./main.go

clean: 	## Remove output files
	@rm -fr $(OUTPUT_DIR) $(LOGS_DIR)

container:	## Build the docker container image
	docker build -t $(BINARY_NAME):$(VERSION) .


## Run:
run:	## Run the binary (with local configuration)
	FILE_PATH=$(CONF_PATH) FILE_NAME=$(CONF_FILE) $(GO) run ./main.go


## Help:
help:	## Show this help
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)