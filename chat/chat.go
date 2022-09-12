package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Println("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
