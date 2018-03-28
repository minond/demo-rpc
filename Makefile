.PHONY: proto server

build:
	protoc -I src/ src/friends.proto --go_out=plugins=grpc:src/gen

server:
	go run src/server.go
