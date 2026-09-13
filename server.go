package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func handleConn(conn net.Conn, message string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Fatal("Error writing to connection:", err)
	}
}

func main() {

	message := "OK\n"

	var wg sync.WaitGroup

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		wg.Add(1)
		go handleConn(conn, message, &wg)
	}
}
