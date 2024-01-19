package main

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func createProblem(str string) (Problem, error) {
	var err error
	var problem Problem

	problem.listVars, err = initializeVariables(str)
	if err != nil {
		return problem, err
	}

	err = initVariableRange(problem.listVars, str)
	if err != nil {
		return problem, err
	}

	err = problem.initInequalities(str)
	if err != nil {
		return problem, err
	}

	err = problem.checkData()
	if err != nil {
		return problem, err
	}

	problem.numberOfPoints, err = getNumberOfPoints(str)
	if err != nil {
		return problem, err
	}

	return problem, nil
}

func handleConnection(conn net.Conn, mainInputChannel chan<- mainInputContainer) {

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	for {
		var err error
		err = nil

		// Read data input
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

		problemString := string(data)

		var problem Problem
		problem, err = createProblem(problemString)

		if err != nil {
			_, errWrite := conn.Write([]byte(err.Error()))
			if errWrite != nil {
				println("Error sending error to ", conn.RemoteAddr(), " : ", errWrite)
			}
		} else {

			baseMaxN := 10000

			maxN := baseMaxN / problem.getProblemSize()

			outChannel := make(chan subOutputContainer, 30)
			resultChannel := make(chan float64, 3)

			N := problem.numberOfPoints

			go receiveDataForRequest(outChannel, resultChannel, N)

			for N > 0 {
				tempN := 0

				if N < maxN {
					tempN = N
				} else {
					tempN = maxN
				}

				N = N - tempN

				mainInputChannel <- mainInputContainer{problem: problem, outputChannel: outChannel, N: tempN}

			}

			result := <-resultChannel
			close(resultChannel)
			close(outChannel)
			volume := getSpaceVolume(problem.listVars)

			final := result * volume
			_, errWrite := conn.Write([]byte(strconv.FormatFloat(final, 'f', -1, 64)))
			if errWrite != nil {
				println("Error sending result to ", conn.RemoteAddr(), " : ", errWrite)
			}
		}
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

	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(mainInputChannel)
	}

	go acceptTcpConnections(mainInputChannel, stopServerChannel)

	select {}

	//close(stopServerChannel)

}
