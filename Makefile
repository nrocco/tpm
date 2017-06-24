BINARY := tpm
PKG := github.com/nrocco/tpm
VERSION := $(shell git describe --always --long --dirty)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

LDFLAGS = "-X ${PKG}/cmd.Version=${VERSION} -X ${PKG}/pkg/client.Version=${VERSION}"

PREFIX = /usr/local

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(GO_FILES)
	go build -i -v -o ${BINARY} -ldflags ${LDFLAGS} ${PKG}

lint:
	@for file in ${GO_FILES}; do \
		golint $${file}; \
	done

vet:
	@go vet ${PKG_LIST}

test:
	@go test -short ${PKG_LIST}

install: tpm
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp $< $(DESTDIR)$(PREFIX)/bin/tpm
	cp completion.zsh $(DESTDIR)$(PREFIX)/share/zsh/site-functions/_tpm

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/tpm
	rm -f $(DESTDIR)$(PREFIX)/share/zsh/site-functions/_tpm

.PHONY: clean
clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi
