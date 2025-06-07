package main

import (
	"fmt"
	"golang-course/lesson_12/internal/document_store"
	"golang-course/lesson_12/internal/general"
	"net"
	"strings"
)

var store *document_store.Store

func main() {
	store = document_store.NewStore()

	l, err := net.Listen("tcp", ":7777")
	defer l.Close()

	if err != nil {
		panic("Error listening: " + err.Error())
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			panic("Error accepting: " + err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connection from: " + conn.RemoteAddr().String())

	readBuffer := make([]byte, 1024)

	for {
		numberOfBytes, err := conn.Read(readBuffer)

		if err != nil {
			fmt.Println("Error reading: " + err.Error())
			return
		}

		messageFromClient := string(readBuffer[:numberOfBytes])

		clear(readBuffer)

		fmt.Println("Message from client:", messageFromClient)

		command, args, _ := strings.Cut(messageFromClient, " ")

		var response []byte

		switch command {
		case general.GetUserCommand:
			response, err = getUser(args)
		case general.AddUserCommand:
			response, err = addUser(args)
		case general.DeleteUserCommand:
			response, err = deleteUser(args)
		case general.GetUsers:
			response, err = getUsers()
		default:
			response = []byte("Command not recognized")
		}

		if err != nil {
			_, err = conn.Write([]byte(err.Error()))
		} else {
			_, err = conn.Write(response)
		}

		if err != nil {
			fmt.Println("Error writing: " + err.Error())
			return
		}
	}
}
