package backend

import (
	"fmt"
	"strconv"
)

func InserMessageToDb(senderName, recieverName, content, created string) {

	db := OpenDb()
	defer db.Close()

	statement, err := db.Prepare(`INSERT INTO messages(
				senderName,
				recieverName,
				content,
				created
			) VALUES (?,?,?,?)`)

	CheckErr(err)
	defer statement.Close()

	statement.Exec(senderName, recieverName, content, created)
	fmt.Printf("%v message added to DB.\n", senderName)

}

func GetLast10MessagesFromDb(senderUsername, recieverUsername, messagesLoaded string) []ChatMessage {

	messages := []ChatMessage{}
	db := OpenDb()
	defer db.Close()

	// Select rows from messages table where senderName and recieverName match
	// query := `SELECT * FROM messages WHERE (senderName=? AND recieverName=?) OR (senderName=? AND recieverName=?) `
	intMessagesLoaded, err := strconv.Atoi(messagesLoaded)
	if err != nil {
		fmt.Println(err)
	}
	query := `
	SELECT * FROM (
		SELECT * FROM messages WHERE (senderName=? AND recieverName=?) OR (senderName=? AND recieverName=?)
		ORDER BY id DESC
		LIMIT 10 OFFSET ?	
	)
	ORDER BY id ASC
	`

	rows, err := db.Query(query, senderUsername, recieverUsername, recieverUsername, senderUsername, intMessagesLoaded)
	CheckErr(err)
	for rows.Next() {
		message := ChatMessage{}
		err = rows.Scan(&message.ID, &message.SenderName, &message.RecieverName, &message.Content, &message.Created)
		CheckErr(err)
		messages = append(messages, message)
	}

	// fmt.Println("Messages queried from database: ", messages)
	// fmt.Println("Messages loaded", intMessagesLoaded)
	return messages

}

//checks if ${username} have any messages in DB. Return boolean value.
func isUserHaveMessages(username string) bool {
	messages := []ChatMessage{}
	db := OpenDb()
	defer db.Close()

	rows, err := db.Query(`SELECT id FROM messages WHERE senderName = ? OR recieverName = ?`, username, username)
	if err != nil {
		fmt.Print("isUserHaveMessages golang function: ")
		CheckErr(err)
	}
	for rows.Next() {
		message := ChatMessage{}
		rows.Scan(&message.ID)
		messages = append(messages, message)
	}
	return len(messages) > 0
}

//Check database for messages existence, get userlist and sort them according last messages.
func GetUsersMessagedWith(username string) []User {
	userList := []string{}
	messages := []ChatMessage{}
	usersToReturn := []User{}

	db := OpenDb()
	defer db.Close()

	rows, err := db.Query(`SELECT senderName, recieverName FROM messages 
							WHERE senderName = ? OR recieverName = ?
								ORDER BY id DESC`, username, username)
	if err != nil {
		fmt.Print("isUserHaveMessages golang function: ")
		CheckErr(err)
	}
	for rows.Next() {
		message := ChatMessage{}
		rows.Scan(&message.SenderName, &message.RecieverName)
		messages = append(messages, message)
	}
	for _, v := range messages {
		if v.SenderName == username {
			userList = append(userList, v.RecieverName)
		} else {
			userList = append(userList, v.SenderName)
		}
	}
	userList = removeDuplicates(userList)
	fmt.Println("Userlist: ", userList)

	reversedUserList := userList
	//var reversedUserList []string
	//for i := len(userList); i > 0; i-- {
	//	reversedUserList = append(reversedUserList, userList[i-1])
	//}
	//fmt.Println(reversedUserList)

	allUsers := getOtherUsersFromDb(username)
	for _, v := range allUsers {
		fmt.Println("allusers: ", v.NickName)
	}
	//for _, user := range reversedUserList {
	for _, name := range allUsers {
		//if user == name.NickName {
		//	break
		//}
		reversedUserList = append(reversedUserList, name.NickName)

	}

	//}
	fmt.Println("reversed: ", reversedUserList)

	noDuplicatesUserList := (removeDuplicates(reversedUserList))
	fmt.Println("no duplicated: ", noDuplicatesUserList)

	for _, username := range noDuplicatesUserList {
		user := User{}
		row := db.QueryRow(`SELECT * FROM users WHERE Nickname = ?`, username)
		row.Scan(&user.Id, &user.Uuid, &user.NickName, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		usersToReturn = append(usersToReturn, user)
	}
	for _, v := range usersToReturn {
		fmt.Println("Usr: ", v.NickName)
	}
	fmt.Println("----------------")

	return usersToReturn

}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func removeDuplicates(strList []string) []string {
	list := []string{}
	for _, item := range strList {
		//fmt.Println(item)
		if !contains(list, item) {
			list = append(list, item)
		}
	}
	return list
}
