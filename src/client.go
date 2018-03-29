package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	gen "github.com/minond/demo-rpc/src/gen"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50007"
)

var (
	create = flag.String("create", "", "Name of your new friend.")
	search = flag.String("search", "", "Search for friends by their names.")
)

func init() {
	flag.Parse()
}

func main() {
	if *create == "" && *search == "" {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		return
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to RPC server: %v", err)
	}

	defer conn.Close()
	client := gen.NewFriendsClient(conn)
	ctx := context.Background()

	if *create != "" {
		friend := gen.Friend{Name: *create}

		if ack, err := client.Create(ctx, &friend); err != nil {
			log.Printf("Error making RPC call: %v\n", err)
		} else if ack != nil {
			fmt.Printf("Ok: %t\n", ack.Ok)
			fmt.Printf("Message: %s\n", ack.Msg)
		} else {
			log.Println("Did not get an ack back from all to Create")
		}
	} else if *search != "" {
		search := gen.SearchRequest{Name: *search}

		if res, err := client.Search(ctx, &search); err != nil {
			log.Printf("Error making RPC call: %v\n", err)
		} else {
			for {
				friend, err := res.Recv()

				if err == io.EOF {
					break
				} else if err != nil {
					log.Printf("Error making RPC call: %v\n", err)
					break
				} else if friend == nil {
					break
				} else {
					fmt.Printf("Result: %v\n", friend)
				}
			}
		}
	}
}
