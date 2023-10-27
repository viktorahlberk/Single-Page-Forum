package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetOnlineHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sessions := getSessionsFromDb()
	onlineSessions := []Session{}
	for _, session := range sessions {
		if !isExpired(session) {
			onlineSessions = append(onlineSessions, session)
		}
	}
	response, err := json.Marshal(onlineSessions)
	CheckErr(err)
	w.Write(response)
}

func getHashFromPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func isPasswordEqualtoDb(loginOrEmail string, password string) bool {
	userId := GetUserIdFromDb(loginOrEmail)
	var hash string
	db := OpenDb()
	defer db.Close()
	err := db.QueryRow("SELECT Password FROM users WHERE Id=?", userId).Scan(&hash)
	if err != nil {
		log.Println(err)
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil {
		fmt.Println("Password is Correct.")
		return true
	} else {
		fmt.Println("Password is INCORRECT.")
		return false
	}
}

func isLoginIsEmail(login string) bool {
	return strings.Contains(login, "@")
}
func isEmailExistInDb(email string) bool {
	db := OpenDb()
	defer db.Close()
	var row int
	err := db.QueryRow("SELECT Id FROM users WHERE email=?", email).Scan(&row)
	if err != nil {
		fmt.Println(err)
	}
	if row != 0 {
		return true
	}
	return false
}
func isNicknameExistInDb(nickname string) bool {
	db := OpenDb()
	defer db.Close()
	var row int
	err := db.QueryRow("SELECT Id FROM users WHERE Nickname=?", nickname).Scan(&row)
	if err != nil {
		fmt.Println(err)
	}
	if row != 0 {
		return true
	}
	return false
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/index.html")
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	credentials := LoginCredentials{}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &credentials)
	fmt.Println(credentials.Login)

	if isLoginIsEmail(credentials.Login) {
		if isEmailExistInDb(credentials.Login) {
			if isPasswordEqualtoDb(credentials.Login, credentials.Password) {
				createSession(w, credentials.Login)
				isValidated(w, true)
				return
			}
			fmt.Println("This e-mail Exist in DB.") //TODO
		} else {
			isValidated(w, false)

			fmt.Println("This e-mail DONT EXIST in DB.") //TODO
			return
		}
	} else {
		if isNicknameExistInDb(credentials.Login) {
			if isPasswordEqualtoDb(credentials.Login, credentials.Password) {
				createSession(w, credentials.Login)
				isValidated(w, true)
				return
			}
			fmt.Println("This nickname EXIST in DB.") //TODO
		} else {
			isValidated(w, false)
			fmt.Println("This nickname DONT EXIST in DB.") //TODO
			return
		}
	}

}

func RegHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("reghandler")
	if r.Method == "POST" {

		//d, err := ioutil.ReadAll(r.Body)
		//CheckErr(err)
		//fmt.Println(string(d))
		err := r.ParseMultipartForm(64)
		CheckErr(err)

		newUser := User{
			Uuid:      uuid.NewV4().String(),
			NickName:  r.FormValue("nickname"),
			Age:       r.FormValue("age"),
			Gender:    r.FormValue("gender"),
			FirstName: r.FormValue("fname"),
			LastName:  r.FormValue("lname"),
			Email:     r.FormValue("email"),
			Password:  getHashFromPassword([]byte(r.FormValue("passw"))),
		}
		//fmt.Println(newUser)

		InsertUser(newUser)

		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		a := "Registered"

		//ans, _ := json.Marshal(a)
		w.Write([]byte(a))

	}
}
func isValidated(w http.ResponseWriter, isValidated bool) {
	w.Header().Set("Content-Type", "application/json")
	a := Answer{}
	a.IsAuthorised = isValidated
	ans, _ := json.Marshal(a)
	w.Write(ans)

}
func ChangeCategoryHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	CheckErr(err)
	fmt.Println("YO", string(b))
	w.Header().Set("text/plain", "charset=UTF-8")
	w.Write(b)
}
func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("newPostHandler func")

	if r.Method == "POST" {
		c, err := r.Cookie("session_token")
		if err == http.ErrNoCookie {
			w.WriteHeader(400)
			w.Write([]byte("Session expired!"))
			return
		}
		CheckErr(err) // return if user is not logged in ?
		err = r.ParseMultipartForm(64)
		CheckErr(err)

		newPost := Post{
			AuthorID:   getUserIdFromUserNameFromDb(getUserNameFromSessionDb(c.Value)),
			AuthorName: getUserNameFromSessionDb(c.Value),
			Title:      r.FormValue("title"),
			Body:       r.FormValue("body"),
			Created:    time.Now(),
			Categories: r.FormValue("cat1") + r.FormValue("cat2") + r.FormValue("cat3") + r.FormValue("cat4") + r.FormValue("cat5") + r.FormValue("cat6"),
		}
		fmt.Println(newPost)

		insertPostToDb(newPost)

		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		a := "Submitted"

		w.Write([]byte(a))
	}
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := getAllUsersFromDb()
	response, err := json.Marshal(users)
	CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	//fmt.Println(idString)
	idInt, err := strconv.Atoi(idString)
	CheckErr(err)

	comments := GetCommentsFromDbByPostId(idInt)
	//fmt.Println(comments)
	response, err := json.Marshal(comments)
	CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
