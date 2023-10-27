package backend

import "time"

type Session struct {
	Id        int
	UserName  string
	Uuid      string
	Expires   int64
	IsExpired int
}

/*type DbSession struct {
	Id      int
	UserId  int
	Expires time.Time
}*/

type LoginCredentials struct {
	Login    string
	Password string
}
type Answer struct {
	IsAuthorised bool
}
type User struct {
	Id        int
	Uuid      string
	NickName  string
	Age       string
	Gender    string
	FirstName string
	LastName  string
	Email     string
	Password  string
}
type Post struct {
	ID         int
	AuthorID   int
	AuthorName string
	Title      string
	Body       string
	Created    time.Time
	Categories string
}
type Category struct {
	PostID int
	Name   string
}

type Comment struct {
	ID         int       `json:"id"`
	PostID     int       `json:"postId"`
	AuthorID   int       `json:"authorId"`
	AuthorName string    `json:"authorName"`
	Body       string    `json:"body"`
	Created    time.Time `json:"created"`
}

type ChatMessage struct {
	ID           int    `json:"id"`
	SenderName   string `json:"sender"`
	RecieverName string `json:"reciever"`
	Content      string `json:"content"`
	Created      string `json:"created"`
}
