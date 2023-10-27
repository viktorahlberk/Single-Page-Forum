package backend

import (
	"fmt"
)

func InsertUser(user User) {
	db := OpenDb()
	defer db.Close()

	statement, err := db.Prepare(`INSERT INTO users(
				Uuid,
				Nickname,
				Age,
				Gender,
				Firstname,
				Lastname,
				Email,
				Password
			) VALUES (?,?,?,?,?,?,?,?)`)

	CheckErr(err)
	defer statement.Close()
	statement.Exec(user.Uuid, user.NickName, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
	fmt.Printf("User %v registred & added to DB.\n", user.NickName)
}
func GetUserIdFromDb(loginOrEmail string) int {
	db := OpenDb()
	defer db.Close()
	var row int

	if isLoginIsEmail(loginOrEmail) {
		err := db.QueryRow("SELECT Id FROM users WHERE Email=?", loginOrEmail).Scan(&row)
		CheckErr(err)
	} else {
		err := db.QueryRow("SELECT Id FROM users WHERE Nickname=?", loginOrEmail).Scan(&row)
		CheckErr(err)
	}
	return row
}

func getAllUsersFromDb() []User {
	db := OpenDb()
	defer db.Close()

	users := []User{}
	query := `SELECT * FROM users`
	rows, err := db.Query(query)
	CheckErr(err)
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Uuid, &user.NickName, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		CheckErr(err)
		users = append(users, user)
	}

	return users
}

//get all users from Db except user in argument.
func getOtherUsersFromDb(nickname string) []User {
	db := OpenDb()
	defer db.Close()

	users := []User{}
	query := `SELECT * FROM users WHERE NOT Nickname=? ORDER BY Nickname ASC`
	rows, err := db.Query(query, nickname)
	CheckErr(err)
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Uuid, &user.NickName, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		CheckErr(err)
		users = append(users, user)
	}

	return users
}
