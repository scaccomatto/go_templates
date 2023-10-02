# Golang simple app

This can be a base structure for golang module or simple app. It is not a web server.  

In the Makefile you can find some useful commands for creating mocks, run tests, build...  

This is a "layered architecture" application. Here a quick description:  
* cmd: it contains the main file that points to the application staring point.
* config: it contains a config file. 
* pkg: it contains the "public data".
* internal: it contains all code that I want to keep it private: business logic, services, etc...

``` 
├── cmd
│   └── main.go
├── config
│   └── app_config.yaml
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── app.go
│   ├── configuration
│   │   └── application.go
│   ├── database
│   │   ├── database.go
│   │   ├── mapdb.go
│   │   └── mysql.go
│   ├── service
│   │   ├── foo.go
│   │   └── foo_test.go
│   └── test
│       ├── mocks
│       │   └── database
│       │       └── database_mock.go
│       └── test_fooservice.go
├── Makefile
├── pkg
│   └── dto
│       └── dto.go
└── README.md

``` 

