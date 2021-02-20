package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	NETWORK string = "tcp"
	HOST    string = "localhost"
	PORT    string = "8080"
)

func processResponse(response string) string {
	return "HTTP/1.1 200 OK\r\n" + "Content-Length: " + fmt.Sprintf("%s", len(response)) + "\r\nContent-Type: text/html\r\n" + fmt.Sprintf("\fr\n%s\r\n", response)
}

func handleRequest(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	//receive the user name
	buff := make([]byte, 256)
	clientNameBytes, _ := c.Read(buff)
	clientName := string(buff[0:clientNameBytes])
	fmt.Println(clientName)

	line, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("ok")
	}
	//broadcast client message
	message := clientName + ":" + string(line)
	fmt.Println(message)
	response := processResponse("world")
	fmt.Println(response)
	c.Write([]byte(response))
}

func main() {
	ln, err := net.Listen(NETWORK, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			break
		}
		// new go coroutine
		go handleRequest(conn)
	}
}
