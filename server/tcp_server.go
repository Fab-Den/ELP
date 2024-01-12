package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	var data []byte

	buffer := make([]byte, 1024)
	for !strings.Contains(string(data), "end") {
		// Read data from the client

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		data = append(data, buffer[:n]...)
	}

	problem := string(data)
	// Print the data received from the client
	fmt.Println("Data received from client:", problem)
}

func acceptTcpConnections() {

	ln, err := net.Listen("tcp", ":8080")
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
