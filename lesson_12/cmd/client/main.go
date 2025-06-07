package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")

	if err != nil {
		fmt.Println("Error connecting: ", err)
	}

	readerBuffer := make([]byte, 1024)

	for {
		consoleReader := bufio.NewReader(os.Stdin)

		command, _, err := consoleReader.ReadLine()

		if err != nil {
			fmt.Println("Error reading from console: ", err)
		}

		_, err = conn.Write(command)

		if err != nil {
			fmt.Println("Error writing to connection: ", err)
		}

		numberOfSymbols, err := conn.Read(readerBuffer)

		if err != nil {
			fmt.Println("Error reading from connection: ", err)
		}

		fmt.Println("Message from server: ", string(readerBuffer[:numberOfSymbols]))

		clear(readerBuffer)
	}
}
