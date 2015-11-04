SOURCES := $(shell find . -name "*.go")
BINARIES := shikakugen


all: $(BINARIES)

$(BINARIES): $(SOURCES)
	go build -o $@ ./cmd/$@
