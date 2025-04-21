# gotostudy

Go to Study!

## Structure of Project

```bash
.
├── adapters
│   ├── inbound
│   │   └── http
│   │       └── controllers
│   │           └── user_controller.go
│   └── outbound
│       └── persistence
│           └── user_repository.go
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
├── controllers
│   └── health_controller.go
├── core
│   ├── domain
│   │   ├── task.go
│   │   └── user.go
│   ├── ports
│   │   ├── user_repository.go
│   │   └── user_service.go
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
```
