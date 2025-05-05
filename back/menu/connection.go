package menu

import "database/sql"

func CreateConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}

	return db
}

func CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
