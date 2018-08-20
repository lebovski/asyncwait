SHELL := /bin/bash

all: test

test:
	go test -test.v github.com/lebovski/asyncwait

.PHONY: all
