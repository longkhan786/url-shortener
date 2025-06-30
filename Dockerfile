# Stage 1: Build Go binary with CGO in Alpine
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies for CGO
RUN apk add --no-cache build-base gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Enable CGO for go-sqlite3
ENV CGO_ENABLED=1
ENV GOOS=linux

RUN go build -o url-shortener main.go

# Stage 2: Alpine runtime
FROM alpine:latest

WORKDIR /app

# Runtime dependencies for CGO binary + SQLite
RUN apk add --no-cache sqlite-libs

COPY --from=builder /app/url-shortener .

EXPOSE 8080
CMD ["./url-shortener"]
