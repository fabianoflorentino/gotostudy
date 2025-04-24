# gotostudy

Go to Study!

## Structure of Project

```bash
.
├── adapters
│   ├── inbound
│   │   └── http
│   │       ├── controllers
│   │       │   ├── health_controller.go
│   │       │   └── user_controller.go
│   │       ├── handlers
│   │       │   └── user_handler.go
│   │       ├── helpers
│   │       │   └── user_helpers.go
│   │       └── requests
│   │           └── register_user_request.go
│   └── outbound
│       └── persistence
│           └── postgres
│               ├── task_model.go
│               ├── user_model.go
│               └── user_repository.go
├── build
│   ├── air.toml
│   ├── Dockerfile.dev
│   └── Dockerfile.prod
├── cmd
│   └── gotostudy
│       └── main.go
├── config
│   ├── _env
│   └── config.go
├── core
│   ├── domain
│   │   ├── task.go
│   │   └── user.go
│   ├── ports
│   │   └── user_repository.go
│   └── services
│       └── user_service.go
├── database
│   └── postgres.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── container.go
│   └── server
│       └── http.go
├── LICENSE
├── README.md
├── repositories
│   └── health_repository.go
├── services
│   └── health_services.go
├── tmp
│   └── build-errors.log
└── tree.log

26 directories, 30 files
```
