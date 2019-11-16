.DEFAULT_GOAL := help
.PHONY: run build-icon build clean help lint install-linters

# Application info (for dumping)
ORG_DOMAIN		:= simelo.tech.org
ORG_NAME		:= Simelo.Tech
APP_NAME		:= FiberCrypto Wallet
APP_DESCRIPTION	:= Multi-coin cryptocurrency wallet
APP_VERSION		:= 0.27.0
LICENSE			:= GPLv3
COPYRIGHT		:= Copyright © 2019 $(ORG_NAME)

UNAME_S = $(shell uname -s)
DEFAULT_TARGET ?= desktop
DEFAULT_ARCH ?= linux
## In future use as a parameter tu make command.
COIN = skycoin
COVERAGEPATH = src/coin/$(COIN)
COVERAGEFILE = $(COVERAGEPATH)/coverage.out
COVERAGEHTML = $(COVERAGEPATH)/coverage.html

# Icons
APP_ICON_PATH	:= resources/images/icons/appIcon
ICONS_BUILDPATH	:= resources/images/icons/appIcon/build
ICONSET			:= resources/images/icons/appIcon/appIcon.iconset
CONVERT			:= convert
SIPS			:= sips
ICONUTIL		:= iconutil
UNAME_S         = $(shell uname -s)
DEFAULT_TARGET  ?= desktop
DEFAULT_ARCH    ?= linux

# Platform-specific switches
ifeq ($(OS),Windows_NT)
	CONVERT		 = ./convert.exe
	WINDRES		:= windres
	RC_FILE		:= resources/platform/windows/winResources.rc
	RC_OBJ		:= winResources.syso
else
#	Get the UNIX operating system
	OS			:= $(shell uname -s)
#	If Linux...
	ifeq ($(OS),Linux)
	endif
#	If Darwin
	ifeq ($(OS),Darwin)
		DARWIN_RES	:= darwin
		PLIST		:= resources/platform/darwin/info.plist
	endif
endif

deps: ## Add dependencies
	dep ensure
	rm -rf rm -rf vendor/github.com/therecipe

# Targets
run: build ## Run FiberCrypto Wallet.
	@echo "Running $(APP_NAME)..."
	@./deploy/linux/FiberCryptoWallet

install-deps-no-envs: ## Install therecipe/qt with -tags=no_env set
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
	wget -O magick.zip -A "ImageMagick-*-portable-Q16-x64.zip" -r -nc -np -nd -e robots=off https://imagemagick.org/download/binaries
	unzip magick.zip convert.exe

install-deps: install-deps-$(UNAME_S) install-linters ## Install dependencies
	@echo "Dependencies installed"

build-docker: ## Build project using docker
	@echo "Building FiberCrypto Wallet..."
	qtdeploy -docker build $(DEFAULT_TARGET)
	@echo "Done."

build-icon-Windows_NT: ## Build the application icon in Windows
	mkdir -p $(ICONS_BUILDPATH)
	$(CONVERT) "$(APP_ICON_PATH)/appIcon-wallet.png" -resize 16x16 "$(ICONS_BUILDPATH)/appIcon_1_16x16.png"
	$(CONVERT) "$(APP_ICON_PATH)/appIcon-wallet.png" -resize 24x24 "$(ICONS_BUILDPATH)/appIcon_2_24x24.png"
	$(CONVERT) "$(APP_ICON_PATH)/appIcon.png" -resize 32x32 "$(ICONS_BUILDPATH)/appIcon_3_32x32.png"
	$(CONVERT) "$(APP_ICON_PATH)/appIcon.png" -resize 48x48 "$(ICONS_BUILDPATH)/appIcon_4_48x48.png"
	$(CONVERT) "$(APP_ICON_PATH)/appIcon.png" -resize 64x64 "$(ICONS_BUILDPATH)/appIcon_5_64x64.png"
	$(CONVERT) "$(APP_ICON_PATH)/appIcon.png" -resize 96x96 "$(ICONS_BUILDPATH)/appIcon_6_96x96.png"
	$(CONVERT) "$(APP_ICON_PATH)/appIcon.png" -resize 128x128 "$(ICONS_BUILDPATH)/appIcon_7_128x128.png"
	$(CONVERT) "$(APP_ICON_PATH)/appIcon.png" -resize 256x256 "$(ICONS_BUILDPATH)/appIcon_8_256x256.png"
	$(CONVERT) "$(APP_ICON_PATH)/appIcon.png" -resize 512x512 "$(ICONS_BUILDPATH)/appIcon_9_512x512.png"
	$(CONVERT) "$(ICONS_BUILDPATH)/appIcon_*.png" "$(APP_ICON_PATH)/appIcon.ico"

