


## Running the code

### generate protoc 

only required if changing the .proto source

```
protoc -I learngrpc/ learngrpc/learngrpc.proto --go_out=plugins=grpc:learngrpc
```

### Start server

```
go run server/main.go
```

### Start Client

```
go run client/main.go
```