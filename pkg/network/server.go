package network

import (
    "bufio"
    "fmt"
    "net"
)

// this func handles all the individual client connections 
func handleConnection(conn net.Conn, messages chan string) {
    // make sure connection is closed on return
    defer conn.Close()
    
    // create a reader for the connection
    reader := bufio.NewReader(conn)

    for {
        // read messages from the client connection and log errors
        msg, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Connection closed:", err)
            return
        }

        // send the received message to the message channel
        messages <- msg
    }
}

func StartServer(address string, messages chan string) error {
    // start a server listening on port 8080
    listener, err := net.Listen("tcp", address)

    // log errors
    if err != nil {
        return fmt.Errorf("Error starting server: %v", err)
    }
    // close listener on return
    defer listener.Close()

    fmt.Println("Server listening on", address)
    
    // accept connections or throw error
    for {
        conn, err := listener.Accept()
        
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        // handle each connection in its own go routine
        go handleConnection(conn, messages)
    }
}