build-icon-Darwin: ## Build the application icon in Darwin
	mkdir -p $(ICONSET)
	# For macOS icons we will use the `sips` and `iconutil` tools as provided by Apple
	$(SIPS) -z 16 16 "$(APP_ICON_PATH)/appIcon-wallet.png" --out "$(ICONSET)/icon_16x16.png"
	$(SIPS) -z 32 32 "$(APP_ICON_PATH)/appIcon.png" --out "$(ICONSET)/icon_16x16@2x.png"
	$(SIPS) -z 32 32 "$(APP_ICON_PATH)/appIcon.png" --out "$(ICONSET)/icon_32x32.png"
	$(SIPS) -z 64 64 "$(APP_ICON_PATH)/appIcon.png" --out "$(ICONSET)/icon_32x32@2x.png"
	$(SIPS) -z 128 128 "$(APP_ICON_PATH)/appIcon.png" --out "$(ICONSET)/icon_128x128.png"
	$(SIPS) -z 256 256 "$(APP_ICON_PATH)/appIcon.png" --out "$(ICONSET)/icon_128x128@2x.png"
	$(SIPS) -z 256 256 "$(APP_ICON_PATH)/appIcon.png" --out "$(ICONSET)/icon_256x256.png"
	$(SIPS) -z 512 512 "$(APP_ICON_PATH)/appIcon.png" --out "$(ICONSET)/icon_256x256@2x.png"
	$(SIPS) -z 512 512 "$(APP_ICON_PATH)/appIcon.png" --out "$(ICONSET)/icon_512x512.png"
	$(SIPS) -z 1024 1024 $(APP_ICON_PATH)/appIcon.png --out "$(ICONSET)/icon_512x512@2x.png"
	$(ICONUTIL) --convert icns --output "$(APP_ICON_PATH)/appIcon.icns" "$(ICONSET)"

build-icon-Linux: ## Build the application icon in Linux
	@echo "Icons cannot be embedded in ELF executables."

build-icon: build-icon-$(OS) ## Build the application icon (Windows_NT and Darwin systems)
	@echo "Builded $(OS) icon..."

build-Linux: ## Build FiberCryptoWallet in Windows
	@echo "Building on Linux"

build-Windows_NT: build-icon ## Build FiberCrypto Wallet in Windows
	@echo "Building on windows"
	$(WINDRES) -i "$(RC_FILE)" -o "$(RC_OBJ)"


build-Darwin: build-icon ## Build FiberCrypto Wallet in Darwin
	@echo "Building on Darwin"
	mkdir -p "$(DARWIN_RES)/Content/Resources"
	cp "$(PLIST)" "$(DARWIN_RES)/Content/"
	cp "$(APP_ICON_PATH)/appIcon.icns" "$(DARWIN_RES)/Content/"

build: build-$(OS)  ## Build FiberCrypto Wallet
	@echo "Building $(APP_NAME)..."
# 	Add the flag `-quickcompiler` when making a release
	qtdeploy build $(DEFAULT_TARGET)
	@echo "Done."

prepare-release: ## Change the resources in the app and prepare to release the app
	./setup_release.sh

clean: ## Clean project FiberCrypto Wallet.
	@echo "Cleaning project $(APP_NAME)..."
	rm -rf deploy/
	rm -rf linux/
	rm -rf windows/
	rm -rf rcc.cpp
	rm -rf rcc.qrc
	rm -rf rcc_cgo_*.go
	rm -rf rcc_*.cpp
	rm -rf rcc__*
	find . -path "*moc.*" -delete
	find . -path "*moc_*" -delete
	rm -rf "$(ICONS_BUILDPATH)"
	rm -rf "$(RC_OBJ)"
	rm -rf "$(ICONSET)"

	@echo "Done."

test-sky: ## Run Skycoin plugin test suite
	go test -cover -timeout 30s github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin
	go test -coverprofile=$(COVERAGEFILE) -timeout 30s github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models

test-clean:
	rm $(COVERAGEFILE)
	rm $(COVERAGEHTML)

test-sky-launch-html-cover:
	go test -cover -timeout 30s github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin
	go test -coverprofile=$(COVERAGEFILE) -timeout 30s github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models
	go tool cover -html=$(COVERAGEFILE) -o $(COVERAGEHTML)

test-cover: test-sky-launch-html-cover ## Show more details of test coverage

test: test-sky ## Run project test suite

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	cat ./.travis/install-golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.10.2

lint: ## Run linters. Use make install-linters first.
	# src needs separate linting rules
	golangci-lint run -c .golangci.yml ./src/coin/...
	golangci-lint run -c .golangci.yml ./src/core/...
	golangci-lint run -c .golangci.yml ./src/main/...
	golangci-lint run -c .golangci.yml ./src/util/...

help:
	@echo "$(APP_NAME) v$(APP_VERSION)"
	@echo "$(APP_DESCRIPTION)"
	@echo "$(COPYRIGHT)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
