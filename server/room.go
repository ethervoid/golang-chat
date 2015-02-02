package main

import (
	"fmt"
	"net"
)

const DEFAULT_ROOM_NAME = "Lobby"

type Room struct {
	name          string
	joins         chan net.Conn
	inputMessage  chan string
	outputMessage chan string
	users         map[string]*User
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
	user := NewUser()
	room.users[user.id] = user

	go room.listenUsersMessages(user)
	go room.listenUsersDisconnections(user)

	user.listen(conn)
}

func (room *Room) listenUsersMessages(user *User) {
	for {
		room.outputMessage <- <-user.inputMessage
	}
}

func (room *Room) listenUsersDisconnections(user *User) {
	for {
		userToDelete := <-user.disconnect
		delete(room.users, userToDelete)
	}
}

func (room *Room) showMessage(userMessage string) {
	//Iteramos por todos los usuarios y lanzamos el mensaje en el canal
	for _, user := range room.users {
		user.outputMessage <- "[" + room.name + "] " + userMessage
	}
}
