.DEFAULT_GOAL := help
.PHONY: help all clean install
.PHONY: build-go build-js build-c build-py
.PHONY: install-deps-go install-deps-js install-deps-nanopb install-protoc
.PHONY: clean-go clean-js clean-c clean-py

REPO_ROOT := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

UNAME_S  = $(shell uname -s)
PYTHON  ?= python
PIP     ?= pip3
PIPARGS ?=
GOPATH  ?= $(HOME)/go

ifeq ($(TRAVIS),true)
  OS_NAME=$(TRAVIS_OS_NAME)
else
  ifeq ($(UNAME_S),Linux)
    OS_NAME=linux
  endif
  ifeq ($(UNAME_S),Darwin)
    OS_NAME=osx
  endif
endif

PROTOC_VERSION      ?= 3.6.1
PROTOC_ZIP          ?= protoc-$(PROTOC_VERSION)-$(OS_NAME)-x86_64.zip
PROTOC_URL          ?= https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_ZIP)
PROTOC_GOGO_URL      = github.com/gogo/protobuf
PROTOC_NANOPBGEN_DIR = nanopb/vendor/nanopb/generator
ifeq ($(UNAME_S),Linux)
  SED_FLAGS = -i
else ifeq ($(UNAME_S),Darwin)
  SED_FLAGS = -i '' -e
endif

PROTOB_SPEC_DIR = protob
PROTOB_MSG_DIR  = $(PROTOB_SPEC_DIR)/messages

PROTOB_GO_DIR = go
PROTOB_JS_DIR = js
PROTOB_PY_DIR = py
PROTOB_C_DIR  = c
PROTOB_SRC_DIR  = $(GOPATH)/src/$(PROTOC_GOGO_URL)

# Use default value when including this repository in vendor/ with dep
# Set it explicitly either when generating to acustom location 
# or when this repository is included at a subpath other than vendor
# e.g. go library should set this to github.com/skycoin/hardware-wallet-go/src/device-wallet/messages
GO_IMPORT ?= github.com/fibercrypto/skywallet-protob
GO_IMPORT_SED = $(shell echo $(GO_IMPORT) | sed 's/\//\\\//g')

