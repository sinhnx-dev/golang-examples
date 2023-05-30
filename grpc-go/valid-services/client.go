package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"sinhnx.dev/grpc/valid-service/valid"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":2011", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := valid.NewValidServiceClient(conn)

	response, err := c.ValidEmail(context.Background(), &valid.StringData{Data: "contact@sinhnx.dev"})
	if err != nil {
		log.Fatalf("Error when calling ValidEmail: %s", err)
	}
	log.Printf("Valid Email Response from server is: %t", response.Value)
}
