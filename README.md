# Dionid Go boilerplate

# Prerequisites

1. Create `.env` from `.env.example`
1. Create `./int-tests/test.env` from `./int-tests/test.env.example`
1. Run `make setup`
1. (if private gitlab) Add your token to `~/.netrc` and add `GONOSUMDB` (https://docs.gitlab.com/ee/user/packages/go_proxy/)
1. ???
1. Profit

# How to start

1. Rename `go-boiler` in `go.mod`, `proto/go-boiler` to name of your project
1. Rename `cmd/core` to name of your app
1. Create `/cmd/core/app.env` from `/cmd/core/app.env.example`
1. `make run`
1. ???
1. Profit

# Whats inside

1. FOP
1. DB
    1. Fully typed-safe SQL (raw sql on sqlc + query builder on qbik)
    1. Migrations
    1. Introspection
1. Protobuf
1. Error handling with terrors
1. Benchmarks
1. Migrations
1. Integration tests
1. PG pool
1. gRPC
1. Graceful-shutdown
1. Swagger
1. pre-commit

# Stack

1. DB
    1. qbik
    1. sqlc
    1. xo
    1. pg
1. Transport
    1. protobuf
    1. amqp091-go
    1. echo
    1. grpc
1. Utils
    1. zap
    1. jwt
    1. uuid
    1. crypto
1. Development
    1. go 1.21.6
    1. golangci
    1. docker
    1. testify

# Folder structure

1. `/api` – compiled protobuf files
1. `/bench` – benchmarks
1. `/cmd` – applications
1. `/dbs` – databases (migrations, fixtures, introspection)
1. `/docker` – files for docker setup
1. `/features` – business logic
1. `/for-setup` – golang setup files
1. `/internal` – internal packages
    1. `/int-tests` – integration tests
    1. `/auth` – internal auth package
1. `/pkg` – packages that can be used as a library for another project
1. `/proto` – protobuf files
    1. `/go-boiler` – application protobuf files for go-boiler
1. `/scripts` – scripts
1. `/vendor` – vendor folder


# How to add new Feature

1. Add `${feature_name}CallRequest` and `${feature_name}CallResponse` to `/proto/go-boiler/calls.proto`
1. Add `rpc ${feature_name}` to `/proto/go-boiler/calls.proto` to `MainApi`
1. Run `make generate-protobuf`
1. Add file `features/${feature_name}.go`
1. Write business logic in it
1. When you need SQL:
    1. You have options
        1. If static SQL
            1. Use generated qbik: `mainDb.${operation}${table_name}${something_else}(ctx, db, ...${args})`
            1. Use generated sqlc:`deps.mainDbQueries.${operation}${table_name}${something_else}(ctx, ...${args})`
            1. Use raw sql to sqlc
                1. Create `features/${feature_name}.mainDb.sql`
                1. Write raw sql with sqlc comments
                1. Run `make generate-sqlc`
                1. Use generated sqlc:`deps.mainDbQueries.${operation}${table_name}${something_else}(ctx, ...${args})`
        1. If dynamic SQL
            1. User qbik as QueryBuilder (examples in `/pkg/qbik/qbik_test.go`)
    1. Run `make maindb-introspect-and-generate`
1. Add method to `cmd/core/http/grpc.go`
1. Run `make run`
1. ???
1. Profit

# How to create Feature integration tests

1. Create `features/${feature_name}/${feature_name}_test.go` from `features/test-template_test.go`
1. !!! All integration tests functions must be named as `TestInt`
1. `make test-int` (run 2 or 3 times if docker containers were restarted)
1. ???
1. Profit

# How to create Feature unit tests

1. Create `features/${feature_name}/${feature_name}_test.go`
1. !!! All integration tests functions must be named as `TestUnit`
1. `make test-unit` (run 2 or 3 times if docker containers were restarted)
1. ???
1. Profit

# How to add migration

1. Run `make maindb-migrate-create name=${migration_name}` OR for go `make maindb-migrate-create name=${migration_name} migration-type=go`
1. Write migrations in `/dbs/maindb/migrations/${timestamp}_${migration_name}.sql`
1. Run `make maindb-migrate-up`
1. Run `make maindb-introspect-and-generate`

# Libs

## DF

Distribution Functions are framework for building microservices. Work still is in progress.

## QBik

QBik is a Query Builder for SQL. It is used to build dynamic, but type-safe SQL queries.

It introspects your database and generates code for you, based on `xo` and `sqlx``:

1. Structures, types and enums
1. Constant names like table and column names
1. SELECT on primary key, unique key, foreign key, compound index
1. UPDATE on primary key, unique key, foreign key, compound index
1. INSERT with returning on primary key, unique key, foreign key, compound index or without

All details and examples are inside `pkg/qbik`

## Terrors

Typed Errors stack-based serializable errors for simple error handling.

Extend your custom errors by `BaseErrorSt` / `PublicError` / `PrivateError`.