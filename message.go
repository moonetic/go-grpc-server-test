package main

import (
	"github.com/moonetic/grpc-proto-test"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SendMessage(ctx context.Context, messageToSend *message.Message) (*message.Message, error) {
	log.Printf("Received message body from client: %s", messageToSend.Body)

	return &message.Message{Body: "Message received successfully"}, nil
}
