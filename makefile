SHELL := /bin/bash

run: 
	go run app/services/sales-api/main.go | go run app/tooling/logfmt/main.go

migrate: 
	go run app/tooling/admin/main.go

tidy: 
	go mod tidy
	go mod vendor

# For testing load on the service
# hey -m GET -c 100 -n 10000 http://localhost:3000/v1/test

# For monitoring the service 
# expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"
