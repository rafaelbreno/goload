# Build stage
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache ca-certificates tzdata git

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X main.version=${VERSION:-dev}" \
    -trimpath \
    -o /goload

# Final stage
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /goload /goload

LABEL org.opencontainers.image.title="goload" \
      org.opencontainers.image.description="CLI tool for HTTP load testing" \
      org.opencontainers.image.vendor="rafaelbreno"

# Use non-root user (numeric UID for scratch compatibility)
USER 65534:65534

ENTRYPOINT ["/goload"]
