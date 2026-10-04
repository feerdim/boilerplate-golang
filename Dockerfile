# ==========================================
# Stage 1: Build binary
# ==========================================
FROM golang:1.27.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-w -s -extldflags '-static'" \
    -o /app/bin/server ./cmd/server/main.go && \
    CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-w -s -extldflags '-static'" \
    -o /app/bin/migrate ./cmd/migrate/main.go

# ==========================================
# Stage 2: Production runtime
# ==========================================
FROM alpine:3.24.1

RUN apk add --no-cache ca-certificates tzdata && \
    addgroup -S appgroup -g 10001 && \
    adduser -S appuser -u 10001 -G appgroup

WORKDIR /app

COPY --from=builder /app/bin/server /app/server
COPY --from=builder /app/bin/migrate /app/migrate
COPY --from=builder /app/database/migration ./database/migration
COPY --from=builder /app/src ./src

USER 10001:10001

EXPOSE 8000

CMD ["/app/server"]
