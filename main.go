package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to server gRPC over port 9000: %v", err)
	}

}
