SOURCES := $(shell find . -name "*.go")
BINARIES := shikakugen
VERSION := $(shell cat .goxc.json | jq -c .PackageVersion | sed 's/"//g')


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


.PHONY: build-docker
build-docker: contrib/docker/.docker-container-built
	@echo "now you can 'docker push moul/shikaku'"


dist/latest/shikaku_latest_linux_386: $(SOURCES)
	mkdir -p dist
	rm -f dist/latest
	(cd dist; ln -s $(VERSION) latest)
	goxc -bc="linux,386" xc
	cp dist/latest/shikaku_$(VERSION)_linux_386 $@


contrib/docker/.docker-container-built: dist/latest/shikaku_latest_linux_386
	cp $< contrib/docker/shikaku
	docker build -t moul/shikaku:latest contrib/docker
	docker tag -f moul/shikaku:latest moul/shikaku:$(shell echo $(VERSION) | sed 's/\+/plus/g')
	docker run -it --rm moul/shikaku --width=4 --height=4 --blocks=5 --draw-map --draw-solution
	docker inspect --type=image --format="{{ .Id }}" moul/shikaku > $@.tmp
	mv $@.tmp $@
