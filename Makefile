.DEFAULT_GOAL := help
.PHONY: run-wallet build-wallet clean help

run-wallet:  ## Run FiberCrypto Wallet.
	@echo "Running FiberCrypto Wallet..."
	@./deploy/linux/FiberCryptoWallet

build-wallet:  ## Build FiberCrypto Wallet.
	@echo "Building FiberCrypto Wallet..."
	# Add the flag `-quickcompiler` when making a release
	@qtdeploy -debug build desktop
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
	find . -path "*moc.*"
	find . -path "*moc_*"
	@echo "Done."

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
