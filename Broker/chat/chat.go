package chat

import (
	"log"

	"golang.org/x/net/context"
)

var ipActual int = 1

type Server struct {
}

func (s *Server) RecibirDeAdmin(ctx context.Context, in *Message) (*Message, error) { //cuando un admin envia una peticion
	log.Printf("Cliente solicita servidor DNS. Petición: %s", in.Mensaje)
	if ipActual == 1 {
		ipActual++
		return &Message{Mensaje: "10.10.28.155"}, nil
	} else if ipActual == 2 {
		ipActual++
		return &Message{Mensaje: "10.10.28.156"}, nil
	} else {
		ipActual = 1
		return &Message{Mensaje: "10.10.28.157"}, nil
	}
	return &Message{Mensaje: "10.10.28.157"}, nil //No debería llegar aqui
}
