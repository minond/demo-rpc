package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	gen "github.com/minond/demo-rpc/src/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50007"
)

var (
	nextId    uint32 = 0
	friendsDb        = make(map[uint32]gen.Friend)
)

type server struct{}

func (s *server) Search(search *gen.SearchRequest, server gen.Friends_SearchServer) error {
	for _, friend := range friendsDb {
		if strings.Contains(friend.Name, search.Name) {
			server.Send(&friend)
		}
	}

	return nil
}

func (s *server) Create(ctx context.Context, friend *gen.Friend) (*gen.Ack, error) {
	if friend.Id != 0 {
		return nil, errors.New("Cannot create a new friend that has an Id.")
	}

	nextId += 1
	friend.Id = nextId
	friendsDb[nextId] = *friend

	ack := &gen.Ack{
		Ok:  true,
		Msg: fmt.Sprintf("Created %v", friend),
	}

	return ack, nil
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
