package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)
	ans, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}

	if ans != "OK\n" {
		log.Fatal("Unexpected response from server:", ans)
		return
	}
	fmt.Println("Received response from server:", ans)
	fmt.Println("\nConnection closed.")
}
