package main

import (
	"context"
	"fmt"
	"log"

	gen "github.com/minond/demo-rpc/src/gen"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50007"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to RPC server: %v", err)
	}

	defer conn.Close()
	client := gen.NewFriendsClient(conn)
	ctx := context.Background()
	friend := gen.Friend{Name: "Marcos"}

	if x, err := client.Create(ctx, &friend); err != nil {
		log.Printf("Error making RPC call: %v\n", err)
	} else {
		fmt.Printf("%+q\n", x)
	}
}
