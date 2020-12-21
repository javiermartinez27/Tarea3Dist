package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"google.golang.org/grpc"
)

func sendAccion(accion string) string {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.RecibirDeAdmin(context.Background(), &chat.Message{Mensaje: accion})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	//log.Printf("Ip: %s", response.Respuesta)
	return response.Mensaje
}

func main() {
	fmt.Println(sendAccion("create google.com 127.0.0.1"))
}
