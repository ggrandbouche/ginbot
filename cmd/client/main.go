package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "18.189.13.250:8080")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Connected to the server. Type your message:")

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        msg := scanner.Text() + "\n"
        _, err := conn.Write([]byte(msg))
        if err != nil {
            fmt.Println("Error sending message:", err)
            return
        }
    }
}
