package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
    serverType = "tcp"
    host = "localhost"
    port = "8080"
    serverAddr = host + ":" + port
)

func main() {
	// connect to the server
	conn, err := net.Dial(serverType, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server")
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")
	// open reader for std in for user input and for reading the server
	// connection
	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)
	for {
		// Ask user for a message
		fmt.Print("Enter a message to send: ")
		clientMessage, _ := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		// write the message to the connection
		_, err = conn.Write([]byte(clientMessage))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
		// get the server response
		serverResponse, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading server response:", err)
			return
		}
		fmt.Print("Server response: " + serverResponse)
	}
}
