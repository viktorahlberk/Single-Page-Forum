package backend

import "fmt"

func insertCommentToDb(c Comment) {
	db := OpenDb()
	defer db.Close()

	query := `INSERT INTO comments (
		PostId,
		AuthorID,
		AuthorName,
		Body,
		Created
	)VALUES(?,?,?,?,?)`
	_, err := db.Exec(query, c.PostID, c.AuthorID, c.AuthorName, c.Body, c.Created)
	CheckErr(err)
	fmt.Printf("%s's comment was inserted to Db.\n", c.AuthorName)

}

func GetCommentsFromDbByPostId(postId int) []Comment {
	db := OpenDb()
	defer db.Close()
	var comments = []Comment{}

	query := `SELECT * FROM comments WHERE PostId=?`
	rows, err := db.Query(query, postId)
	CheckErr(err)
	for rows.Next() {
		comment := Comment{}
		err = rows.Scan(&comment.ID, &comment.PostID, &comment.AuthorID, &comment.AuthorName, &comment.Body, &comment.Created)
		CheckErr(err)
		comments = append(comments, comment)
	}
	return comments
}
