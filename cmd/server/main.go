
package main

import (
    "bufio"
    "fmt"
    "net"
)

func handleConnection(conn net.Conn, messages chan string) {
    defer conn.Close()
    reader := bufio.NewReader(conn)

    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Connection closed:", err)
            return
        }

        messages <- msg
    }
}

func main() {
    listener, err := net.Listen("tcp", ":8080")

    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server listening on port 8080")

    messages := make(chan string)
    go func() {
        for msg := range messages {
            fmt.Println("Message received:", msg)
        }
    }()

    for {
        conn, err := listener.Accept()

        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn, messages)
    }
}
