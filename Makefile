all: shikakugen

shikakugen: $(shell find . -name "*.go")
	go build -o $@ ./cmd/$@
