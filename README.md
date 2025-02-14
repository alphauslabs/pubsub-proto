
You need to install few things to start contributing:

- The [protoc](https://protobuf.dev/installation/) tool, a protocol buffer compiler.

- The following compiler plugins:

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```