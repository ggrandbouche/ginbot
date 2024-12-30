package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Creates a connection to the ec2 server with the public ip and the choosen
    // port
    conn, err := net.Dial("tcp", "18.118.160.185:8080")
    // if there is an error connecting we return and print the error
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    //  make sure the connection is closed once we return 
    defer conn.Close()

    fmt.Println("Connected to the server. Type your message:")

    // scanner used to read user input from terminal
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        msg := scanner.Text() + "\n"
        // send the message from use to the server, again logging errors 
        _, err := conn.Write([]byte(msg))
        if err != nil {
            fmt.Println("Error sending message:", err)
            return
        }
    }
}
