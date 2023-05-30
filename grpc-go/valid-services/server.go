package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"

	"google.golang.org/grpc"
	pb "sinhnx.dev/grpc/valid-service/valid"
)

type Server struct {
	pb.ValidServiceServer
}

func (s *Server) ValidEmail(ctx context.Context, in *pb.StringData) (*pb.BoolValue, error) {
	log.Printf("Receive string to valid from client: %s", in.Data)
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	var rtValue pb.BoolValue
	rtValue.Value = re.MatchString(in.Data)
	return &rtValue, nil
}

var (
	port = flag.Int("port", 2011, "The server port")
)

func main() {
	fmt.Println("Go gRPC Valid Service Server is running...")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterValidServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
