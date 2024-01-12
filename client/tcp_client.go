package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send data to the server
	data := []byte("x y\n~x->-1:1\n~y->-1:1\n#x*x+y*y<1\n")
	_, err = conn.Write(data)
	_, err = conn.Write([]byte("end"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	time.Sleep(10000000000000)
}
