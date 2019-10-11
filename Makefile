.DEFAULT_GOAL := help
.PHONY: run build clean help lint install-linters

run:  ## Run FiberCrypto Wallet.
	@echo "Running FiberCrypto Wallet..."
	@./deploy/linux/FiberCryptoWallet

build:  ## Build FiberCrypto Wallet.
	@echo "Building FiberCrypto Wallet..."
	# Add the flag `-quickcompiler` when making a release
	@qtdeploy build desktop
	@echo "Done."

clean: ## Clean project FiberCrypto Wallet.
	@echo "Cleaning project FiberCrypto Wallet..."
	rm -rf deploy/
	rm -rf linux/
	rm -rf rcc.cpp
	rm -rf rcc.qrc
	rm -rf rcc_cgo_linux_linux_amd64.go
	rm -rf rcc_*.cpp
	rm -rf rcc__*
	find . -path "*moc.*" -delete
	find . -path "*moc_*" -delete
	@echo "Done."

test: ## Run project test suite
	go test github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	cat ./.travis/install-golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.10.2

lint: ## Run linters. Use make install-linters first.
	vendorcheck ./...
	# src needs separate linting rules
	golangci-lint run -c .golangci.yml ./src/...
	# The govet version in golangci-lint is out of date and has spurious warnings, run it separately
	# go vet -all ./...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
