package main

import (
	"database/sql"
	"flag"

	_ "github.com/mattn/go-sqlite3"
)

/*

Table definition:
CREATE TABLE urlinfo(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	original TEXT NOT NULL,
	short TEXT NOT NULL,
	created TEXT
);

*/

func main() {

	addPtr := flag.Bool("add", false, "a bool")
	trnPtr := flag.Bool("translate", false, "a bool")
	delPtr := flag.Bool("delete db", false, "a bool")

	db, err := sql.Open("sqlite3", "./url-shortener.db")
	if err != nil {
		panic(err)
	}

	res, err := insert_into_db(db, "https://www.google.com", "non.zero/UNIQUE_SHORT_HASH")
	if err != nil {
		panic(err)
	}

	/*
		affect, err = res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)
	*/
	db.Close()

}

func delete_from_db(db *sql.DB, where string) (sql.Result, error) {
	stmt, err := db.Prepare("DELETE FROM urlinfo WHERE ?")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(where)
	if err != nil {
		return nil, err
	}

	return res, err
}

func insert_into_db(db *sql.DB, url string, short string) (sql.Result, error) {
	stmt, err := db.Prepare("INSERT INTO urlinfo(original, short, created) values(?,?, datetime('now'))")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(url, short)
	if err != nil {
		return nil, err
	}

	return res, err
}

func select_from_db(db *sql.DB, where string) (sql.Result, error) {
	stmt, err := db.Prepare("SELECT * FROM urlinfo WHERE ?")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(where)
	if err != nil {
		return nil, err
	}

	return res, err
}
