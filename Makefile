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
	go get ./...
	go build -o $@ ./cmd/$@

.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	goconvey -cover -port=10042 -workDir="$(realpath .)" -depth=-1


.PHONY: goapp_serve
goapp_serve:
	goapp serve ./cmd/appspot/app.yaml


.PHONY: goapp_deploy
goapp_deploy:
	goapp deploy -application shikaku-as-a-service ./cmd/appspot/app.yaml


.PHONY: release
release:
	goxc
