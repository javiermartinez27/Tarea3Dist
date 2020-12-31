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

var consistencia [][]string

func sendToBroker(accion string) string {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se logró conectar: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.RecibirDeCliente(context.Background(), &chat.Message{Mensaje: accion})
	if err != nil {
		log.Fatalf("Error al tratar de conectar con el Broker: %s", err)
	}
	return response.Mensaje
}

func main() {
	fmt.Println("Bienvenido al Cliente; ingrese el comando  get nombre.dominio, o ingrese 'exit' para salir:\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		comando := scanner.Text()
		if comando == "exit" {
			fmt.Println("Informacion de los dominios que ha solicitado:")
			fmt.Println(consistencia)
			fmt.Println("Saliendo...")
			break
		}
		check := strings.Split(comando, " ")
		if check[0] == "get" {
			mensaje := sendToBroker(comando) //string
			fmt.Println("Mensaje:")
			fmt.Println(mensaje)
			if len(consistencia) == 0 {
				paraAppend := []string{check[1], mensaje}
				consistencia = append(consistencia, paraAppend)
				// fmt.Println(consistencia)
				// fmt.Println(reflect.TypeOf(consistencia))
				// fmt.Println("if")
			} else {
				for i, s := range consistencia {
					// mensaje := sendToBroker(comando)
					// fmt.Println(s)

					fmt.Println("esto es s")
					fmt.Println(s)

					// fmt.Println("esto es s como string")
					// fmt.Println(strings.Join(s, " "))

					// fmt.Println("esto es check[1]")
					// fmt.Println(check[1])
					// fmt.Println("esto es type of check[1]")
					// fmt.Println(reflect.TypeOf(check[1]))
					// fmt.Println("esto es check[1] +  mensaje")
					// test := check[1] + " " + mensaje
					// fmt.Println(test)
					// sComoString := strings.Join(s, " ")
					// fmt.Println(strings.Contains(sComoString, check[1]))
					relojObtenidoDelMensaje := strings.Split(mensaje, " ")
					fmt.Println("Reloj Obtenido")
					fmt.Println(relojObtenidoDelMensaje[0])

					// fmt.Println("esto es mensaje+check[1]")
					// test2 := mensaje + " " + check[1]
					// fmt.Println(test2)
					if s[i] == check[1] {
						fmt.Println("pase por aqui")

						s = []string{check[1], mensaje}
						consistencia[i] = s
					} // else {
					// 	paraAppend := []string{check[1], mensaje}
					// 	consistencia = append(consistencia, paraAppend)
					// 	fmt.Println(consistencia)
					// 	fmt.Println("else")
					// }
					// // paraAppend := []string{check[1], mensaje}
					// // consistencia = append(consistencia, paraAppend)
				}
			}
			// for i, s := range consistencia {
			// 	mensaje := sendToBroker(comando)
			// 	fmt.Println(len(consistencia))

			// 	if s[i] == check[1] {
			// 		fmt.Println("lo contiene")
			// 		s = []string{mensaje}
			// 		consistencia[i] = s
			// 	} else {
			// 		fmt.Println("tessst")
			// 		paraAppend := []string{check[1], mensaje}
			// 		consistencia = append(consistencia, paraAppend)
			// 	}
			// 	paraAppend := []string{check[1], mensaje}
			// 	consistencia = append(consistencia, paraAppend)
			// }
			// mensaje := sendToBroker(comando)
			// fmt.Println(mensaje)
			// paraAppend := []string{check[1], mensaje}
			// consistencia = append(consistencia, paraAppend)
		} else {
			fmt.Println("Por favor, ingrese un comando válido")
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Se encontró un error:", err)
	}
}
