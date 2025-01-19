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
	serverReader := bufio.NewReader(conn)
	for {
		// get the server response
		serverResponse, err := serverReader.ReadString('>')
		if err != nil {
			fmt.Println("Error reading server response:", err)
			return
		}
        if serverResponse[0:2] == "I:" {
            serverResponse = serverResponse[2:] 
            fmt.Print(serverResponse)
            getInput(conn)
        } else {
            fmt.Println(serverResponse[:len(serverResponse)-1])
        }
    }
}


func getInput(conn net.Conn) {
    reader := bufio.NewReader(os.Stdin)
    clientMessage, inputErr := reader.ReadString('\n')
    if inputErr != nil {
        fmt.Println("Error reading input:", inputErr)
        return
    }
    // write the message to the connection
    _, err := conn.Write([]byte(clientMessage))
    if err != nil {
        fmt.Println("Error sending message:", err)
        return
    }
}
