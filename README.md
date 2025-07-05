# Project - What are we building
# gRPC key-value store service
- Set(key, value) -> saves a key-value pair
- Get(key) -> retrieve a value
- Delete(key) -> removes a key.

Everything will b e test-driver (TDD), and we'll eventually add unit tests, streaming, and auth.

# Set up
## Install `protoc`
This is the Protocol Buffer complier userd to generate Go code from `.proto` files.
```
brew install protobuf
```
Check version
```
protoc --version
```

Make sure you have the **protoc** Go plugin installed and its in on your `PATH`.
## Go gRPC Plugins needed
```shell
go install google.golang.org/protobuf/cmd/proto-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
Verify the plugins are installed
```shell
which protoc-gen-go
which protoc-gen-go-grpc
```
These should return the path to the installed binaries.
## Troubleshooting if there plugins aren't found
1. Check where Go is installed
   ```
   which go
   ```
   If it shows something like `/opt/homebrew/bin/go`, then you're using the Homebrew-installed Go
2. Check where Go installs binaries.
   ```
   go env GOPATH
   ```
   You should get something like: `/Users/<your-username>/go`
   Now Check if the `bin/` directory under this path exists:
   ```
   ls "$(go env GOPATH)/bin"
   ```
   You should ideally see binaries like `protoc-gen-go` and `protoc-gen-go-grpcf` after running the install commands below.
3. Run the plugin installation commands again from above
4. Add the binary path to your shell's `PATH`. Make sure this line is in your `~/.zshrc`
   ```
   export PATH="$PATH:$(go env GOPATH)/bin"
   ```
   Then reload your terminal or run: `source ~/.zshrc`.
5. Confirm again the plugings are installed and visible.

# Generating Go code from .proto file
## Goal of generating go code from .proto file
1. We have **Go structs** for the requests/response messages.
2. We get a **Go interface** for the gRPC server to implement.
3. We get a **Go client** type to make RPC calls.

## Setup:
```
kvstore/
├── proto/
│   └── kvstore.proto   // Your .proto file
├── gen/                // Where generated code will go
└── ...
```
Suppose this is the project structure.
You're working with a file like `proto\kvstore.proto` that defines messages and an RPC service.

The following command
```shell
protoc \
  --go_out=gen \
  --go-grpc_out=gen \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  proto/kvstore.proto
```
`protoc` is the **Protocol Buffers compiler**, used to conver `.proto` files into target language code, e.g. Go

`--go_out=gen` tell `protoc` to generate Go code for **messages (structs)** and place the files in  the `gen/` directory.

`--go-grpc_out=gen` tells `protoc` to also generate Go code for **gRPC services** - i.e., the server interface and the client stub.
This server interface needs to be implemented by the Go application you build the api for.
The client stub, can be used by your client to call APIs on your Go server.

`source_relative` is used to generate the code files relative to the `.proto` file, rather than the root.

### Output Example
```
gen/
└── kvstorepb/
    ├── kvstore.pb.go         # All the message definitions (SetRequest, GetResponse, etc.)
    └── kvstore_grpc.pb.go    # gRPC interface and client stubs
```

## Generated code files
With above command, we generated code files in Go.

### kvstore.pb.go
This file includes:
- Message Structs: 
  Example:
  ```proto
  type SetRequest struct {
    Key   string
    Value string
  }
  ```
  These are generated from the `.proto` definitions like:
  ```proto
  message SetRequest {
    string key = 1;
    string value = 2;
  }
  ```
  These implement all the protobuf marshaling/unmarshaling logc you don't need to worry about.

### kvstore_grpc.pb.go
This file includes two very important things:

✅ `KVStoreServer` **interface**
```go
type KVStoreServer interface {
    Set(context.Context, *SetRequest) (*SetResponse, error)
    Get(context.Context, *GetRequest) (*GetResponse, error)
}
```
This gets implemented in the Go server. This basically servers a Controller interface in REST world. The Go server should implement a service that would process the request and return a response.

✅ `KVStoreClient` **struct**
```go
type KVStoreClient interface {
    Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
    Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}
```
This is used by a client to make **RPC** calls to the server.

