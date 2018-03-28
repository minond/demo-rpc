package main

import (
	pb "github.com/minond/demo-rpc/friends"
)

const port = ":50007"

var db = make(map[string]pb.Friend)

type server struct{}

// func (s *server)

func main() {
}
