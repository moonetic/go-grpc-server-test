package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"moonetic.com/grpc-server-test/message"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := message.Server{}

	grpcServer := grpc.NewServer()

	message.RegisterMessageServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
