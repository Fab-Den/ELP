package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	var input string
	var choice string
	var str string

	for {
		//Ask the user if he has an input file ready
		fmt.Print("Do you have an input file? [y/n] ")
		scanner_choice := bufio.NewScanner(os.Stdin)
		if scanner_choice.Scan() {
			choice = scanner_choice.Text()
		}
		//if yes, ask for the path, retrieve it open the file and reads it
		if choice == "y" || choice == "Y" {
			fmt.Println("Enter file path: ")
			scanner_file := bufio.NewScanner(os.Stdin)
			if scanner_file.Scan() {
				file_path := scanner_file.Text()
				fmt.Println(file_path)
				input = file_reading(file_path)
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
	}
}

func file_reading(path string) string {
	var str string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
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

		}
	}(file)

	return str
}
