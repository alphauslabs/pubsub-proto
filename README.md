# PubSub Proto

Protocol Buffer definitions for [PubSub](https://github.com/alphauslabs/pubsub).

## Prerequisites

Before contributing to this project, you need to install:

1. The [protoc](https://protobuf.dev/installation/) tool (Protocol Buffer compiler)

2. The required compiler plugins:

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Development Workflow

To update generated files:

```bash
$ make generate_grpc_code
```
