FROM golang:1.20-alpine AS builder
LABEL maintainer="github:@rusq"

WORKDIR /build
COPY . .

RUN go build -ldflags="-s -w" ./cmd/aklapi

FROM alpine:3.17
LABEL maintainer="github:@rusq"

RUN apk add --no-cache ca-certificates


WORKDIR /app
COPY --from=builder /build/aklapi .

EXPOSE 8080

ENTRYPOINT ["/app/aklapi"]
