.DEFAULT_GOAL := help
.PHONY: run build clean help

UNAME_S = $(shell uname -s)
DEFAULT_TARGET ?= desktop
DEFAULT_ARCHi ?= linux

run: build ## Run FiberCrypto Wallet.
	@echo "Running FiberCrypto Wallet..."
	@./deploy/linux/FiberCryptoWallet

install-deps-no-envs: ##  Install whithout 
	go get -v -tags=no_env github.com/therecipe/qt/cmd/...

install-docker-deps: ## Install docker images for project compilation using docker
	@echo "Downloading images..."
	docker pull therecipe/qt:$(DEFAULT_ARCH)
	@echo "Download finished."

install-deps-Linux: ## Install Linux dependencies
	go get -u -v github.com/therecipe/qt/cmd/... 
	$(go env GOPATH)/bin/qtsetup -test=false
	go get -t -d -v ./...

install-deps-Darwin: ## Install osx dependencies
	xcode-select --install
	go get -u -v github.com/therecipe/qt/cmd/... 
	$(go env GOPATH)/bin/qtsetup -test=false
	go get -t -d -v ./...

install-deps-Windowns: ## Install Windowns dependencies
	go get -u -v github.com/therecipe/qt/cmd/... 
	%v\bin\qtsetup -test=false
	go get -t -d -v ./...

install-deps: install-deps-$(UNAME_S) ## 
	@echo "Dependencies installed"

build-docker: ## Build project using docker
	@echo "Building FiberCrypto Wallet..."
	$(GOPATH)/bin/qtdeploy -docker build $(DEFAULT_TARGET)
	@echo "Done."
	

build: ## Build FiberCrypto Wallet.
	@echo "Building FiberCrypto Wallet..."
	# Add the flag `-quickcompiler` when making a release
	@qtdeploy build $(DEFAULT_TARGET)
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

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
