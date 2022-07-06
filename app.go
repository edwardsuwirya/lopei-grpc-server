package main

import "enigmacamp.com/lopei_grpc_srv/delivery"

// 1. install golang grpc
// go get -u google.golang.org/grpc
// 2. compile protobuf
// protoc --go_out=./service --go-grpc_out=./service model/lopei.proto

func main() {
	delivery.Server().Run()
}
