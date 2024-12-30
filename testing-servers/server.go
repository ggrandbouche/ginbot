package main

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

func main() {
    // start a server listening on port 8080
    listener, err := net.Listen("tcp", ":8080")

    // log errors
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    // close listener on return
    defer listener.Close()

    fmt.Println("Server listening on port 8080")

    // make the channel to handle client messages
    messages := make(chan string)
    // go routine to print the messages received from clients
    go func() {
        for msg := range messages {
            fmt.Println("Message received:", msg)
        }
    }()

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
