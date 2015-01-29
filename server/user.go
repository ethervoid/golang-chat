package main

import (
	"fmt"
	"net"
)

type User struct {
	inputMessage  chan string
	outputMessage chan string
}

func (user *User) listen(conn net.Conn) {
	go user.showMessage(conn)
	go user.sendMessage(conn)
}

func (user *User) sendMessage(conn net.Conn) {
	for {
		message := make([]byte, 1024)

		_, err := conn.Read(message)
		if err != nil {
			fmt.Println("[Error] Reading incoming message:", err.Error())
		}

		user.inputMessage <- string(message)
	}
}

func (user *User) showMessage(conn net.Conn) {
	fmt.Println("[Debug] Showing user message")
	for {
		message := <-user.outputMessage
		fmt.Println("Showing message: ", message)
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("[Server Error] Writing message:", err.Error())
		}
	}
}
