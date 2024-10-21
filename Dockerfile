FROM golang:1.23-alpine AS build

WORKDIR /work
ENV GOROOT=/usr/local/go

# Download and cache dependencies.
COPY go.mod go.sum ./

RUN go mod download

# Add sources

COPY api/ ./api
COPY cmd/ ./cmd
COPY pkg/ ./pkg
COPY Taskfile.yaml .

RUN env GOBIN=/bin go install github.com/go-task/task/v3/cmd/task@latest

FROM build AS build-server

RUN task build:server


FROM alpine:latest AS server

COPY --from=build-server /work/out/server /opt/grpc-fs/server
ENTRYPOINT ["/opt/grpc-fs/server"]


