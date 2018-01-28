BIN := tpm
PKG := github.com/nrocco/tpm
VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git describe --always --dirty)
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
PKG_LIST := $(shell go list ${PKG}/... | grep -v ${PKG}/vendor/)
GO_FILES := $(shell git ls-files '*.go')

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
LDFLAGS = "-d -s -w -X ${PKG}/cmd.version=${VERSION} -X ${PKG}/cmd.commit=${COMMIT} -X ${PKG}/cmd.buildDate=${DATE} -X ${PKG}/pkg/client.Version=${VERSION}"
BUILD_ARGS = -a -tags netgo -installsuffix netgo -ldflags $(LDFLAGS)

PREFIX = /usr/local

.DEFAULT_GOAL: build

build/${BIN}-$(GOOS)-$(GOARCH): $(GO_FILES)
	mkdir -p build
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build ${BUILD_ARGS} -o $@ ${PKG}

.PHONY: deps
deps:
	dep ensure

.PHONY: lint
lint:
	golint -set_exit_status ${PKG_LIST}

.PHONY: vet
vet:
	go vet -v ${PKG_LIST}

.PHONY: test
test:
	go test -short ${PKG_LIST}

.PHONY: coverage
coverage:
	mkdir -p coverage && rm -rf coverage/*
	for package in ${PKG_LIST}; do go test -covermode=count -coverprofile "coverage/$${package##*/}.cov" "$$package"; done
	echo mode: count > coverage/coverage.cov
	tail -q -n +2 coverage/*.cov >> coverage/coverage.cov
	go tool cover -func=coverage/coverage.cov

.PHONY: version
version:
	@echo ${VERSION}

.PHONY: clean
clean:
	rm -rf build

.PHONY: build
build: build/${BIN}-${GOOS}-${GOARCH}

.PHONY: build-all
build-all:
	$(MAKE) build GOOS=linux GOARCH=amd64
	$(MAKE) build GOOS=darwin GOARCH=amd64

.PHONY: install
install: build/$(BIN)-$(GOOS)-$(GOARCH)
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp $< $(DESTDIR)$(PREFIX)/bin/${BIN}
	cp completion.zsh $(DESTDIR)$(PREFIX)/share/zsh/site-functions/_${BIN}

.PHONY: uninstall
uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/${BIN}
	rm -f $(DESTDIR)$(PREFIX)/share/zsh/site-functions/_${BIN}
