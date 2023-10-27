package backend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Sender   string `json:"sender"`
	Reciever string `json:"reciever"`
	Created  string `json:"created"`
	Content  string `json:"content"`

	// UsersOnline []string `json:"usersOnline"`
}

type UsersStatus struct {
	UserJoined string `json:"userJoined"`
	UserLeft   string `json:"userLeft"`
}

var users = make(map[string]*websocket.Conn)

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Get user cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		log.Println(err)
	}

	userNickname := getUserNameFromSessionDb(cookie.Value)

	if _, ok := users[userNickname]; !ok {
		users[userNickname] = ws
		fmt.Println("Created new connection!")

	}

	var usersStatus UsersStatus
	usersStatus.UserJoined = userNickname
	for userName := range users {
		if userName == userNickname {
			continue
		}

		if err = users[userName].WriteJSON(&usersStatus); err != nil {
			log.Println("Couldnt send message", err.Error())
			continue
		}

	}

	// Connections
	fmt.Println("Ongoing connections: ", len(users))

	for {
		var msg Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("WebSocket closed", err.Error())
			delete(users, userNickname)
			var usersStatus UsersStatus
			usersStatus.UserLeft = userNickname

			for userName := range users {
				if err = users[userName].WriteJSON(&usersStatus); err != nil {
					log.Println("Cloud not send Message to ", err.Error())
					continue
				}

			}
			break

		}

		for userNickname1 := range users {
			if userNickname1 == msg.Reciever || userNickname1 == msg.Sender {
				if err = users[userNickname1].WriteJSON(msg); err != nil {
					log.Println("Cloud not send Message to ", err.Error())
					continue
				}
			}

		}

		InserMessageToDb(msg.Sender, msg.Reciever, msg.Content, msg.Created)

	}

	// if err := ws.Close(); err != nil {
	// 	log.Println("Websocket could not be closed", err.Error())
	// }

	ws.Close()
	fmt.Println("Websocket closed")
	// log.Println("Number of client still connected ...", len(users))
	// log.Println("Client connected:", ws.RemoteAddr().String())
	// log.Println("Number client connected ...", len(users))
}

/*var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}*/
