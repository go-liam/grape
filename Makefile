SHELL := /bin/bash

BASEDIR = $(shell pwd)

config:
	echo ${BASEDIR}

test:
	cd ${BASEDIR}; \
	go mod tidy; \
	export PROJECT_ENV="unit" ; \
    go test $(`go list ./... | grep -v /proto/ | grep -v /temp-demo/ | grep -v /frontend/ `) -coverprofile=coverage.data ./...  ;  \



