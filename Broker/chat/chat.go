package chat

import (
	"log"

	"golang.org/x/net/context"
)

var ipActual int = 1

type Server struct {
}

func (s *Server) RecibirDeAdmin(ctx context.Context, in *Message) (*Message, error) { //cuando un admin envia una peticion
	log.Printf("Administrador solicita servidor DNS. Petición: %s", in.Mensaje)
	if ipActual == 1 {
		ipActual++
		return &Message{Mensaje: ":9001"}, nil //10.10.28.155
	} else if ipActual == 2 {
		ipActual++
		return &Message{Mensaje: ":9002"}, nil //10.10.28.156
	} else {
		ipActual = 1
		return &Message{Mensaje: ":9003"}, nil //10.10.28.157
	}
	return &Message{Mensaje: "10.10.28.157"}, nil //No debería llegar aqui
}

//La gracia seria que una vez se Activa RecibirDeCliente, mande la accion al broker y luego reciba la respuesta
// y esa respuesta sea enviada de nuevo al cliente
// func sendToBroker(accion string, puerto string) string {
// 	var conn *grpc.ClientConn
// 	conn, err := grpc.Dial(puerto, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("No se logró conectar: %s", err)
// 	}
// 	defer conn.Close()

// 	c := chat.NewChatServiceClient(conn)

// 	response, err := c.RecibirDeCliente(context.Background(), &chat.Message{Mensaje: accion})
// 	if err != nil {
// 		log.Fatalf("Error al tratar de conectar con el Broker: %s", err)
// 	}
// 	return response.Mensaje
// }

func (s *Server) RecibirDeCliente(ctx context.Context, in *Message) (*Message, error) { //cuando un admin envia una peticion
	log.Printf("Cliente solicita servidor DNS. Petición: %s", in.Mensaje)
	if ipActual == 1 {
		ipActual++
		// IpResponse := sendToBroker(in.Mensaje, ":9001")
		return &Message{Mensaje: "Test 1"}, nil //10.10.28.155
	} else if ipActual == 2 {
		ipActual++
		return &Message{Mensaje: ":9002"}, nil //10.10.28.156
	} else {
		ipActual = 1
		return &Message{Mensaje: ":9003"}, nil //10.10.28.157
	}
	return &Message{Mensaje: "10.10.28.157"}, nil //No debería llegar aqui
}
