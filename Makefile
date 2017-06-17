PREFIX = /usr/local

GO_FILES = main.go $(shell find cmd -type f -print)

tpm: $(GO_FILES)
	go build

install: tpm
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp $< $(DESTDIR)$(PREFIX)/bin/tpm

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/tpm
