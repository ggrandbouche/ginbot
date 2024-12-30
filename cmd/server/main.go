package main

import (
    "fmt"
    "log"
    "github/ggrandbouche/ginbot/pkg/network"
)

func main() {
    serverAddress := "3.139.65.34:8080"
    messages := make(chan string)

    go func() {
        if err := network.StartServer(serverAddress, messages); err != nil {
            log.Fatal("Server error:", err)
        }
    }()

    go func() {
        for msg := range messages {
            fmt.Println("Received message:", msg)
        }
    }()

    select {}
}
