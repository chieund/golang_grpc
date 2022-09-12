# run build docker
```
docker-composer up --build
```

# create file pb
```
docker-composer exec grpc protoc --go_out=plugins=grpc:chat chat.proto
```

# run server
```
docker-composer exec grpc go run server/main.go
2022/09/12 15:01:18 Receive message body from client: %s Hello From Client

```

# run client
```
docker-composer exec grpc go run client/main.go
2022/09/12 15:01:33 Response from server: Hello From the Server!

```