package main

import (
	"context"
	"log"
	"net"

	pb "example.com/pb"
	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.GreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply {
		Message: req.Name,
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// Create new gRPC server instance
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})
	
	log.Println("Server start : 50051")

	// Start gRPC server
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}