# grpc-test

## Example of gRPC API usage in Client-Server communication

1. Get packages:

```bash
$ go get -u github.com/joho/godotenv github.com/golang/protobuf/protoc-gen-go google.golang.org/grpc github.com/go-redis/redis
```

2. Generate gRPC code

```bash
$ export PATH="$PATH:$HOME/go/bin"
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto
```

3. Build containers
```
$ docker-compose build
$ docker-compose up
```

## Links

- [gRPC](https://grpc.io/)
- [Tutorial for gRPC starters](https://youtube.com/playlist?list=PLrSqqHFS8XPYu-elDr1rjbfk0LMZkAA4X)
- [Tutorial for Redis, gRPC, Go microservice](https://medium.com/@felipe.infantino.moreno/microservice-in-golang-using-redis-and-grpc-865587aa260c)
