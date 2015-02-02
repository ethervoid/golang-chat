package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"time"
)

type User struct {
	id            string
	inputMessage  chan string
	outputMessage chan string
	disconnect    chan string
}

func NewUser() *User {
	userId := generateUserHash()
	user := &User{
		userId,
		make(chan string),
		make(chan string),
		make(chan string)}
	return user
}

func generateUserHash() string {
	hasher := md5.New()
	now := time.Now()
	return hex.EncodeToString(hasher.Sum([]byte(now.String())))
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
			if err == io.EOF {
				fmt.Println("User " + user.id + " disconnected")
				conn.Close()
				user.disconnect <- user.id
				break
			}
			fmt.Println("[Server Error] Reading incoming message:", err.Error())
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
