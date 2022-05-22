# grpc-test

## Example of gRPC API usage in Client-Server communication

1. **Unary Service Method**
2. **Persistant user data**
3. **Protojson**
4. **Redis DB**
5. **From REST to gRPC**

- Get packages:

```bash
$ go get -u github.com/joho/godotenv github.com/golang/protobuf/protoc-gen-go google.golang.org/grpc
```

- Generate gRPC code

```bash
$ export PATH="$PATH:$HOME/go/bin"
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/userManagement.proto
```

## Links

- [gRPC](https://grpc.io/)
- [Used tutorial](https://youtube.com/playlist?list=PLrSqqHFS8XPYu-elDr1rjbfk0LMZkAA4X)
