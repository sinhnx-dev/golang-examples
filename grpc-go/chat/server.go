package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "sinhnx.dev/grpc/chat/chat"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	// lis, err := net.Listen("tcp", ":9000")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// grpcServer := grpc.NewServer()

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %s", err)
	// }

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := pb.Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
