SOURCE_FILES?=./...
GOLANGCI_VERSION=v1.52.2
COVERAGE=coverage.out

export GO111MODULE := on
export PATH := ./bin:$(PATH)

default: build

.PHONY: link-git-hooks
link-git-hooks:
	find .git/hooks -type l -exec rm {} \;
	find githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

.PHONY: build
build:
	go install $(SOURCE_FILES)

.PHONY: test
test:
	go test $(SOURCE_FILES) -coverprofile $(COVERAGE) -timeout=30s -parallel=4 -cover -race

.PHONY: fmt
fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w  ./**/*.go

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

.PHONY: lint
lint:
	golangci-lint run $(SOURCE_FILES)

.PHONY: check
check: test lint-fix

.PHONY: tools
tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_VERSION)

.PHONY: addcopy
addcopy:
	@scripts/add-copy.sh

TAG=$(patsubst v%,%,$(shell git describe --tags --dirty --always))
.PHONY: check-version
check-version:
	scripts/check-version.sh "$(TAG)"

.PHONY: openapi-pipeline
openapi-pipeline:
	echo "Running OpenAPI Generation and Validation process"
	$(MAKE) -C tools clean_client
	echo "Running client generation"
	$(MAKE) -C tools generate_client
	echo "Validating generated SDK"
	go test go.mongodb.org/atlas-sdk/test

.PHONY: generate-docs
generate-docs:
	$(MAKE) -C tools generate_docs
	./scripts/toc.sh
