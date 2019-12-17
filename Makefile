.DEFAULT_GOAL := help
.PHONY: deps install-deps-no-envs install-docker-deps install-deps install-linters
.PHONY: install-deps-Linux install-deps-Darwin install-deps-Windows
.PHONY: build build-docker build-icon
.PHONY: prepare-release gen-mocks
.PHONY: run help
.PHONY: test test-core test-sky test-sky-launch-html-cover test-cover lint
.PHONY: clean-test clean-build clean clean-Windows

# Application info (for dumping)
ORG_DOMAIN		:= simelo.tech.org
ORG_NAME		:= Simelo.Tech
APP_NAME		:= FiberCryptoWallet
APP_DESCRIPTION	:= Multi-coin cryptocurrency wallet
APP_VERSION		:= 0.27.0
LICENSE			:= GPLv3
COPYRIGHT		:= Copyright © 2019 $(ORG_NAME)

UNAME_S = $(shell uname -s)
OSNAME = $(shell echo $(UNAME_S) | tr A-Z a-z)
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

# Files
GIT := $(shell which git)
ifeq ($(GIT),)
  ALLFILES    := $(shell find . type f | grep -v .git | grep -v vendor)
else
	GIT_BRANCH  := $(shell git rev-parse --abbrev-ref HEAD)
	ALLFILES    := $(shell git ls-tree -r $(GIT_BRANCH) --name-only | grep -v vendor)
endif

GOFILES    := $(shell echo "$(ALLFILES)" | grep  '.go$$')
QRCFILES   := $(shell echo "$(ALLFILES)" | grep  '.qrc$$')
QMLFILES   := $(shell echo "$(ALLFILES)" | grep  '.qml$$')
TSFILES    := $(shell echo "$(ALLFILES)" | grep  '.ts$$')
SVGFILES   := $(shell echo "$(ALLFILES)" | grep  '.svg$$')
JSFILES    := $(shell echo "$(ALLFILES)" | grep  '.js$$')
PNGFILES   := $(shell echo "$(ALLFILES)" | grep  '.png$$')
OTFFILES   := $(shell echo "$(ALLFILES)" | grep  '.otf$$')
ICNSFILES  := $(shell echo "$(ALLFILES)" | grep  '.icns$$')
ICOFILES   := $(shell echo "$(ALLFILES)" | grep  '.ico$$')
RCFILES    := $(shell echo "$(ALLFILES)" | grep  '.rc$$')
PLISTFILES := $(shell echo "$(ALLFILES)" | grep  'Info.plist$$')
QTCONFFILES := $(shell echo "$(ALLFILES)" | grep  'qtquickcontrols2.conf$$')

QMLFILES      := $(shell echo "$(QMLFILES) $(JSFILES)")
QTFILES       := $(shell echo "$(QRCFILES) $(TSFILES) $(PLISTFILES) $(QTCONFFILES)")
RESOURCEFILES := $(shell echo "$(SVGFILES) $(PNGFILES) $(OTFFILES) $(ICNSFILES) $(ICOFILES) $(RCFILES)")
SRCFILES      := $(shell echo "$(QTFILES) $(RESOURCEFILES) $(GOFILES)")

BINPATH_Linux      := deploy/linux/fibercryptowallet
BINPATH_Windows_NT := deploy/windows/fibercryptowallet.exe
BINPATH_Darwin     := deploy/darwin/fibercryptowallet.app/Contents/MacOS/fibercryptowallet
BINPATH            := $(BINPATH_$(UNAME_S))

deps: ## Add dependencies
	dep ensure
	rm -rf rm -rf vendor/github.com/therecipe

# Targets
run: $(BINPATH) ## Run FiberCrypto Wallet.
	@echo "Running $(APP_NAME)..."
	@$(BINPATH)

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
	wget -O magick.zip https://sourceforge.net/projects/imagemagick/files/im7-exes/ImageMagick-7.0.7-25-portable-Q16-x64.zip
	unzip magick.zip convert.exe

