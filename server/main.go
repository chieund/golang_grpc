package main

import (
	"github.com/chieund/golang_grpc/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalln("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("failed to serve: %s", err)
	}
}
