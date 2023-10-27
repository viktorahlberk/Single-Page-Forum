package main

import (
	"fmt"
	"net/http"
	"rt-forum/backend"
)

func main() {
	backend.CreateTables()
	//fmt.Println(
	//backend.GetUsersMessagedWith("vic86")

	server := &http.Server{
		Addr:    ":8081",
		Handler: setRoutes(),
	}
	fmt.Printf("Server started at http://localhost" + server.Addr + "\n")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func setRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/favicon", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/", backend.IndexHandler)
	mux.HandleFunc("/login", backend.LoginHandler)
	mux.HandleFunc("/reg", backend.RegHandler)
	mux.HandleFunc("/logout", backend.DeleteSession)
	mux.HandleFunc("/newpost", backend.NewPostHandler)
	mux.HandleFunc("/getcomments", backend.GetCommentsHandler)
	mux.HandleFunc("/getallusers", backend.GetAllUsersHandler)
	mux.HandleFunc("/getonline", backend.GetOnlineHandler)

	//mux.HandleFunc("/ws", backend.HandleConnections)
	mux.HandleFunc("/ws", backend.WebsocketHandler)

	// API
	mux.HandleFunc("/allposts", backend.GetAllPosts)
	mux.HandleFunc("/post", backend.GetPostAndComments)
	mux.HandleFunc("/commentpost", backend.CreatePostComment)
	mux.HandleFunc("/userauth", backend.IsUserAuthenticated)
	mux.HandleFunc("/getotherusers", backend.GetOtherUsers)
	mux.HandleFunc("/getlast10messages", backend.GetLast10Messages)

	fs := http.FileServer(http.Dir("./frontend"))
	//go backend.HandleMessages()
	mux.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	return mux
}
