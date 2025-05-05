include .env
export $(shell sed 's/=.*//' .env)

PROJECT_NAME=go-boiler
BINARY_NAME=${PROJECT_NAME}

MAIN_DB_PATH=./dbs/maindb
MAIN_DB_PG="postgres://${MAIN_PG_USERNAME}:${MAIN_PG_PASSWORD}@${MAIN_PG_HOST}:${MAIN_PG_PORT}/${MAIN_PG_DB}?sslmode=disable"
MAIN_DB_PG_TEST="postgres://${MAIN_PG_USERNAME_TEST}:${MAIN_PG_PASSWORD_TEST}@${MAIN_PG_HOST_TEST}:${MAIN_PG_PORT_TEST}/${MAIN_PG_DB_TEST}?sslmode=disable"

migration-type=sql

# DB

## MIGRATIONS

migrate-maindb-up:
	cd ./dbs/maindb/migrations && goose postgres ${MAIN_DB_PG} up

migrate-maindb-test-up:
	cd ./dbs/maindb/migrations && goose postgres ${MAIN_DB_PG_TEST} up

migrate-maindb-down:
	cd ./dbs/maindb/migrations && goose postgres ${MAIN_DB_PG} down

migrate-maindb-down-to-zero:
	cd ./dbs/maindb/migrations && goose postgres ${MAIN_DB_PG} down-to 0

migrate-maindb-test-down-to-zero:
	cd ./dbs/maindb/migrations && goose postgres ${MAIN_DB_PG_TEST} down-to 0

migrate-maindb-reup:
	make migrate-maindb-down-to-zero
	make migrate-maindb-up

migrate-maindb-test-reup:
	make migrate-maindb-test-down-to-zero
	make migrate-maindb-test-up

migrate-maindb-create:
	cd ./dbs/maindb/migrations && goose create $(name) $(migration-type)

## INTROSPECT

instrospect-maindb:
	sqli generate -o dbs/maindb ${MAIN_DB_PG}

introspect-and-generate-maindb:
	make instrospect-maindb

# Run

run:
	go fmt ./...
	docker-compose up -d
	cd ./cmd/core && go run .

prepare:
	docker-compose up -d
	make migrate-maindb-up
	make introspect-and-generate-maindb

prepare-and-run:
	make prepare
	make run

# Tests

test-unit:
	grc go test ./... -v -count=1 -run "^TestUnit" 

test-int-prepare:
	docker-compose -f docker-compose.tests.yaml up -d
	docker logs main-db-postgres-test 2>&1 | grep -q "database system is ready to accept connections"
	make migrate-maindb-test-reup

test-int:
	make test-int-prepare
	grc go test ./... -count=1 -run "^TestInt" -v

test-custom:
	make test-int-prepare
	grc go test ./... -count=1 -run $(name) -v

test-all:
	make test-int-prepare
	grc go test ./... -count=1

# Profiling

pprof:
	go tool pprof ${SERVER_HOST}/debug/pprof/${PPROF_PROFILE}

# Linter

lint:
	golangci-lint run -c ./.golangci.yml ./...

# Build

build:
	make generate-protobuf
	make build-mac && make build-linux

build-mac:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin ./cmd/core

clean-mac:
	go clean
	rm ${BINARY_NAME}-darwin

build-linux:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ./cmd/core

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux

# Protobuf and gRPC

generate-protobuf-schema:
	protoc \
		-I=./proto/${PROJECT_NAME} \
		-I=./proto/ \
		--go-grpc_out=api/v1/go/proto \
		--go_out=api/v1/go/proto \
		--go_opt paths=source_relative \
		--go-grpc_opt paths=source_relative \
		./proto/${PROJECT_NAME}/*
	protoc-go-inject-tag -input="./api/v1/go/proto/*.pb.go"

generate-protobuf-gateway:
	protoc \
		-I=./proto/${PROJECT_NAME} \
		-I=./proto/ \
		--grpc-gateway_out=api/v1/go/proto \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt logtostderr=true \
		./proto/${PROJECT_NAME}/*

generate-protobuf-openapi:
	protoc \
		./proto/${PROJECT_NAME}/calls.proto \
		-I=./proto/${PROJECT_NAME} \
		-I=./proto/ \
		--openapi_out=./api/v1/http

generate-protobuf:
	make generate-protobuf-schema
	make generate-protobuf-gateway
	make generate-protobuf-openapi

# Benchmarks

bench-select:
	ali \
	-H "Accept: application/json" \
	-H "Content-type: application/json" \
	--rate 100 \
	-w 10 \
	--body-file="./bench/select.json" \
	-m "POST" \
	${SERVER_HOST}/

# Pre-commit

pre-commit:
	make lint
	go mod tidy
	git add ./go.mod
	git add ./go.sum
	make test-all
	build
	make build-mac
	make clean-mac

setup:
	ifeq ($(UNAME_S),Linux)
		apk update && apk add --no-cache make protobuf-dev && \
		apk add --no-cache graphviz && \
		apk add --no-cache grc &&

	endif
	ifeq ($(UNAME_S),Darwin)
		brew install protobuf
		brew install graphviz
		brew install grc
	endif
	echo "make pre-commit" > .git/hooks/pre-commit
	chmod ug+x .git/hooks/pre-commit
	cp -R ./for-setup/.grc/ ~/.grc
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/favadi/protoc-go-inject-tag@latest \
		github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2 \
		github.com/nakabonne/ali@latest \
		github.com/google/gnostic/cmd/protoc-gen-openapi@latest \
		github.com/Dionid/sqli/cmd/sqli@latest
	go install -mod=mod github.com/bufbuild/buf/cmd/buf
	make prepare
