# Build stage

FROM golang:1.22-alpine as builder

WORKDIR /app
COPY . .

# Tidy and build the gRPC server from new location
RUN go mod tidy
RUN go build -o keyserver-grpc ./cmd/server/main.go

# Runtime stage
FROM alpine
COPY --from=builder /app/kvstore-grpc /usr/local/bin/kvstore-grpc

EXPOSE 50051
ENTRYPOINT ["kvstore-grpc"]