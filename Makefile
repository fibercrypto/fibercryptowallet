.DEFAULT_GOAL := help
.PHONY: run build-icon build clean help

# Application info (for dumping)
ORG_DOMAIN		:= simelo.tech.org
ORG_NAME		:= Simelo.Tech
APP_NAME		:= FiberCrypto Wallet
APP_DESCRIPTION	:= Multi-coin cryptocurrency wallet
APP_VERSION		:= 0.27.0
LICENSE			:= GPLv3
COPYRIGHT		:= Copyright © 2019 $(ORG_NAME)

# Icons
APP_ICON_PATH	:= resources/images/icons/appIcon
ICONS_BUILDPATH	:= resources/images/icons/appIcon/build
ICONSET			:= resources/images/icons/appIcon/appIcon.iconset
CONVERT			:= convert
SIPS			:= sips
ICONUTIL		:= iconutil
UNAME_S = $(shell uname -s)

# Platform-specific switches
ifeq ($(OS),Windows_NT)
	CONVERT		= magick convert
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

# Targets
run:  ## Run FiberCrypto Wallet
	@echo "Running $(APP_NAME)..."
	@./deploy/linux/FiberCryptoWallet

build-icon-Windows_NT: ## Build the application icon in Windows
	mkdir -p $(ICONS_BUILDPATH)
	# For Windows icons we need the `convert` tool provided by "Imagemagick"
	$(CONVERT) -resize 16x16 "$(APP_ICON_PATH)/appIcon-wallet.png" "$(ICONS_BUILDPATH)/appIcon_16x16.png"
	$(CONVERT) -resize 24x24 "$(APP_ICON_PATH)/appIcon-wallet.png" "$(ICONS_BUILDPATH)/appIcon_24x24.png"
	$(CONVERT) -resize 32x32 "$(APP_ICON_PATH)/appIcon.png" "$(ICONS_BUILDPATH)/appIcon_32x32.png"
	$(CONVERT) -resize 48x48 "$(APP_ICON_PATH)/appIcon.png" "$(ICONS_BUILDPATH)/appIcon_48x48.png"
	$(CONVERT) -resize 64x64 "$(APP_ICON_PATH)/appIcon.png" "$(ICONS_BUILDPATH)/appIcon_64x64.png"
	$(CONVERT) -resize 96x96 "$(APP_ICON_PATH)/appIcon.png" "$(ICONS_BUILDPATH)/appIcon_96x96.png"
	$(CONVERT) -resize 128x128 "$(APP_ICON_PATH)/appIcon.png" "$(ICONS_BUILDPATH)/appIcon_128x128.png"
	$(CONVERT) -resize 256x256 "$(APP_ICON_PATH)/appIcon.png" "$(ICONS_BUILDPATH)/appIcon_256x256.png"
	$(CONVERT) -resize 512x512 "$(APP_ICON_PATH)/appIcon.png" "$(ICONS_BUILDPATH)/appIcon_512x512.png"
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

build-Windows_NT: ## Build FiberCrypto Wallet in Windows
	@echo "Building on windows"
	$(WINDRES) -i "$(RC_FILE)" -o "$(RC_OBJ)"


build-Darwin: ## Build FiberCrypto Wallet in Darwin 
	@echo "Building on Darwin"
	mkdir -p "$(DARWIN_RES)/Content/Resources"
	cp "$(PLIST)" "$(DARWIN_RES)/Content/"
	cp "$(APP_ICON_PATH/appIcon.icns)" "$(DARWIN_RES)/Content/"

build:  ## Build FiberCrypto Wallet
#ifeq ($(OS),Windows_NT)
#	@echo "Building $(OS) resources..."
#	$(WINDRES) -i "$(RC_FILE)" -o "$(RC_OBJ)"
#endif
#ifeq ($(OS),Darwin)
#	@echo "Building $(OS) resources..."
#	mkdir -p "$(DARWIN_RES)/Content/Resources"
#	cp "$(PLIST)" "$(DARWIN_RES)/Content/"
#	cp "$(APP_ICON_PATH/appIcon.icns)" "$(DARWIN_RES)/Content/"
#endif
	@echo "Building $(APP_NAME)..."
# 	Add the flag `-quickcompiler` when making a release
	qtdeploy build desktop
	@echo "Done."

clean-Windows_NT: ## Clean project in Windows
	# Windows actions
	rm -rf "$(ICONS_BUILDPATH)"
	rm -rf "$(RC_OBJ)"

clean-Darwin: ## Clean project in Darwin
	# Darwin actions
	rm -rf "$(ICONSET)"

clean-Linux: ## Clean project in Linux
	# Linux actions
	@echo "Cleaned"

clean: clean-$(OS) ## Clean project FiberCrypto Wallet
	# Regular generated files
	@echo "Cleaning project $(APP_NAME)..."
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
	@echo "$(APP_NAME) v$(APP_VERSION)"
	@echo "$(APP_DESCRIPTION)"
	@echo "$(COPYRIGHT)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
