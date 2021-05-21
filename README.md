This directory contains a Rest HTTP API server prototype in Go. Following stacks are picked:
	
	net/http package to start and serve HTTP server
  
	Gorilla mux to handle routes
  
  	Swagger in lorder to serve a REST API compliant with OpenAPI specs

Pre-requisits
  Install Go in 1.13 version minimum
  
  install swaggern https://github.com/go-swagger/go-swagger/blob/master/docs/install.md
  
  define api pkg/swagger/swagger.yml, generate code
  
  swagger generate server --quiet --target server --name signup-api --spec swagger.yml --exclude-main

Build the app
$ go build -o bin/rest-go-server internal/main.go

or
$ make build

Run the app
$ ./bin/rest-go-server

Test the app
$ curl http://localhost:8080/signup

