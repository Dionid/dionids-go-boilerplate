# Dionids Go boilerplate

# Prerequisites

1. Rename `go-boiler` to `your_app_name` in `go.mod`, `proto/go-boiler`, `docker-compose.yaml` and `PROJECT_NAME` in `Makefile`
1. Rename `cmd/core` to `cmd/${your_app_name}`
1. Create `cmd/${your_app_name}/app.env` from `cmd/${your_app_name}/app.env.example`
1. Create `.env` from `.env.example`
1. Create `internal/int-tests/test.env` from `internal/int-tests/test.env.example`
1. `make setup`
1. `make run`

# Whats inside

1. API
    1. Protobuf
    1. gRPC Gateway
    1. gRPC HTTP Gateway
    1. gRPC to Swagger
1. DB
    1. Fully typed-safe SQL on [sqli](https://github.com/Dionid/sqli)
    1. Migrations
    1. Introspection
    1. PG
1. [FOP](https://fop.davidshekunts.com)
1. [FDD](https://fdd.davidshekunts.com)
1. Error handling with [terrors](./pkg/terrors)
1. Unit & Integration tests with DB setup
1. Graceful-shutdown
1. pre-commit

# Stack

1. DB
    1. [sqli](https://github.com/Dionid/sqli)
1. Transport
    1. protobuf
    1. echo
    1. grpc
    1. cmux
1. Utils
    1. zap
    1. jwt
    1. uuid
    1. crypto
1. Development
    1. go (1.24.2)
    1. golangci
    1. docker
    1. testify

# Folder structure

1. `/api` – compiled protobuf files
1. `/bench` – benchmarks
1. `/cmd` – applications
1. `/dbs` – databases (migrations, fixtures, introspection, models)
1. `/docker` – files for docker setup
1. `/features` – business logic
1. `/for-setup` – golang setup files
1. `/internal` – internal packages
    1. `/int-tests` – integration tests
    1. `/auth` – internal auth package
1. `/pkg` – packages that can be used as a library for another project
1. `/proto` – protobuf files
    1. `/go-boiler` – application protobuf files for go-boiler
1. `/scripts` – custom scripts to run
1. `/vendor` – vendor folder


# How to add new Feature

1. Add `${feature_name}CallRequest` and `${feature_name}CallResponse` to `/proto/go-boiler/calls.proto`
1. Add `rpc ${feature_name}` to `/proto/go-boiler/calls.proto` to `MainApi`
1. Run `make generate-protobuf`
1. Add file `features/${feature_name}/${feature_name}.go`
1. Write business logic in it
1. Add method to `cmd/${your_app_name}$/http/grpc.go`
1. Run `make run`

# How to create Feature integration tests

1. Create `features/${feature_name}/${feature_name}_test.go` from `features/test-template_test.go`
1. (!) All integration tests functions must be named as `TestInt...`
1. `make test-int` (run 2 or 3 times if docker containers were restarted)

# How to create Feature unit tests

1. Create `features/${feature_name}/${feature_name}_test.go`
1. (!) All integration tests functions must be named as `TestUnit...`
1. `make test-unit` (run 2 or 3 times if docker containers were restarted)

# How to add migration

1. Run `make maindb-migrate-create name=${migration_name}` OR for go `make maindb-migrate-create name=${migration_name} migration-type=go`
1. Write migrations in `/dbs/maindb/migrations/${timestamp}_${migration_name}.sql`
1. Run `make maindb-migrate-up`
1. Run `make maindb-introspect-and-generate` to generate new models

# Architecture

## FOP

https://fop.davidshekunts.com

## FDD

https://fdd.davidshekunts.com

## DF

Distribution Functions are framework for building microservices. Work still is in progress.

## Terrors

Typed Errors stack-based serializable errors for simple error handling.

Extend your custom errors by `BaseErrorSt` / `PublicError` / `PrivateError`.

# Project layout

Done with best practices in mind. See https://github.com/golang-standards/project-layout.