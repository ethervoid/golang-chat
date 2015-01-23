package main

import (
	"fmt"
	"net"
	"os"
)

func receiveMessages(conn net.Conn, channel chan string) {
	for {
		message := make([]byte, 1024)

		_, err := conn.Read(message)
		if err != nil {
			fmt.Println("[Error] Reading incoming message:", err.Error())
		}

		channel <- string(message)
	}
}

func sendMessages(conn net.Conn) {
	for {
		var input string
		fmt.Scanln(&input)

		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("[Error] Writing message:", err.Error())
		}
	}
}

func listen(conn net.Conn) {
	inputMessage := make(chan string)
	go receiveMessages(conn, inputMessage)
	sendMessages(conn)

	go func() {
		for {
			select {
			case message := <-inputMessage:
				fmt.Println(message)
			}
		}
	}()
}

func main() {
	address, err := net.ResolveTCPAddr("tcp", "localhost:9595")
	if err != nil {
		fmt.Println("[Error] Trying to resolve server address: ", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, address)
	if err != nil {
		fmt.Println("[Error] Trying to connect to the server: ", err.Error())
		os.Exit(1)
	}

	listen(conn)
}