PROTOB_MSG_FILES = $(shell ls -1 $(PROTOB_MSG_DIR)/*.proto)
PROTOB_MSG_SPECS = $(patsubst %,$(PROTOB_MSG_DIR)/%,$(notdir $(PROTOB_MSG_FILES)))
PROTOB_MSG_GO    = $(patsubst %,$(PROTOB_GO_DIR)/%,$(notdir $(PROTOB_MSG_FILES:.proto=.pb.go)))
PROTOB_MSG_JS    = $(patsubst %,$(PROTOB_JS_DIR)/%,$(notdir $(PROTOB_MSG_FILES:.proto=.pb.js)))
PROTOB_MSG_PY    = $(patsubst %,$(PROTOB_PY_DIR)/%,$(notdir $(PROTOB_MSG_FILES:.proto=_pb2.py)))
PROTOB_MSG_C     = $(patsubst %,$(PROTOB_C_DIR)/%,$(notdir $(PROTOB_MSG_FILES:.proto=.pb.c)))

# Output dirs

OUT_GO ?= $(PROTOB_GO_DIR)
OUT_JS ?= $(PROTOB_JS_DIR)
OUT_PY ?= $(PROTOB_PY_DIR)
OUT_C  ?= $(PROTOB_C_DIR)

all: build-go build-js build-c build-py ## Generate protobuf classes for all languages

install: install-deps-go install-deps-js install-deps-nanopb install-protoc ## Install protocol buffer tools

clean: clean-go clean-js clean-c clean-py ## Delete temporary and output files
	rm -rf \
		$$( find . -name '*.swp' ) \
		$$( find . -name '*.swo' ) \
		$$( find . -name '*.orig' )

install-protoc: /usr/local/bin/protoc

/usr/local/bin/protoc:
	echo "Downloading protobuf from $(PROTOC_URL)"
	curl -OL $(PROTOC_URL)
	echo "Installing protoc"
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
	rm -f $(PROTOC_ZIP)

#----------------
# Go lang
#----------------

install-deps-go: install-protoc ## Install tools to generate protobuf classes for go lang
	@if [ -e $(PROTOB_SRC_DIR) ] ; then \
		echo 'Detected $(PROTOC_GOGO_URL) on local file system. Checking v1.2.0' ; \
		cd $(PROTOB_SRC_DIR) && git checkout v1.2.0 ; \
	else \
		echo 'Cloning $(PROTOC_GOGO_URL)' ; \
		git clone --branch v1.2.0 --depth 1 https://$(PROTOC_GOGO_URL) $(PROTOB_SRC_DIR) ; \
	fi
	( cd $(PROTOB_SRC_DIR)/protoc-gen-gogofast && go install )

build-go: install-deps-go $(PROTOB_MSG_GO) $(OUT_GO)/google/protobuf/descriptor.pb.go ## Generate protobuf classes for go lang

$(OUT_GO)/google/protobuf/descriptor.pb.go: $(OUT_GO)/types.pb.go
	protoc -I./$(PROTOC_NANOPBGEN_DIR)/proto --gogofast_out=$(OUT_GO) $(PROTOC_NANOPBGEN_DIR)/proto/google/protobuf/descriptor.proto
	sed $(SED_FLAGS) 's/import\ protobuf\ \"google\/protobuf\"/import\ protobuf\ \"$(GO_IMPORT_SED)\/go\/google\/protobuf\"/g' $(OUT_GO)/types.pb.go

$(OUT_GO)/%.pb.go: $(PROTOB_MSG_DIR)/%.proto
	protoc -I./$(PROTOC_NANOPBGEN_DIR)/proto/ -I protob/messages --gogofast_out=$(OUT_GO) $<

clean-go:
	rm -rf $$( find $(OUT_GO) -name '*.pb.go' )

check-go: build-go
	grep -xq 'import\ protobuf\ \"$(GO_IMPORT_SED)\/go\/google\/protobuf\"' $(OUT_GO)/types.pb.go || exit 1

#----------------
# Javascript
#----------------

install-deps-js: ## Install tools to generate protobuf classes for javascript
	cd $(REPO_ROOT)/js && npm install

build-js: install-deps-js ## Generate protobuf classes for javascript
	cd $(REPO_ROOT)/js && npm run gen-proto

clean-js:
	rm -rf $(OUT_JS)/skycoin.js $(OUT_JS)/node_modules

#----------------
# C with nanopb
#----------------

install-deps-nanopb: install-protoc ## Install tools to generate protobuf classes for C and Python with nanopb
	make -C $(PROTOC_NANOPBGEN_DIR)/proto/
	$(PIP) install $(PIPARGS) "protobuf==$(PROTOC_VERSION)" ecdsa

build-c: install-deps-nanopb $(PROTOB_MSG_C) $(OUT_C)/messages_map.h ## Generate protobuf classes for C with nanopb

$(OUT_C)/%.pb.c: $(OUT_C)/%.pb $(PROTOB_MSG_DIR)/%.options
#c/%.pb.c: c/%.pb $(PROTOB_MSG_DIR)/%.options
	$(eval PROTOBUF_FILE_OPTIONS := $(subst pb,options,$<))
	$(eval PROTOBUF_FILE_OPTIONS = $(subst c/,,$(PROTOBUF_FILE_OPTIONS)))
	$(PYTHON) $(PROTOC_NANOPBGEN_DIR)/nanopb_generator.py -f $(PROTOB_MSG_DIR)/$(PROTOBUF_FILE_OPTIONS) $< -L '#include "%s"' -T

$(OUT_C)/%.pb: $(PROTOB_MSG_DIR)/%.proto
	protoc -I./$(PROTOC_NANOPBGEN_DIR)/proto/ -I. -I./$(PROTOB_MSG_DIR) $< -o $@

$(OUT_C)/messages_map.h: $(OUT_PY)/messages_map.py $(OUT_PY)/messages_pb2.py $(OUT_PY)/types_pb2.py
	PYTHONPATH="$$PYTHONPATH:$(REPO_ROOT)/$(OUT_PY)" $(PYTHON) $< > $@

clean-c: clean-py
	rm -rf $(OUT_C)/messages_map.h \
		$$( find $(OUT_C) -name '*.pb.c' ) \
		$$( find $(OUT_C) -name '*.pb.h' ) \
		$$( find $(OUT_C) -name '*.d' ) \
		$$( find $(OUT_C) -name '*.i' ) \
		$$( find $(OUT_C) -name '*.s' ) \
		$$( find $(OUT_C) -name '*.o' )

#----------------
# Python with nanopb
#----------------

build-py: install-deps-nanopb $(PROTOB_MSG_PY) ## Generate protobuf classes for Python with nanopb

$(OUT_PY)/%_pb2.py: $(PROTOB_MSG_DIR)/%.proto
	protoc -I./$(PROTOC_NANOPBGEN_DIR)/proto/ -I./$(PROTOB_MSG_DIR) $< --python_out=$(OUT_PY)

clean-py:
	rm -rf \
		$$( find . -name '__pycache__' ) \
		$$( find . -name '*_pb2.py' ) \
		$$( find . -name '*.pyc' ) \
		$$( find . -name '*.pyd' ) \
		$$( find . -name '*.pyo' )

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

