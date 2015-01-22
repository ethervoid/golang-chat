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
	go user.receiveMessage(conn)
	go user.sendMessage(conn)
}

func (user *User) sendMessage(conn net.Conn) {
	for {
		message := make([]byte, 1024)

		_, err := conn.Read(message)
		if err != nil {
			fmt.Println("[Error] Reading incoming message:", err.Error())
		}

		user.outputMessage <- string(message)
	}
}

func (user *User) receiveMessage(conn net.Conn) {
	for {
		message := <-user.outputMessage
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("[Error] Writing message:", err.Error())
		}
	}
}
