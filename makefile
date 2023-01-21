SHELL := /bin/bash

run: 
	go run app/services/sales-api/main.go | go run app/tooling/logfmt/main.go

migrate: 
	go run app/tooling/admin/main.go

tidy: 
	go mod tidy
	go mod vendor

# Testing Auth
# curl -il http://localhost:3000/v1/testauth
# curl -il -H "Authorization: Bearer ${TOKEN}" http://localhost:3000/v1/testauth
