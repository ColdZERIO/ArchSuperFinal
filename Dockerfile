# Build stage
FROM golang:1.25.6 AS builder

WORKDIR /src

# Copy go.mod and go.sum first for dependency download caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the full source tree and build the binary
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/app ./cmd/app

# Final stage
FROM gcr.io/distroless/base-debian11

WORKDIR /app

# Copy binary and static web assets and SQL schema
COPY --from=builder /app/app ./app
COPY --from=builder /src/web ./web
COPY --from=builder /src/pkg/db/scheduler.sql ./pkg/db/scheduler.sql

EXPOSE 7540

ENV TODO_PORT=":7540"
ENV TODO_DBFILE="scheduler.db"

CMD ["/app/app"]
