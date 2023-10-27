package backend

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func createSession(w http.ResponseWriter, userName string) string {
	sessionToken := uuid.NewV4().String()

	s := Session{
		UserName:  userName,
		Uuid:      sessionToken,
		Expires:   time.Now().Add(5 * time.Minute).Unix(),
		IsExpired: 0,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(12 * 30 * 24 * time.Hour), // expires on one year
	})
	insertSessionToDb(s)
	return sessionToken
}
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	//uuid := getUuidFromSessionCookie(r)
	deleteSessionFromDb(getUuidFromSessionCookie(r))

	c := http.Cookie{
		Name:   "session_token",
		MaxAge: -1}
	http.SetCookie(w, &c)

	w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	_, err := w.Write([]byte("Logged out"))
	CheckErr(err)

}
func getUuidFromSessionCookie(r *http.Request) string {
	sessionToken, err := r.Cookie("session_token")
	CheckErr(err)
	uuid := sessionToken.Value
	return uuid
}
