package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/randomtask1155/grpcsample/learngrpc"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewDirectorClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SayHello(ctx, &pb.HelloRequest{Name: name, Duration: 0, Cancel: false, Fail: true})
	if err != nil {
		log.Printf("Simulated Server Error: %v", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SayHello(ctx, &pb.HelloRequest{Name: name, Duration: 2, Cancel: false, Fail: false})
	if err != nil {
		log.Printf("Simulated Client timeout: %v", err)
	}

	ctx, cancel = context.WithCancel(context.Background())
	go func() {
		_, err := c.SayHello(ctx, &pb.HelloRequest{Name: name, Duration: 0, Cancel: true, Fail: false})
		if err != nil {
			log.Printf("Simulated Client cancel: %v", err)
		}
	}()
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	//log.Printf("Greeting: %s", r.Message)
}
