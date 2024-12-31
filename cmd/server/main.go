package main

import (
    "bufio"
    "fmt"
    "net"
)

type Message struct  {
    conn net.Conn
    text []byte
}

func handleConnection(conn net.Conn, chan messages) {
    defer conn.Close()
    // read the channel, then send stuff to the client
    clients := make(map[string]net.Conn)
    msg := <-message

    clients[msg.Conn.RemoteAddr().String()] = msg.Conn
    fmt.Println("New client = ", msg.Conn.RemoteAddr().String())
}

func main() {
    listener, err := net.Listen("tcp", ":8080")

    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()

    messages := make(chan Message)

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            return
        }
        messages <- Message{conn: conn, text: nil}
        fmt.Println("Connection accepted: ", conn.RemoteAddr())
        go handleConnection(conn, messages)
    }
}
