package backend

func CreateTables() {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS users (
			"Id" INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
			"Uuid" TEXT NOT NULL,
			"Nickname" TEXT NOT NULL,
			"Age" INTEGER NOT NULL,
			"Gender" TEXT NOT NULL,
			"Firstname" TEXT NOT NULL,
			"Lastname" TEXT NOT NULL,
			"Email" TEXT NOT NULL,
			"Password" TEXT NOT NULL)`,

		`CREATE TABLE IF NOT EXISTS sessions (
			"Id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"UserName" TEXT NOT NULL,
			"Uuid" TEXT NOT NULL,
			"Expires" INTEGER,
			"IsExpired" INTEGER
		)`,

		`CREATE TABLE IF NOT EXISTS posts (
			"Id" INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
			"AuthorId" INTEGER NOT NULL,
			"AuthorName" TEXT NOT NULL, 
			"Title" TEXT NOT NULL,		
			"Body" TEXT NOT NULL,
			"Created" DATETIME,
			"Categories" TEXT
		)`,

		`CREATE TABLE IF NOT EXISTS comments (
			"ID" INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
			"PostId" INTEGER NOT NULL,
			"AuthorID" TEXT NOT NULL,
			"AuthorName" TEXT NOT NULL,
			"Body" TEXT NOT NULL,
			"Created" DATETIME
		)`,
	}

	db := OpenDb()
	defer db.Close()

	for _, table := range tables {
		statement, err := db.Prepare(table)
		CheckErr(err)
		defer statement.Close()
		statement.Exec()
	}
}
