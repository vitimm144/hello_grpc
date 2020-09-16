package main

import (
	"context"
	"log"
	"net"

	"vitimm144/hello_grpc/pb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Hello (ctx context.Context, request *.pb.HelloRequest) (*pb.HelloResponse, error){
	result := "Hello "+request.GetName()
	res := &pb.HelloResponse{
		Msg: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to List %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
