package main

import (
	"bufio"
	"fmt"
	"net"
    "github/ggrandbouche/ginbot/pkg/gin"
    "sync"
)

const (
    serverType = "tcp"
    host = "localhost"
    port = ":8080"
)

type Connection struct {
    conn net.Conn
    player int
}

func main() {
    // open listener
	listener, err := net.Listen(serverType, port)
	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started")
    // initialize channel and connections slice
    var connections []Connection
    var wg sync.WaitGroup

    input := make(chan string)
    output := make(chan gin.Output)
    startGin := make(chan string)

    wg.Add(1)
    go func() {
        defer wg.Done()
        <-startGin
        fmt.Print("started gin")
        go gin.Gin(input, output)

        for outputCP, ok := <-output; ok; outputCP, ok = <-output {
            for _, conn := range connections {
                if conn.player == outputCP.Player || outputCP.Player == 2 {
                    fmt.Println("writing, connection: ", conn.conn.LocalAddr())
                    conn.conn.Write([]byte(outputCP.Output))
                }
            }
        }
    }()

    for i := 0; i < 2; i++ {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err)
            continue
		}
        connections = append(connections, Connection{conn: conn, player: i})
		go handleConnection(conn, input)
	}
    startGin<- " "
    wg.Wait()
}

func handleConnection(conn net.Conn, input chan string) {
	defer conn.Close()
    fmt.Println("Client connected:", conn.RemoteAddr().String())
    reader := bufio.NewReader(conn)

	for {
		tempInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}
        input<- tempInput
	}
}

