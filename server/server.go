package main

import (
	"fmt"
	"net"
	"os"
)

type Server struct {
	config Config
}

func (server *Server) startHandlingConnections(ln net.Listener, room Room) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error! ", err)
		}

		fmt.Println("Connection: ", conn)

		room.joins <- conn
	}
}

func (server *Server) start() {
	fmt.Println("Starting chat server...")
	config := server.config

	room := Room{
		DEFAULT_ROOM_NAME,
		make(chan net.Conn, 10),
		make(chan string),
		make(chan string),
		make([]*User, 0),
	}
	room.open()

	ln, err := net.Listen("tcp", config.hostName())
	if err != nil {
		fmt.Println("[Error]", err)
		os.Exit(1)
	}

	server.startHandlingConnections(ln, room)
}

func (server *Server) stop() {
	fmt.Println("Stopping chat server...")
	os.Exit(0)
}

func serverInit() Server {
	config := Config{DEFAULT_ADDRESS, DEFAULT_PORT, DEFAULT_MAX_CLIENTS}
	server := Server{config}

	return server
}

func main() {
	server := serverInit()
	server.start()
}
