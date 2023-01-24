SHELL := /bin/bash

run: 
	go run app/services/sales-api/main.go | go run app/tooling/logfmt/main.go


# if you want to set env variables such dbhost, and port etc, you can also set it in the cmdline
# go run app/services/sales-api/main.go --web-api-host=127.0.0.1:8000
#--desc=information
# available flags are 
# --web-api-host
# --web-debug-host
# --web-read-timeout
# --web-write-timeout
# --web-idle-timeout
# --web-shutdown-timeout
# --auth-keys-folder
# --auth-active-kid
# --db-user
# --db-password
# --db-host
# --db-name=postgres
# --db-max-idle-conns
# --db-max-open-conns


migrate: 
	go run app/tooling/admin/main.go

tidy: 
	go mod tidy
	go mod vendor

# For testing load on the service
# hey -m GET -c 100 -n 10000 http://localhost:3000/v1/test

# For monitoring the service 
# expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"