install-deps: install-deps-$(UNAME_S) install-linters ## Install dependencies
	@echo "Dependencies installed"

build-docker: ## Build project using docker
	@echo "Building $(APP_NAME)..."
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

build-qt:
	@echo "Building on $(UNAME_S)"
	@echo "Building $(APP_NAME)..."
# Add the flag `-quickcompiler` when making a release
	qtdeploy build $(DEFAULT_TARGET)
	@echo "Done."

$(BINPATH_Linux): $(SRCFILES)
	make build-qt

build-res-Windows_NT: $(RC_FILE)
	@echo "Building on windows"
	$(WINDRES) -i "$(RC_FILE)" -o "$(RC_OBJ)"

$(BINPATH_Windows_NT): $(SRCFILES)
	make build-icon
	make build-res-Windows_NT
	make build-qt

build-res-Darwin: $(PLIST) $(APP_ICON_PATH)/appIcon.icns
	@echo "Building on Darwin"
	mkdir -p "$(DARWIN_RES)/Content/Resources"
	cp "$(PLIST)" "$(DARWIN_RES)/Content/"
	cp "$(APP_ICON_PATH)/appIcon.icns" "$(DARWIN_RES)/Content/"

$(BINPATH_Darwin): $(SRCFILES)
	make build-icon
	make build-res-Darwin
	make build-qt

build: $(BINPATH)  ## Build FiberCrypto Wallet
	@echo "Output => $(BINPATH)"

prepare-release: ## Change the resources in the app and prepare to release the app
	./.travis/setup_release.sh

clean-test: ## Remove temporary test files
	rm -f $(COVERAGEFILE)
	rm -f $(COVERAGEHTML)

clean-build: ## Remove temporary files
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

clean: clean-test clean-build ## Remove temporary files

gen-mocks: ## Generate mocks for interface types
	mockery -all -output src/coin/mocks -outpkg mocks -dir src/core

test-sky: ## Run Skycoin plugin test suite
	go test -cover -timeout 30s github.com/fibercrypto/fibercryptowallet/src/coin/skycoin
	go test -coverprofile=$(COVERAGEFILE) -timeout 30s github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models

test-core: ## Run tests for API core and helpers
	go test -cover -timeout 30s github.com/fibercrypto/fibercryptowallet/src/util

test-sky-launch-html-cover:
	go test -cover -timeout 30s github.com/fibercrypto/fibercryptowallet/src/coin/skycoin
	go test -coverprofile=$(COVERAGEFILE) -timeout 30s github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models
	go tool cover -html=$(COVERAGEFILE) -o $(COVERAGEHTML)

test-cover-travis:
	go test -covermode=count -coverprofile=$(COVERAGEFILE) -timeout 30s github.com/fibercrypto/fibercryptowallet/src/util
	$(HOME)/gopath/bin/goveralls -coverprofile=$(COVERAGEFILE) -service=travis-ci -repotoken 1zkcSxi8TkcxpL2zTQOK9G5FFoVgWjceP
	go test -coverprofile=$(COVERAGEFILE) -timeout 30s github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models
	$(HOME)/gopath/bin/goveralls -coverprofile=$(COVERAGEFILE) -service=travis-ci -repotoken 1zkcSxi8TkcxpL2zTQOK9G5FFoVgWjceP
	go test -cover -covermode=count -coverprofile=$(COVERAGEFILE) -timeout 30s github.com/fibercrypto/fibercryptowallet/src/coin/skycoin
	$(HOME)/gopath/bin/goveralls -coverprofile=$(COVERAGEFILE) -service=travis-ci -repotoken 1zkcSxi8TkcxpL2zTQOK9G5FFoVgWjceP


test-cover: clean-test test-sky-launch-html-cover ## Show more details of test coverage

test: clean-test test-core test-sky ## Run project test suite

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	cat ./.travis/install-golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.21.0

install-coveralls: ## Install coveralls
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls

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
