default: test build

pwd = $(shell pwd)

test:
	for d in helpers vault cli ; do \
		cd $(pwd)/$$d GOPATH=$(GOPATH) ; \
		go test || exit 1 ; \
	done

build: deps
	./scripts/build || exit 1

deps:
	GOBIN=$(GOPATH)/bin GOPATH=$(GOPATH) go get

.PHONY: build
