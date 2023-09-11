PROJECT := amclient
.DEFAULT_GOAL := amclienttest
UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)
CACHE_BASE ?= $(HOME)/.cache/$(PROJECT)
CACHE := $(CACHE_BASE)/$(UNAME_OS)/$(UNAME_ARCH)
CACHE_BIN := $(CACHE)/bin
CACHE_VERSIONS := $(CACHE)/versions
CACHE_GOBIN := $(CACHE)/gobin
CACHE_GOCACHE := $(CACHE)/gocache
export PATH := $(abspath $(CACHE_BIN)):$(PATH)
SHELL := /usr/bin/env bash -o pipefail
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory

MOCKGEN_VERSION ?= 0.2.0
MOCKGEN := $(CACHE_VERSIONS)/mockgen/$(MOCKGEN_VERSION)
$(MOCKGEN):
	@rm -f $(CACHE_BIN)/mockgen
	@mkdir -p $(CACHE_BIN)
	@env GOBIN=$(CACHE_BIN) go install go.uber.org/mock/mockgen@v$(MOCKGEN_VERSION)
	@chmod +x $(CACHE_BIN)/mockgen
	@rm -rf $(dir $(MOCKGEN))
	@mkdir -p $(dir $(MOCKGEN))
	@touch $(MOCKGEN)

.PHONY: amclienttest
amclienttest: $(MOCKGEN)
	mockgen -typed -destination=./amclienttest/mock_ingest.go -package=amclienttest go.artefactual.dev/amclient IngestService
	mockgen -typed -destination=./amclienttest/mock_processing_config.go -package=amclienttest go.artefactual.dev/amclient ProcessingConfigService
	mockgen -typed -destination=./amclienttest/mock_transfer.go -package=amclienttest go.artefactual.dev/amclient TransferService
	mockgen -typed -destination=./amclienttest/mock_v2_jobs.go -package=amclienttest go.artefactual.dev/amclient JobsService
	mockgen -typed -destination=./amclienttest/mock_v2_package.go -package=amclienttest go.artefactual.dev/amclient PackageService
	mockgen -typed -destination=./amclienttest/mock_v2_task.go -package=amclienttest go.artefactual.dev/amclient TaskService
