package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

const defaultPort = "8080"

func main() {
	var conn net.Conn
	var err error

	for {
		fmt.Println("Server address : ")
		addressChoice := bufio.NewScanner(os.Stdin)

		addressChoice.Scan()

		// Connect to the server
		conn, err = net.Dial("tcp", addressChoice.Text()+":"+defaultPort)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			break
		}
	}

	if conn != nil {

		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				fmt.Println("Error when closing connection : ", err)
			}
		}(conn)

		var input string
		var choice string
		var str string

		for {
			//Ask the user if he has an input file ready
			fmt.Print("Do you have an input file? [y/n] ")
			scannerChoice := bufio.NewScanner(os.Stdin)
			if scannerChoice.Scan() {
				choice = scannerChoice.Text()
			}

			//if yes, ask for the path, retrieve it open the file and reads it
			if choice == "y" || choice == "Y" {
				fmt.Println("Enter file path: ")
				scannerFile := bufio.NewScanner(os.Stdin)
				if scannerFile.Scan() {
					filePath := scannerFile.Text()
					fmt.Println(filePath)
					input = fileReading(filePath)
				}

				// if not, ask the user to directly submit the data
			} else if choice == "N" || choice == "n" {
				fmt.Println("Enter your data: ")
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					line := scanner.Text()

					if line == "" {
						break
					}
					input += line + "\n"
				}

				if err := scanner.Err(); err != nil {
					fmt.Println("Error reading input:", err)
				}
			} else if choice == "quit" {
				fmt.Println("Closing client")
				return
			}
			// if the input, from the file or the console, is empty then return an error
			if input == "" || input == "error" {
				fmt.Println("Error in your data, unable to send")
				//else add an identification sequence at the end of the string
			} else {
				str = input + "end"
				fmt.Println(str)
			}

			// Send data to the server
			data := []byte(str)
			_, err = conn.Write(data)
			if err != nil {
				fmt.Println("Error:", err)
			}

			timeStart := time.Now()

			var rec []byte
			buffer := make([]byte, 1024)
			// Read data from the server
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err)
				return
			}
			rec = append(rec, buffer[:n]...)

			fmt.Println(string(rec))

			println("Execution time for :", time.Since(timeStart)/time.Millisecond, "ms")

		}
	}
}

func fileReading(path string) string {
	var str string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file : ", err)
		return "error"
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		str += line + "\n"
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file : ", err)
		}
	}(file)

	return str
}
