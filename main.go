package main

import (
	"log"
	"net"

	pb "github.com/KazuyaMito/task-manager-backend/gen/go/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTaskServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("falied to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, &server{})
	log.Println("Server is ruunning on port: 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
