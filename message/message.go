package message

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SendMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)

	return &Message{Body: "Message received successfully"}, nil
}