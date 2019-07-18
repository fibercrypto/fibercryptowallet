.DEFAULT_GOAL := help
.PHONY: run-wallet build-wallet clean help

run-wallet:  ## Run fiberCryptoWallet.
	./deploy/linux/FiberCryptoWallet

build-wallet:  ## Build fiberCryptoWallet.
	@echo "Building Fiber Wallet"
	@echo "..."
	@qtdeploy -debug -quickcompiler build desktop
	@echo "Wallet builded"

clean: ## Clean project.
	rm -rf deploy/
	rm -rf rcc.cpp
	rm -rf rcc.qrc
	rm -rf rcc_cgo_linux_linux_amd64.go
	rm -rf rcc_*.cpp
	rm -rf rcc__*

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'