PREFIX = /usr/local

GO_FILES = main.go $(shell find cmd pkg -type f -print)

tpm: $(GO_FILES)
	go build

install: tpm
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp $< $(DESTDIR)$(PREFIX)/bin/tpm
	cp completion.zsh $(DESTDIR)$(PREFIX)/share/zsh/site-functions/_tpm

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/tpm
	rm -f $(DESTDIR)$(PREFIX)/share/zsh/site-functions/_tpm
