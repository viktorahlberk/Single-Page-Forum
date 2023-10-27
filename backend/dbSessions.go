package backend

import "time"

func isExpired(s Session) bool {
	if time.Now().Unix() > s.Expires {
		db := OpenDb()
		defer db.Close()
		query := `
			UPDATE sessions
				SET IsExpired=?
					WHERE UserName=?`
		_, err := db.Exec(query, 1, s.UserName)
		CheckErr(err)
		return true
	}
	return false
}
func prolongateSessionOnDb(username string) {
	db := OpenDb()
	defer db.Close()

	query := `
		UPDATE sessions
			SET Expires=?,
				IsExpired=?
					WHERE UserName=?`
	_, err := db.Exec(query, time.Now().Add(5*time.Minute).Unix(), 0, username)
	CheckErr(err)
}

/*func DeleteExpiredDbSessions() {
	sessions := getSessionsFromDb()
	for _, session := range sessions {
		if isExpired(session) {
			deleteSessionFromDb(session.Uuid)
		}
	}
}*/

func insertSessionToDb(s Session) {
	db := OpenDb()
	defer db.Close()

	_, err := db.Exec(`INSERT INTO sessions(
		UserName,
		Uuid,
		Expires,
		IsExpired
	) VALUES (?,?,?,?)`, s.UserName, s.Uuid, s.Expires, s.IsExpired)
	CheckErr(err)
}
func deleteSessionFromDb(Uuid string) {
	db := OpenDb()
	defer db.Close()

	_, err := db.Exec(`DELETE FROM sessions WHERE Uuid=?`, Uuid)
	CheckErr(err)
}
func getUserNameFromSessionDb(sessionUuid string) string {
	db := OpenDb()
	defer db.Close()
	var userName string
	err := db.QueryRow(`SELECT Username FROM sessions WHERE Uuid=?`, sessionUuid).Scan(&userName)
	CheckErr(err)
	return userName
}
func getUserIdFromUserNameFromDb(userName string) int {
	db := OpenDb()
	defer db.Close()
	var userId int
	err := db.QueryRow(`SELECT Id FROM users WHERE Nickname=?`, userName).Scan(&userId)
	CheckErr(err)
	return userId
}

//Get all sessions from database and return slice of them.
func getSessionsFromDb() []Session {
	db := OpenDb()
	defer db.Close()

	sessions := []Session{}
	query := `SELECT * FROM sessions`
	rows, err := db.Query(query)
	CheckErr(err)
	for rows.Next() {
		session := Session{}
		err = rows.Scan(&session.Id, &session.UserName, &session.Uuid, &session.Expires, &session.IsExpired)
		CheckErr(err)
		sessions = append(sessions, session)
	}

	return sessions
}
