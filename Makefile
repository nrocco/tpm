GO_FILES = main.go $(shell find cmd -type f -print)

tpm: $(GO_FILES)
	go build
