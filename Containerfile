
# syntax=docker/dockerfile:1

# ============================================================
# Stage 1: Build
# ============================================================
FROM golang:1.26-bookworm AS builder

WORKDIR /src

# Pin the versions used by your project for reproducible builds.
ARG TEMPL_VERSION=latest
ARG SQLC_VERSION=v1.30.0

# Install Templ and sqlc.
RUN go install github.com/a-h/templ/cmd/templ@${TEMPL_VERSION} && \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@${SQLC_VERSION}

# Download Go dependencies first for better build caching.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project, including:
# - main.go
# - sqlc.yaml
# - SQL schema and queries
# - .templ files
# - static/
COPY . .

# Generate type-safe Go code from SQL.
RUN sqlc generate

# Generate Go code from Templ templates.
RUN templ generate

# Build the root Go package.
# modernc.org/sqlite does not require CGO.
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/my-go-app \
    .

# ============================================================
# Stage 2: Runtime
# ============================================================
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Copy the compiled Go application.
COPY --from=builder /out/my-go-app ./my-go-app

# Copy static assets.
COPY --from=builder /src/static ./static

# Expose the application port.
EXPOSE 6969

# Run without root privileges.
USER nonroot:nonroot

ENTRYPOINT ["/app/my-go-app"]
