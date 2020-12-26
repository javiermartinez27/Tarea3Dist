package chat

import (
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"
)

var ipActual int = 1

type Server struct {
}

func crearRegistro(registro string) {
	registroSeparado := strings.Split(registro, ".")
	nombre := "registros_zf/registro_" + registroSeparado[1] + ".txt"
	f, err := os.OpenFile(nombre, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := f.WriteString(registro + "\n")
	if err2 != nil {
		log.Fatal(err)
	}
}

func (s *Server) RecibirDeAdmin(ctx context.Context, in *Message) (*Message, error) { //cuando un admin envia una peticion
	log.Printf("Administrador envía petición: %s", in.Mensaje)
	separar := strings.Split(in.Mensaje, " ")
	if separar[0] == "create" {
		crearRegistro(separar[1])
	} else if separar[0] == "update" {
		//updateregistro(separar[1], separar[2])
	} else {
		//borrarregistro(separar[1])
	}
	return &Message{Mensaje: "Recibido"}, nil
}
