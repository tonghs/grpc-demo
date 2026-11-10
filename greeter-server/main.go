package main

import (
	"context"
	"flag"
	"fmt"
	greeterv1 "greeter-api/gen/greeter/v1"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type greeterServer struct {
	greeterv1.UnimplementedGreeterServiceServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *greeterv1.SayHelloRequest) (*greeterv1.SayHelloResponse, error) {
	log.Printf("Received SayHello request: %q", req.GetName())

	return &greeterv1.SayHelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func main() {
	addr := flag.String("addr", ":50051", "The address to listen on for gRPC requests.")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greeterv1.RegisterGreeterServiceServer(s, &greeterServer{})

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		log.Println("Shutting down server...")
		s.GracefulStop()
	}()

	log.Printf("Server listening on %s", *addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Server stopped.")
}
