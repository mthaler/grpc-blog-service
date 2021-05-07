package main

import (
	"fmt"
	"github.com/mthaler/grpc-blog-service/blogpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	blogpb.UnimplementedBlogServiceServer
}

func main() {
	fmt.Println("Starting calculator service...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
