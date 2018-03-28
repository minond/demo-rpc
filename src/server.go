package main

import (
	"errors"
	"log"
	"net"

	pb "github.com/minond/demo-rpc/src/friends"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50007"
)

var (
	db = make(map[string]pb.Friend)
)

type server struct{}

func (s *server) Search(*pb.SearchRequest, pb.Friends_SearchServer) error {
	return errors.New("Unimplemented method")
}

func (s *server) Create(pb.Friends_CreateServer) error {
	return errors.New("Unimplemented method")
}

func main() {
	conn, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to open tcp connection: %v", err)
	}

	serv := grpc.NewServer()
	pb.RegisterFriendsServer(serv, &server{})
	reflection.Register(serv)

	if err := serv.Serve(conn); err != nil {
		log.Fatalf("Failed to mount RPC service to connection: %v", err)
	}
}
