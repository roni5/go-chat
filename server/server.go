package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/arjunyel/go-chat"
)

const (
	port = ":12893"
)

type server struct{}

func (s *server) Chat(ctx context.Context, in *pb.SendChat) (*pb.ReceiveChat, error) {
	clientName := in.Name
	clientMessage := in.Message

	return &pb.ReceiveChat{Name: clientName, Message: clientMessage}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// Initializes the gRPC server.
	s := grpc.NewServer()

	// Register the server with gRPC.
	pb.RegisterGroupChatServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
