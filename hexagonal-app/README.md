# Go 

This is an example of an application with an hexagonal architecture

├── cmd
│   └── main.go
├── config
│   └── app_config.yaml
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── adapters
│   │   ├── db
│   │   │   └── db.go
│   │   └── grpc
│   │       └── grpc.go
│   ├── application
│   │   ├── api
│   │   │   └── api.go
│   │   ├── app.go
│   │   └── domain
│   │       └── order.go
│   ├── config
│   │   └── config.go
│   └── ports
│       ├── api.go
│       └── db.go
├── Makefile
└── README.md