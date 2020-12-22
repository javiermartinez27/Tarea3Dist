package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"google.golang.org/grpc"
)

func sendToDNS(puerto string, accion string) string {
	var conn *grpc.ClientConn
	puerto = ":9001"
	conn, err := grpc.Dial(puerto, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se logró conectar: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.RecibirDeAdmin(context.Background(), &chat.Message{Mensaje: accion})
	if err != nil {
		log.Fatalf("Error al conectar con el DNS: %s", err)
	}
	return response.Mensaje
}

func sendAccion(accion string) string {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se logró conectar: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.RecibirDeAdmin(context.Background(), &chat.Message{Mensaje: accion})
	if err != nil {
		log.Fatalf("Error al tratar de conectar con el Broker: %s", err)
	}
	return response.Mensaje
}

func main() {
	fmt.Println("Bienvenido al administrador; ingrese el comando que neecsita, o ingrese 'exit' para salir:\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		comando := scanner.Text()
		if comando == "exit" {
			fmt.Println("Saliendo...")
			break
		}

		check := strings.Split(comando, " ")
		if check[0] == "create" || check[0] == "update" || check[0] == "delete" {
			puerto := sendAccion(comando)
			fmt.Println(sendToDNS(puerto, comando))
		} else {
			fmt.Println("Por favor, ingrese un comando válido")
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Se encontró un error:", err)
	}
}
