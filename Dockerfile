FROM golang:1.24.2-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache git ca-certificates tzdata && \
    update-ca-certificates

WORKDIR /src

# Copy go mod files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

RUN GOARCH=amd64 GOOS=linux go build -o /app ./cmd/core/

FROM scratch

COPY --from=builder /app /app
EXPOSE 8080

ENTRYPOINT ["/app"]