REPO_ROOT=$(shell git rev-parse --show-toplevel)

# Build variables
PACKAGE = github.com/gezacorp/licsensei
BINARY_NAME ?= licsensei
BUILD_DIR ?= build
BUILD_PACKAGE = ${PACKAGE}/cmd/licsensei
VERSION ?= $(shell (git symbolic-ref -q --short HEAD || git describe --tags --exact-match) | tr "/" "-")
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
LDFLAGS += -X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH} -X main.buildDate=${BUILD_DATE} -X main.configDir=${CONFIG_DIR}
GOPRIVATE = github.com/gezacorp

export CGO_ENABLED ?= 0
ifeq (${VERBOSE}, 1)
ifeq ($(filter -v,${GOARGS}),)
	GOARGS += -v
endif
TEST_FORMAT = short-verbose
endif

# Dependency versions
LICENSEI_VERSION = 0.9.0
GOLANGCI_VERSION = 1.64.7
GORELEASER_VERSION = 2.8.1
GOLANG_VERSION = 1.24

.PHONY: list
list: ## List all make targets
	@${MAKE} -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: fmt
fmt: ## Run go fmt against code
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code
	go vet ./...

.PHONY: tidy
tidy: ## Execute go mod tidy
	go mod tidy
	go mod download all

${REPO_ROOT}/bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p ${REPO_ROOT}/bin
	@mkdir -p bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

${REPO_ROOT}/bin/golangci-lint: ${REPO_ROOT}/bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} ${REPO_ROOT}/bin/golangci-lint

.PHONY: lint
lint: ${REPO_ROOT}/bin/golangci-lint ## Run linter
# "unused" linter is a memory hog, but running it separately keeps it contained (probably because of caching)
	${REPO_ROOT}/bin/golangci-lint run --disable=unused -c ${REPO_ROOT}/.golangci.yml --timeout 2m
	${REPO_ROOT}/bin/golangci-lint run -c ${REPO_ROOT}/.golangci.yml --timeout 2m

.PHONY: lint-fix
lint-fix: ${REPO_ROOT}/bin/golangci-lint ## Run linter
	@${REPO_ROOT}/bin/golangci-lint run -c ${REPO_ROOT}/.golangci.yml --fix --timeout 2m

bin/goreleaser: ## Install goreleaser
	@mkdir -p ./bin/
	scripts/install_goreleaser.sh v${GORELEASER_VERSION}

.PHONY: release
release: bin/goreleaser ## Release current tag
	@bin/goreleaser

${REPO_ROOT}/bin/licensei: ${REPO_ROOT}/bin/licensei-${LICENSEI_VERSION}
	@ln -sf licensei-${LICENSEI_VERSION} ${REPO_ROOT}/bin/licensei

${REPO_ROOT}/bin/licensei-${LICENSEI_VERSION}:
	@mkdir -p ${REPO_ROOT}/bin
	@mkdir -p bin
	curl -sfL https://raw.githubusercontent.com/goph/licensei/master/install.sh | bash -s v${LICENSEI_VERSION}
	mv bin/licensei $@

.PHONY: license-check
license-check: ${REPO_ROOT}/bin/licensei ## Run dependencies license check
	${REPO_ROOT}/bin/licensei --config ${REPO_ROOT}/.licensei.toml check
	${REPO_ROOT}/bin/licensei --config ${REPO_ROOT}/.licensei.toml header

${REPO_ROOT}/${BUILD_DIR}/licsensei:
	make build

.PHONY: license-header-check
license-header-check: ${REPO_ROOT}/${BUILD_DIR}/licsensei # Run license header check
	@${REPO_ROOT}/${BUILD_DIR}/licsensei

.PHONY: build
build: ## Build a binary
ifeq (${VERBOSE}, 1)
	go env
endif
ifneq (${IGNORE_GOLANG_VERSION_REQ}, 1)
	@printf "${GOLANG_VERSION}\n$$(go version | awk '{sub(/^go/, "", $$3);print $$3}')" | sort -t '.' -k 1,1 -k 2,2 -k 3,3 -g | head -1 | grep -q -E "^${GOLANG_VERSION}$$" || (printf "Required Go version is ${GOLANG_VERSION}\nInstalled: `go version`" && exit 1)
endif

	@$(eval GENERATED_BINARY_NAME = ${BINARY_NAME})
	@$(if $(strip ${BINARY_NAME_SUFFIX}),$(eval GENERATED_BINARY_NAME = ${BINARY_NAME}-$(subst $(eval) ,-,$(strip ${BINARY_NAME_SUFFIX}))),)
	go build ${GOARGS} -tags "${GOTAGS}" -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/${GENERATED_BINARY_NAME} ${BUILD_PACKAGE}

.PHONY: test
test:
	go test ./... \
		-coverprofile cover.out \
		-v \
		-failfast \
		-test.v \
		-test.paniconexit0
