.PHONY: build clean tool lint help utest config

SHELL := /bin/bash

BASEDIR = $(shell pwd)

config:
	echo ${BASEDIR}

utest:
	cd ${BASEDIR}; \
	go mod tidy; \
	export PROJECT_ENV="unit" ; \
    go test $(`go list ./... | grep -v /api/ | grep -v /docs/ | grep -v /website/ | grep -v /scripts/ | grep -v /assets/ | grep -v /web/  | grep -v /deployments/  | grep -v /test/mock/ `) -coverprofile=coverage.data ./...  ;  \

lint:
	golangci-lint run -c .golangci.yml;


help:
	@echo "make: compile packages and dependencies"
	@echo "make config: get config info"
	@echo "make lint: golint ./..."
	@echo "make utest: go test"