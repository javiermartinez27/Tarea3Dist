package chat

import (
	"log"

	"golang.org/x/net/context"
)

var ipActual int = 1

type Server struct {
}

func (s *Server) RecibirDeAdmin(ctx context.Context, in *Message) (*Message, error) { //cuando un admin envia una peticion
	log.Printf("Cliente envía etición: %s", in.Mensaje)
	//HACER ALGO
	return &Message{Mensaje: "Recibido"}, nil
}
