package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	message := "OK\n"

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		log.Fatal(err)
		return
	}

	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		conn.Write([]byte(message))
		conn.Close()
	}
}
