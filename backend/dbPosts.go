package backend

import "fmt"

func insertPostToDb(post Post) {
	db := OpenDb()
	defer db.Close()

	query := `INSERT INTO posts (
		AuthorId,
		AuthorName,
		Title,
		Body,
		Created,
		Categories
	)VALUES(?,?,?,?,?,?)`
	_, err := db.Exec(query, post.AuthorID, post.AuthorName, post.Title, post.Body, post.Created, post.Categories)
	CheckErr(err)
}
func GetAllPostsFromDb() []Post {

	posts := []Post{}
	db := OpenDb()
	defer db.Close()
	query := `SELECT * FROM posts ORDER BY Id DESC`
	rows, err := db.Query(query)
	CheckErr(err)
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.AuthorID, &post.AuthorName, &post.Title, &post.Body, &post.Created, &post.Categories)
		CheckErr(err)
		posts = append(posts, post)
	}
	return posts

}

func GetPostFromDb(id int) (Post, error) {
	db := OpenDb()
	defer db.Close()

	var post Post

	if err := db.QueryRow("SELECT * FROM posts WHERE Id=?", id).
		Scan(&post.ID, &post.AuthorID, &post.AuthorName, &post.Title, &post.Body, &post.Created, &post.Categories); err != nil {
		fmt.Println("GetPostFromDb", err)

		return post, err
	}
	return post, nil
}
