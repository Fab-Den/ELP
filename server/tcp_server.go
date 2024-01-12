package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	var data []byte

	buffer := make([]byte, 1024)
	for {
		// Read data from the client

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		data = append(data, buffer[:n]...)

		if n == 0 {
			break
		}

	}

	// Print the data received from the client
	fmt.Println("Data received from client:", string(buffer))
}

func acceptTcpConnections() {

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			fmt.Println("Error closing listening:", err)

		}
	}(ln)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)

	}

}

func main() {
	acceptTcpConnections()
}
