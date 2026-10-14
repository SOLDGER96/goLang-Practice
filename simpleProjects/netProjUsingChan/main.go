package main

import (
	"fmt"
	"net"
	"strings"
)

// For Comunication
type Hub struct {
	clients map[net.Conn]bool
	broadcast chan Message
	register chan net.Conn
	unregister chan net.Conn
}

type Message struct {
	sender net.Conn
	payload []byte
}

var PORT string = ":8080"


func (h *Hub) Run() {
	for {
		select {
		case clientConn := <-h.register:
			h.clients[clientConn] = true
		case clientConn := <-h.unregister:
			delete(h.clients, clientConn)
		case msg := <-h.broadcast:
			for clientConn := range h.clients {
				if clientConn != msg.sender {
					clientConn.Write(msg.payload)
				}
			}	
		}
	}
}

func main() {
	// var conn net.Conn
	// clients[conn] = true
	hub := newHub()
	go hub.Run()
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error Starting Server: ", err)
		return
	}

	defer listener.Close()
	fmt.Println("Chat server Listening on port ", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting Connection: ", err)
			continue
		}
		go handleConnection(conn, hub)
	}
}

func handleConnection(conn net.Conn, h *Hub) {
	defer conn.Close()
	
	// Asking for Username
	conn.Write([]byte("Enter your Username: "))

	//Temporary Buffer for accepting username
	nameBuffer := make([]byte, 1024)
	n, err := conn.Read(nameBuffer)
	if err != nil {
		return
	}
	// http.HandlerFunc("POST /api/abc",handleConnection())
	
	username := string(nameBuffer[:n])
	username = strings.TrimSpace(username)

	fmt.Printf("\n A new User: %v got connected!\n", username)
	// Logic for adding user to the map
	h.register <- conn

	for {
		textBuffer := make([]byte, 1024) // Data across internet travels in raw stream of 1s and 0s so we use here byte
		//conn.Write([]byte(">>"))

		n, err := conn.Read(textBuffer)
		if err != nil {
			fmt.Println("User Disconnected!", err)
			h.unregister <- conn
			return
		}
		message := textBuffer[:n]
		message = []byte(strings.TrimSpace(string(message)))
		messageToString := username + ": " + string(message) + "\n" + ">>"
		displayMessage := []byte(messageToString)
		
		msgEnvelope := Message{
			sender: conn,
			payload: displayMessage,
		}
		h.broadcast <- msgEnvelope
	}
}

func newHub() *Hub {
	return &Hub{
		clients: make(map[net.Conn]bool),
		broadcast: make(chan Message),
		register: make(chan net.Conn),
		unregister: make(chan net.Conn),
	}
}

