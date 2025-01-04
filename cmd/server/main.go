package main

import (
	"bufio"
	"fmt"
	"net"
    "os"
)

type Message struct {
	conn    net.Conn
	message string
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started")

	ch := make(chan Message)
	var connections []net.Conn

	go func() {
		for msg := range ch {
			for _, conn := range connections {
				if conn != msg.conn {
                    if msg.message == "quit" {
                        os.Exit(1)
                    }
					writer := bufio.NewWriter(conn)
					_, err := writer.WriteString(msg.message)
					if err != nil {
						fmt.Println("Error writing to connection:", err)
						continue
					}
					writer.Flush()
				}
			}
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err)
            continue
		}
        connections = append(connections, conn)
		go handleConnection(conn, ch)
	}
}

func handleConnection(conn net.Conn, ch chan Message) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr().String())
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}
		ch <- Message{conn: conn, message: msg}
	}
}

