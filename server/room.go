package main

import (
	"fmt"
	"net"
)

const DEFAULT_ROOM_NAME = "lobby"

type Room struct {
	name          string
	joins         chan net.Conn
	inputMessage  chan string
	outputMessage chan string
	users         []*User
}

func (room *Room) open() {
	go func() {
		for {
			select {
			case userJoin := <-room.joins:
				room.newUser(userJoin)
			case userMessage := <-room.outputMessage:
				room.showMessage(userMessage)
			}
		}
	}()
}

func (room *Room) newUser(conn net.Conn) {
	fmt.Println("Conectado: ", conn)
	user := &User{make(chan string), make(chan string)}
	room.users = append(room.users, user)

	fmt.Println(room.users)

	go func() {
		room.outputMessage <- <-user.inputMessage
	}()

	user.listen(conn)
}

func (room *Room) showMessage(userMessage string) {
	//Iteramos por todos los usuarios y lanzamos el mensaje en el canal
	for _, user := range room.users {
		user.outputMessage <- userMessage
	}
}
