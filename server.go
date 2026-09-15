package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func handleConn(conn net.Conn, message string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	err := conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Println("Error setting write deadline:", err)
		return
	}

	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println("Error writing to client:", err)
		return
	}
}

func main() {

	message := "OK\n"

	var wg sync.WaitGroup

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	fmt.Println("Server is listening on port 8080...")

	go func() {

		<-ctx.Done()

		fmt.Println("\nShutting down server...")

		if err := listener.Close(); err != nil {
			log.Println("Error closing listener: ", err)
		}

	}()

	for {
		conn, err := listener.Accept()

		if err != nil {
			if ctx.Err() != nil {
				break
			}

			fmt.Println("Error accepting connection:", err)
			continue
		}

		wg.Add(1)
		go handleConn(conn, message, &wg)
	}

	wg.Wait()

	fmt.Println("Server stopped")
}
