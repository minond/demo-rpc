package main

import (
	"context"
	"errors"
	"log"
	"net"

	gen "github.com/minond/demo-rpc/src/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50007"
)

var (
	db = make(map[string]gen.Friend)
)

type server struct{}

func (s *server) Search(*gen.SearchRequest, gen.Friends_SearchServer) error {
	return errors.New("Unimplemented method")
}

func (s *server) Create(context.Context, *gen.Friend) (*gen.Ack, error) {
	return nil, errors.New("Unimplemented method")
}

func main() {
	conn, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to open tcp connection: %v", err)
	}

	serv := grpc.NewServer()
	gen.RegisterFriendsServer(serv, &server{})
	reflection.Register(serv)

	if err := serv.Serve(conn); err != nil {
		log.Fatalf("Failed to mount RPC service to connection: %v", err)
	}
}
