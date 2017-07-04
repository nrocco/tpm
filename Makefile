BIN := tpm
PKG := github.com/nrocco/tpm
VERSION := $(shell git describe --tags --always --dirty)
PKG_LIST := $(shell go list ${PKG}/... | grep -v ${PKG}/vendor/)
GO_FILES := $(shell find * -type d -name vendor -prune -or -name '*.go' -type f | grep -v vendor)

LDFLAGS = "-d -s -w -X ${PKG}/cmd.Version=${VERSION} -X ${PKG}/pkg/client.Version=${VERSION}"
BUILD_ARGS = -a -tags netgo -installsuffix netgo -ldflags $(LDFLAGS)

PREFIX = /usr/local

.DEFAULT_GOAL: build/$(BIN)

build/$(BIN): $(GO_FILES)
	CGO_ENABLED=0 go build ${BUILD_ARGS} -o build/${BIN} ${PKG}

.PHONY: deps
deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

.PHONY: lint
lint:
	@for file in ${GO_FILES}; do golint $${file}; done

.PHONY: vet
vet:
	@go vet ${PKG_LIST}

.PHONY: test
test:
	@go test ${PKG_LIST}

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: clean
clean:
	rm -rf build

install: build/$(BIN)
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp $< $(DESTDIR)$(PREFIX)/bin/tpm
	cp completion.zsh $(DESTDIR)$(PREFIX)/share/zsh/site-functions/_tpm

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/tpm
	rm -f $(DESTDIR)$(PREFIX)/share/zsh/site-functions/_tpm

.PHONY: build
build:
	mkdir -p build
	for GOOS in darwin linux; do \
		for GOARCH in amd64; do \
		    echo "==> Building ${BIN}-$$GOOS-$$GOARCH"; \
			docker run --rm -v "$(PWD)":/go/src/$(PKG) -w /go/src/$(PKG) -e "CGO_ENABLED=0" -e "GOOS=$$GOOS" -e "GOARCH=$$GOARCH" golang:1.9 \
				go build ${BUILD_ARGS} -o build/${BIN}-$$GOOS-$$GOARCH ${PKG}; \
		done; \
	done
