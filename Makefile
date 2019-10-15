.DEFAULT_GOAL := help
.PHONY: run build clean help lint install-linters

UNAME_S = $(shell uname -s)
DEFAULT_TARGET ?= desktop
DEFAULT_ARCH ?= linux

run: build ## Run FiberCrypto Wallet.
	@echo "Running FiberCrypto Wallet..."
	@./deploy/linux/FiberCryptoWallet

install-deps-no-envs: ##  Install whithout 
	go get -v -tags=no_env github.com/therecipe/qt/cmd/...
	go get -t -d -v ./...
	@echo "Dependencies installed"

install-docker-deps: ## Install docker images for project compilation using docker
	@echo "Downloading images..."
	docker pull therecipe/qt:$(DEFAULT_ARCH)
	@echo "Download finished."

install-deps-Linux: ## Install Linux dependencies
	sudo apt-get update
	sudo apt-get install libgl-dev -y
	go get -u -v github.com/therecipe/qt/cmd/... 
	(qtsetup -test=false | true)
	go get -t -d -v ./...

install-deps-Darwin: ## Install osx dependencies
	xcode-select --install || true
	go get -u -v github.com/therecipe/qt/cmd/... 
	qtsetup -test=false || true
	go get -t -d -v ./...

install-deps-Windows: ## Install Windowns dependencies
	go get -u -v github.com/therecipe/qt/cmd/... 
	qtsetup -test=false -ErrorAction SilentlyContinue 
	go get -t -d -v ./...

install-deps: install-deps-$(UNAME_S) install-linters ## 
	@echo "Dependencies installed"

build-docker: ## Build project using docker
	@echo "Building FiberCrypto Wallet..."
	qtdeploy -docker build $(DEFAULT_TARGET)
	@echo "Done."
	

build: ## Build FiberCrypto Wallet.
	@echo "Building FiberCrypto Wallet..."
	# Add the flag `-quickcompiler` when making a release
	@qtdeploy build $(DEFAULT_TARGET)
	@echo "Done."

clean-Windows: ## Clean project FiberCrypto Wallet.
	@echo "Cleaning project FiberCrypto Wallet..."
	Get-ChildItem $Path -Recurse | Where{$_.Name -Match "moc"} | Remove-Item
	Get-ChildItem $Path -Recurse | Where{$_.Name -Match "deploy"} | Remove-Item -recurse
	Get-ChildItem $Path -Recurse | Where{$_.Name -Match "windows"} | Remove-Item -recurse
	Get-ChildItem $Path -Recurse | Where{$_.Name -Match "rcc"} | Remove-Item -recurse
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
	# src needs separate linting rules
	golangci-lint run -c .golangci.yml ./src/...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
