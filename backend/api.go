package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {

	posts := GetAllPostsFromDb()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		fmt.Println("Couldnt encode JSON!")
	}

}

func GetPostAndComments(w http.ResponseWriter, r *http.Request) {
	type ErrorStruct struct {
		Code    int
		Message string
	}

	idString := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
		return
	}

	post, err := GetPostFromDb(idInt)

	if err != nil {
		if err == sql.ErrNoRows {

			// An example how we would write & handle an error to response
			errorMsg := ErrorStruct{400, "no rows"}
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(errorMsg.Code)
			w.Write([]byte(errorMsg.Message))

			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	if err := encoder.Encode(post); err != nil {
		w.WriteHeader(500)
		fmt.Println("Couldnt encode JSON!")
	}

}

func CreatePostComment(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Postid  int
		Comment string
	}

	var data requestBody
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}

	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(400)
			fmt.Println("Cant create comment because no cookie found")
		} else {
			w.WriteHeader(500)
			fmt.Println("Something unexpected happen while creating a post comment")
		}

		return
	}
	db := OpenDb()
	defer db.Close()

	userid := getUserIdFromUserNameFromDb(getUserNameFromSessionDb(sessionToken.Value))

	commentData := Comment{
		PostID:     data.Postid,
		AuthorID:   userid,
		AuthorName: getUserNameFromSessionDb(sessionToken.Value),
		Body:       data.Comment,
		Created:    time.Now(),
	}

	insertCommentToDb(commentData)

}

func IsUserAuthenticated(w http.ResponseWriter, r *http.Request) {
	type responseStruct struct {
		Authenticated bool   `json:"authenticated"`
		Nickname      string `json:"nickname"`
	}

	var response responseStruct

	cookie, err := r.Cookie("session_token")
	if err != nil {
		response.Authenticated = false
		log.Println(err)
	} else {
		response.Nickname = getUserNameFromSessionDb(cookie.Value)
		if response.Nickname == "" {
			response.Authenticated = false
		} else {
			prolongateSessionOnDb(response.Nickname)
			response.Authenticated = true
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
	}

}

func GetOtherUsers(w http.ResponseWriter, r *http.Request) {
	loggedInUserName := r.URL.Query().Get("user")
	var users1 []User
	if isUserHaveMessages(loggedInUserName) {
		users1 = GetUsersMessagedWith(loggedInUserName)
	} else {
		users1 = sortAlphabetically(getOtherUsersFromDb(loggedInUserName))
	}
	//users := sortAlphabetically(getOtherUsersFromDb(loggedInUserName))

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(users1); err != nil {
		log.Println(err)
	}

}

func GetLast10Messages(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	messages := GetLast10MessagesFromDb(urlValues.Get("sender"), urlValues.Get("reciever"), urlValues.Get("messages_loaded"))

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		log.Println(err)
	}

}
