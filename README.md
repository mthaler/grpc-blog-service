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

## Install MongoDB

Download MongoDB from https://www.mongodb.com/try/download/community and install it. On Debian Linux, run the following commands to start `mongod`:

```bash
# systemctl enable mongod
# systemctl start mongod
# systemctl status mongod
```

The last command should show that `mongod` was started successfully.

## Install Robo 3T

[Robo 3T](https://github.com/Studio3T/robomongo) is a GUI for MongoDB. Install the latest release from GitHub releases.

## Install MongoDB Go Driver

To connect to MongoDB the blog serivce uses the [mongodb-go-driver](https://github.com/mongodb/mongo-go-driver)

To add the dependency, do

```bash
$ go get go.mongodb.org/mongo-driver/mongo
```
