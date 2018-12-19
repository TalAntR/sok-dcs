#!/usr/bin/env make -f
# -*- makefile -*-
#
# This makefile builds ...
# Examples:
#
#    make -  build Golang application;
#


# Path to source code
WORKSPACE ?= $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

export GOPATH ?= $(abspath $(dir $(WORKSPACE)))

# src/github.com/TalAntR/sok-dcs
PROJECT := github.com/TalAntR/sok-dcs

GODEPS := $(addprefix github.com/,go-yaml/yaml)

default: build;


build: dcsd;

clean:
	go clean -i -testcache -x
	@rm -f $(GOPATH)/src/$(PROJECT)/.gobuild

dcsd: $(GOPATH)/src/$(PROJECT)/.gobuild $(GODEPS)
	go fmt $(PROJECT) $(PROJECT)/dcs
	go build -o $@ $(PROJECT)

$(GOPATH)/src/$(PROJECT)/.gobuild:
	@mkdir -p $(dir $(abspath $(dir $@)))
	@ln -sf -T $(WORKSPACE) $(abspath $(dir $@))
	@touch $@

# Download Go dependencies
$(GODEPS):
	go get $@

.PHONY: test
test: $(GOPATH)/src/$(PROJECT)/.gobuild
	go test ./...

help:
	go help

.PHONY: default build
