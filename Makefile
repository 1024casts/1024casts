SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with verison infos
versionDir = "github.com/1024casts/1024casts/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

PROJECT_NAME := "github.com/1024casts/1024casts"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

all: build
dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

run_ci: test install_lint lint
lint:
	golangci-lint run
install_lint:
	curl ... | sh -s ...
test: ## Run unittests
	go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

build: dep
	@go build -v -ldflags ${ldflags} .

clean:
	rm -f backend
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}

swag-init:
	swag init
	@echo "swag init done"
	@echo "see docs by: http://localhost:8080/swagger/index.html"

# gen model struct
model:
	db2struct --host localhost -d database_name -t table_name --package model --struct model_name -p --user root --gorm --json

ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"
	@echo "make swag-init - init swag doc"

.PHONY: clean gotool ca help

# more: https://github.com/xiaomeng79/istio-micro/blob/master/Makefile


