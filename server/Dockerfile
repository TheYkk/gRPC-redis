FROM golang:1.13-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o grpc-server

FROM scratch

COPY --from=builder /app/grpc-server /app/grpc-server

WORKDIR /app

ENTRYPOINT ["/app/grpc-server"]