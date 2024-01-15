package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const numberWorkers = 12

func handleConnection(conn net.Conn, mainInputChannel chan<- mainInputContainer) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	for {

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
		println(conn.RemoteAddr().String())

		listVar, err := initializeVariables(problem)
		if err != nil {
			return
		}

		err = initVariableRange(listVar, problem)
		if err != nil {
			return
		}

		inequalities, err := initializeInequalities(problem)

		N := getNumberOfPoints(problem)
		if N == 0 {
			return
		}

		baseMaxN := 10000

		maxN := baseMaxN / inequalities.getProblemSize()

		println("MAXN : ", maxN)

		outChannel := make(chan subOutputContainer, 30)
		resultChannel := make(chan float64, 3)

		go receiveDataForRequest(outChannel, resultChannel, N)

		for N > 0 {
			tempN := 0

			if N < maxN {
				tempN = N
			} else {
				tempN = maxN
			}

			N = N - tempN

			mainInputChannel <- mainInputContainer{listVar: listVar, inequalities: inequalities, outputChannel: outChannel, N: tempN}

		}

		result := <-resultChannel
		close(resultChannel)
		close(outChannel)
		volume := getSpaceVolume(listVar)

		final := result * volume
		conn.Write([]byte(strconv.FormatFloat(final, 'f', -1, 64)))
	}

}

func acceptTcpConnections(mainInputChannel chan<- mainInputContainer, stopEvent <-chan bool) {

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
		select {
		case <-stopEvent:
			return

		default:
			// set a timeout to loop on the select if there is no connection attempt
			err := ln.(*net.TCPListener).SetDeadline(time.Now().Add(1 * time.Second))

			conn, err := ln.Accept()

			if err != nil {
				if !err.(net.Error).Timeout() {
					fmt.Println("Error accepting connection:", err)
				}
				continue
			}

			go handleConnection(conn, mainInputChannel)
		}

	}
}

func main() {
	stopServerChannel := make(chan bool, 1)

	mainInputChannel := make(chan mainInputContainer, 42)

	for i := 0; i < numberWorkers; i++ {
		go worker(mainInputChannel)
	}

	go acceptTcpConnections(mainInputChannel, stopServerChannel)

	select {}

	//close(stopServerChannel)

}
