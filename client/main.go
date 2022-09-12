package main

import (
	"context"
	"github.com/chieund/golang_grpc/chat"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9002", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client"})
	if err != nil {
		log.Fatalln("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
