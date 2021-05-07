# grpc-blog-service
gRPC blog service from gRPC Master Class: Build Modern API &amp; Microservices on Udemy 

Install the protobuf compiler plugins for Go using the following commands:

```bash
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Update PATH so that the protoc compiler can find the plugins:

```bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Compile a gRPC service

```bash
$ protoc --go_out=. --go-grpc_out=.  blogpb/blog.proto
```
