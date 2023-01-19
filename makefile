SHELL := /bin/bash

run: 
	go run app/services/test-api/main.go | go run app/services/tooling/logfmt/main.go

tidy: 
	go mod tidy
	go mod vendor
