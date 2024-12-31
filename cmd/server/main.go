package main

import (
    "bufio"
    "fmt"
    "net"
)

func handleConnection(conn net.Conn, connections *[]net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)

    fmt.Println(connections)

    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Connection closed:", err)
            return
        }

        fmt.Printf("Message received: %s", msg)

        response := "Echo: " + msg
        _, writeErr := writer.WriteString(response)
        if writeErr != nil {
            fmt.Println("Error writing to connection:", writeErr)
            return
        }

        writer.Flush()
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

    for {
        connections := []net.Conn{}
        conn, err := listener.Accept()

        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        connections = append(connections, conn)
        go handleConnection(conn)
    }
}
