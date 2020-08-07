SHELL := /bin/bash

BASEDIR = $(shell pwd)

config:
	echo ${BASEDIR}

utest:
	cd ${BASEDIR}; \
	go mod tidy; \
	export PROJECT_ENV="unit" ; \
    go test $(`go list ./... | grep -v /api/ | grep -v /docs/ | grep -v /website/ | grep -v /scripts/ | grep -v /assets/ | grep -v /web/  | grep -v /deployments/ `) -coverprofile=coverage.data ./...  ;  \
lint:
	golangci-lint run -c .golangci.yml;


