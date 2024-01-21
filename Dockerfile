FROM golang:1.21 AS builder

WORKDIR /src
COPY . .
RUN go mod download

RUN GOARCH=amd64 GOOS=linux go build -o /app ./cmd/core/

FROM scratch

COPY --from=builder /app /app
EXPOSE 8080

ENTRYPOINT ["/app"]