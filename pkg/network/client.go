package network

import (
    "bufio"
    "fmt"
    "net"
)

func ConnectToServer(address string) (net.Conn, error) {
    // create connection to the server at the provided address 
    conn, err := net.Dial("tcp", address)
    if err != nil {
        return nil, fmt.Errorf("Error connection to server: %v", err)
    }
    return conn, nil
}

func SendMessage(conn net.Conn, msg string) error {
    _, err := conn.Write([]byte(msg + "\n"))
    return err
}

func ReceiveMessage(conn net.Conn) (string, error) {
    reader := bufio.NewReader(conn)
    msg, err := reader.ReadString('\n')
    return msg, err
}
