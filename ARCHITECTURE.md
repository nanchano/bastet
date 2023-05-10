# Service Architecture and Design
The Bastet service provides a REST layer to interact with events, allowing the client to perform Create, Read, Update and Delete on them. It follows clean architecture principles to make developing, maintaining and debugging the service easier. This document provides an overview of the major components, how they fit together, implementation details, and some design decisions.

## App structure
```bash
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── bastet
│       └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── core
│   │   ├── errors.go
│   │   ├── event.go
│   │   ├── operations.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── repository
│   │   ├── migrations
│   │   │   ├── 000001_create_events_table.down.sql
│   │   │   └── 000001_create_events_table.up.sql
│   │   ├── operations.go
│   │   ├── repository.go
│   │   ├── sqlc
│   │   │   ├── db.go
│   │   │   ├── models.go
│   │   │   ├── querier.go
│   │   │   └── queries.sql.go
│   │   └── utils.go
│   └── server
│       ├── event.go
│       ├── operations.go
│       ├── response.go
│       └── server.go
└── sql
    ├── init.sql
    ├── queries.sql
    ├── schema.sql
    └── sqlc.yaml
```

1. ./Dockerfile containerizes the app while docker-compose.yml orchestrates the containers.
2. ./Makefile provides a user friendly layer on some typical commands run on the repo.
3. ./cmd/bastet/main.go provides the main entrypoint of the service.
4. ./internal is a Go specific folder that encapsulates application specific code, restricting its imports into different projects.

  * ./internal/config defines the configuration needed for the app, database and server.

  * ./internal/core represents the main business use case of the app, creates the relevant abstractions to perform them, disregarding implementation details to the lower level layers. It does contain business logic and validation.

  * ./internal/repository contains the repository (database) implementation, in this case for Postgres, heavily based on the SQLC tool. Migrations are included here.

  * ./internal/server represents the REST layer of the service for the client to communicate.

5. ./sql contains the SQLC YAML definitions as well as the schema and queries.

Overall, the idea is to encapsulate the business logic in the `core` package (including validation and errors), while creating the necessary abstractions for the service to work. In this case, the only abstraction is the `Repository` interface, hiding the implementation details and delegating said responsibility to the relevant packages.

The `Repository` interface is implemented in the `repository` package using Postgres SQL (which is heavily based on SQLC). If we were to change the database, it would be as simple as creating a new implementation with the new one, and instantiating it in main.

The `Server` is then injected into the server, which will parse HTTP requests, call the service to perform the operations, and send the relevant responses.

As such, the `main` function serves only as an orchestrator, instantiating the relevant components of the app and injecting them between themselves as needed.

A small caveat, theres a direct dependency on the `slog` package to log rather than hide it over a `Logger` interface to avoid over-engineering, as logging is something that doesn't usually change and the implementation of the interfaces can be tricky and time consuming.
