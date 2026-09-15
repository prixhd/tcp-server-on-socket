package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	conn, err := net.DialTimeout(
		"tcp",
		"localhost:8080",
		5*time.Second,
	)

	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal("Error closing connection:", err)
		}
	}()

	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatal("Error setting read deadline:", err)
	}

	reader := bufio.NewReader(conn)

	ans, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Error reading from server:", err)
	}

	if ans != "OK\n" {
		log.Fatal("Unexpected response from server:", ans)
	}

	fmt.Println("Received response from server:", ans)
	fmt.Println("\nConnection closed.")
}
