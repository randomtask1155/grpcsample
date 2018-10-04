package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/randomtask1155/grpcsample/learngrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	if in.Fail {
		return &pb.HelloReply{Message: "Hello " + in.Name}, fmt.Errorf("failing for no reason at all")
	}

	if in.Cancel {
		select {
		case <-ctx.Done():
			return &pb.HelloReply{}, ctx.Err()
		}
	}

	time.Sleep(time.Duration(in.Duration) * time.Second)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDirectorServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
