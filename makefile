SHELL := /bin/bash

run: 
	go run app/services/sales-api/main.go | go run app/tooling/logfmt/main.go

key: 
	go run app/tooling/admin/main.go

tidy: 
	go mod tidy
	go mod vendor
