# Golang gRPC Sederhana (Client & Server)

Berikut adalah gRPC sederhana (client & server) menggunakan Golang. Repo ini didasarkan pada panduan [gRPC Quickstart](https://grpc.io/docs/quickstart/go.html) and [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html)

Repo ini mengimplementasikan gRPC sederhana dengan fungsi berikut:
- simple RPC
- server-side streaming RPC
- client-side streaming RPC
- bidirectional streaming RPC

# Cara menjalankan aplikasi

1. Install the dependencies

```bash
go mod tidy
```

2. Run the server

```bash
go run server/main.go
```

3. Run the client

```bash
go run client/main.go
```

# Cara mengenerate _grpc.pb dan .pb

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```
OR -

```bash
protoc --go_out=. --go_opt=module=github.com/taufiqoo/learn-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/taufiqoo/learn-grpc proto/greet.proto
```