
# Simple Go Service 

Simple Go API Service with logging and monitoring




## Run Locally with default values

Clone the project

```bash
  git clone git@github.com:ZweWT/backend-go.git
```

Go to the project directory

```bash
  cd backend-go
```

migration

```bash
  make migrate
```

Start the server

```bash
  make run
```
## Run project with customized values

If you want to set env variables such as dbhost, and port etc, you can fix the default config(app/services/sales-api/main.go) or you can also set it in the cmdline.

**Example**
```bash
  go run app/services/sales-api/main.go --web-api-host=127.0.0.1:8000
```

Available flags are - 

`--web-api-host`
`--web-debug-host`
`--web-read-timeout`
`--web-write-timeout`
`--web-idle-timeout`
`--web-shutdown-timeout`
`--auth-keys-folder`
`--auth-active-kid`
`--db-user`
`--db-password`
`--db-host`
`--db-name`
`--db-max-idle-conns`
`--db-max-open-conns`


## Load Tests and Monitoring

To run load tests, install Hey(https://github.com/rakyll/hey) 

```bash
  hey -m GET -c 100 -n 10000 http://localhost:3000/v1/test
```


## Monitoring & Debugging

Debug routes can be found are under port 4000 

```http
  GET /debug/pprof
  GET /debug/pprof/cmdline
  GET /debug/pprof/profile
  GET /debug/pprof/symbol
  GET /debug/pprof/trace
  GET /debug/vars
```
Monitoring via terminal, install expvarmon (https://github.com/divan/expvarmon)

```bash
  expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"
```
## Tech Stack

**Language:** Golang

**Database:** Postgresql

