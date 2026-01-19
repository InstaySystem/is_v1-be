FROM golang:1.25-alpine AS builder

RUN apk add --no-cache tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download && go mod verify

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -trimpath \
    -o main ./cmd/api

FROM alpine:latest

RUN apk add --no-cache tzdata

RUN addgroup -S -g 1001 go && adduser -S -D -u 1001 -G go gin

WORKDIR /app

RUN mkdir -p logs && chown gin:go logs

COPY --from=builder --chown=gin:go /app/main .

USER gin

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./main"]