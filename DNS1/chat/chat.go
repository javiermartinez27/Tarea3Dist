package chat

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/context"
)

var ipActual int = 1

type Server struct {
}

func crearRegistro(registro string, ip string) {
	registroSeparado := strings.Split(registro, ".")
	nombre := "registros_zf/registro_" + registroSeparado[1] + ".txt"
	f, err := os.OpenFile(nombre, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := f.WriteString(registro + " IN A " + ip + "\n")
	if err2 != nil {
		log.Fatal(err)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func updateReloj(registro string) {
	registroSeparado := strings.Split(registro, ".")
	nombre := "relojes/reloj_" + registroSeparado[1] + ".txt"
	if _, err := os.Stat(nombre); err == nil { //actualiza el reloj
		reloj, err := readLines(nombre)
		if err != nil {
			log.Fatal(err)
		}
		separar := strings.Split(reloj[0], ",")
		i, err := strconv.Atoi(separar[0]) //cambiar en DNS2 y 3
		i++
		s := strconv.Itoa(i)
		newReloj := s + ",0,0" //cambiar en otros DNS
		err = ioutil.WriteFile(nombre, []byte(newReloj), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	} else { //primera vez que se añade algo a este registro
		f, err := os.OpenFile(nombre, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		_, err2 := f.WriteString("1,0,0")
		if err2 != nil {
			log.Fatal(err)
		}
	}
}

func borrarRegistro(registro string) {
	registroSeparado := strings.Split(registro, ".")
	nombre := "registros_zf/registro_" + registroSeparado[1] + ".txt"
	input, err := ioutil.ReadFile(nombre)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, registro) {
			lines[i] = ""
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(nombre, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func crearLog(accion string, registro string, ip string) {
	registroSeparado := strings.Split(registro, ".")
	nombre := "logs/log_" + registroSeparado[1] + ".txt"
	f, err := os.OpenFile(nombre, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if accion == "create" {
		_, err2 := f.WriteString(accion + " " + registro + " " + ip + "\n")
		if err2 != nil {
			log.Fatal(err)
		}
	} else if accion == "delete" {
		_, err2 := f.WriteString(accion + " " + registro + "\n")
		if err2 != nil {
			log.Fatal(err)
		}
	} else {
		cambioSeparado := strings.Split(ip, ">")
		ipOName := strings.ReplaceAll(cambioSeparado[1], "<", "")
		_, err2 := f.WriteString(accion + " " + registro + " " + ipOName + "\n")
		if err2 != nil {
			log.Fatal(err)
		}
	}
}

func updateRegistro(registro string, cambio string) {
	registroSeparado := strings.Split(registro, ".")
	nombre := "registros_zf/registro_" + registroSeparado[1] + ".txt"
	input, err := ioutil.ReadFile(nombre)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, registro) {
			cambioSeparado := strings.Split(cambio, ">")
			if cambioSeparado[0] == "<IP" {
				ipSep := strings.ReplaceAll(cambioSeparado[1], "<", "")
				lines[i] = registro + " IN A " + ipSep
			} else {
				guardaIp := strings.Split(lines[i], " IN A ")
				newName := strings.ReplaceAll(cambioSeparado[1], "<", "")
				lines[i] = newName + " IN A " + guardaIp[1]
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(nombre, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) RecibirDeAdmin(ctx context.Context, in *Message) (*Message, error) { //cuando un admin envia una peticion
	log.Printf("Administrador envía petición: %s", in.Mensaje)
	separar := strings.Split(in.Mensaje, " ")
	if separar[0] == "create" {
		crearRegistro(separar[1], separar[2])
		updateReloj(separar[1])
		crearLog(separar[0], separar[1], separar[2])
	} else if separar[0] == "update" {
		updateRegistro(separar[1], separar[2])
		updateReloj(separar[1])
		crearLog(separar[0], separar[1], separar[2])
	} else {
		borrarRegistro(separar[1])
		updateReloj(separar[1])
		crearLog(separar[0], separar[1], "-")
	}
	return &Message{Mensaje: "Recibido"}, nil
}
