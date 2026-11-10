package main

import (
	"context"
	"flag"
	greeterv1 "greeter-api/gen/greeter/v1"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "The address of the gRPC server.")
	name := flag.String("name", "World", "The name to greet.")
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := greeterv1.NewGreeterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &greeterv1.SayHelloRequest{
		Name: *name,
	})

	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}
