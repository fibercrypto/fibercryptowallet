.DEFAULT_GOAL := help
.PHONY: all build
.PHONY: test_unit test_integration test
.PHONY: dep mocks
.PHONY: clean lint check format

GOPATH  ?= $(HOME)/go
UNAME_S ?= $(shell uname -s)
VERSION_RAW  =$(shell ./ci-scripts/version.sh)
VERSION_MAJOR=$(shell echo $(VERSION_RAW) | tr -d v | cut -d. -f1)
VERSION_MINOR=$(shell echo $(VERSION_RAW) | cut -d. -f2)
VERSION_PATCH=$(shell echo $(VERSION_RAW) | cut -d. -f3)
FULL_VERSION =$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_PATCH)

all: build

build:
	cd cmd/cli && ./install.sh

dep: ## Ensure package dependencies are up to date
	dep ensure -v

mocks: ## Create all mock files for unit tests
	echo "Generating mock files"
	mockery -name Devicer -dir ./src/skywallet -case underscore -inpkg -testonly
	mockery -name DeviceDriver -dir ./src/skywallet -case underscore -inpkg -testonly

test-unit: ## Run unit tests
	go test -v github.com/skycoin/hardware-wallet-go/src/skywallet

test-integration-emulator: ## Run emulator integration tests
	./ci-scripts/integration-test.sh -a -m EMULATOR -n emulator-integration

test-integration-wallet: ## Run usb integration tests
	./ci-scripts/integration-test.sh -m USB -n wallet-integration

test: test_unit test-integration-emulator ## Run all tests

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	# For some reason this install method is not recommended, see https://github.com/golangci/golangci-lint#install
	# However, they suggest `curl ... | bash` which we should not do
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

check-version:
	test "$(shell cat ./VERSION)" = "$(shell ./ci-scripts/version.sh)"

check: lint test ## Perform self-tests

lint: ## Run linters. Use make install-linters first.
	vendorcheck ./...
	golangci-lint run -c .golangci.yml ./...

format: ## Formats the code. Must have goimports installed (use make install-linters).
	goimports -w -local github.com/skycoin/hardware-wallet-go ./cmd
	goimports -w -local github.com/skycoin/hardware-wallet-go ./src

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

release: build
	cp $(GOPATH)/bin/skycoin-hw-cli skycoin-hw-cli-$(UNAME_S)-v$(VERSION_RAW)
