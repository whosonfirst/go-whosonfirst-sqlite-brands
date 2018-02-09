CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-sqlite-brands
	cp *.go src/github.com/whosonfirst/go-whosonfirst-sqlite-brands/
	cp -r tables src/github.com/whosonfirst/go-whosonfirst-sqlite-brands/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

# if you're wondering about the 'rm -rf' stuff below it's because Go is
# weird... https://vanduuren.xyz/2017/golang-vendoring-interface-confusion/
# (20170912/thisisaaronland)

# see the way we're deleting the vendor-ed version of go-whosonfirst-sqlite
# from go-whosonfirst-index - if we don't do that everything fails with a 
# lot of duplicate symbol errors (20180206/thisisaaronland)

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-brands"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-sqlite"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-index"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-log"
	rm -rf src/github.com/mattn
	rm -rf src/github.com/shaxbee
	rm -rf src/github.com/whosonfirst/go-whosonfirst-sqlite/vendor/github.com/whosonfirst/go-whosonfirst-log
	rm -rf src/github.com/whosonfirst/go-whosonfirst-index/vendor/github.com/whosonfirst/go-whosonfirst-sqlite

vendor-deps: rmdeps deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt tables/*.go

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/wof-sqlite-index-brands cmd/wof-sqlite-index-brands.go
	@GOPATH=$(GOPATH) go build -o bin/wof-sqlite-query-brands cmd/wof-sqlite-query-brands.go
