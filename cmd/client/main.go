package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "3.17.147.118:8080")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Connected to the server. Type your message:")

    reader := bufio.NewReader(os.Stdin)
    serverReader := bufio.NewReader(conn)
    for {
        fmt.Print("Enter message: ")
        clientMessage, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            return
        }

        _, err = conn.Write([]byte(clientMessage))
        if err != nil {
            fmt.Println("Error sending message:", err)
            return
        }

        serverResponse, err := serverReader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading server response:", err)
            return
        }

        fmt.Print("Server respone: " + serverResponse)
        
    }
}
