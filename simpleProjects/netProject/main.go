package main

import (
	"fmt"
	"net"
	//"net/http"
	"strings"
	"sync"
)

var PORT string = ":8080"
var clients = make(map[net.Conn]bool)
var mutex sync.Mutex

func main() {
	// var conn net.Conn
	// clients[conn] = true
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
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
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	for {
		textBuffer := make([]byte, 1024) // Data across internet travels in raw stream of 1s and 0s so we use here byte
		//conn.Write([]byte(">>"))

		n, err := conn.Read(textBuffer)
		if err != nil {
			fmt.Println("User Disconnected!", err)
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			return
		}
		message := textBuffer[:n]
		message = []byte(strings.TrimSpace(string(message)))
		messageToString := username + ": " + string(message) + "\n" + ">>"
		displayMessage := []byte(messageToString)
		for clientConn := range clients {
			if clientConn != conn {
				clientConn.Write(displayMessage)
			}
		}
	}
}
