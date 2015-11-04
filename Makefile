SOURCES := $(shell find . -name "*.go")
BINARIES := shikakugen


all: build

.PHONY: build
build: $(BINARIES)

.PHONY: test
test:
	go get -t ./...
	go test -v .

.PHONY: cover
cover:
	rm -f profile.out
	go test -covermode=count -coverpkg=. -coverprofile=profile.out

$(BINARIES): $(SOURCES)
	go build -o $@ ./cmd/$@
