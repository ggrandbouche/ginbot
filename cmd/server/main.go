package main

import (
	"bufio"
	"fmt"
	"net"
)

type Message struct {
	conn    net.Conn
	message string
}

func main() {
    // open listener
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started")
    // initialize channel and connections slice
	ch := make(chan Message)
	var connections []net.Conn
    // anonymous go routine to handle reading other connections messages
    // only anonymous cuz passing stuff is annoying
	go func() {
		for msg := range ch {
			for _, conn := range connections {
				if conn != msg.conn {
                    // if the message in the channel isnt from the current connection
                    // then write it to the current connection
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
    // accept connections
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
    // get messages from each client and put in channel
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}
		ch <- Message{conn: conn, message: msg}
	}
}

