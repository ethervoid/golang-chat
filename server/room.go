package main

import (
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
			case userMessage := <-room.inputMessage:
				room.showMessage(userMessage)
			}
		}
	}()
}

func (room *Room) newUser(conn net.Conn) {
	user := &User{make(chan string), make(chan string)}
	room.users = append(room.users, user)

	go func() {
		room.outputMessage <- <-user.outputMessage
	}()

	user.listen(conn)
}

func (room *Room) showMessage(userMessage string) {
	//Iteramos por todos los usuarios y lanzamos el mensaje en el canal
	for _, user := range room.users {
		user.outputMessage <- userMessage
	}
}
