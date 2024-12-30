package main

import (
    "fmt"
    "log"
    "github/ggrandbouche/ginbot/pkg/network"
)

func main() {
    serverAddress := "18.118.160.185:8080"

    conn, err := network.ConnectToServer(serverAddress)
    if err != nil {
        log.Fatal("Error connecting to server:", err)
    }
    defer conn.Close()

    message := "Hello from client"
    if err := network.SendMessage(conn, message); err != nil {
        log.Fatal("Error sending message:", err)
    }

    response, err := network.ReceiveMessage(conn)
    if err != nil {
        log.Fatal("Error receiving message:", err)
    }
    fmt.Println("Server response:", response)
}
