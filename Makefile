# some helpful shortcuts

build:
	go install github.com/aykevl93/plaincast

fmt:
	go fmt . ./apps ./apps/youtube ./apps/youtube/mp ./server

run: build
	$(GOPATH)/bin/plaincast
