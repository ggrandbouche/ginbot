package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting to server")
        return 
    }
    defer conn.Close()
    fmt.Println("Connected to server")
    
    reader := bufio.NewReader(os.Stdin)
    serverReader := bufio.NewReader(conn)
    for {
        fmt.Print("Enter a message to send: ")
        clientMessage, _ := reader.ReadString('\n')
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

        fmt.Print("Server response: " + serverResponse)
    }
}

