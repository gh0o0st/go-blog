package main

import (
	"database/sql"
	"time"
)

//Post represent a blog post
type Post struct {
	Header     string
	Category   string
	Tags       []string
	Body       string
	CreateTime time.Time
}

//Blog represent the entire blog
type Blog struct {
	Title string
	Posts []Post
}

func open() (*sql.DB, error) {
	return sql.Open("sqlite3", dbName)
}

func insert(header, body string) error {
	db, err := open()
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO Post(header, body) values(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(header, body)
	if err != nil {
		return err
	}
	return nil
}

func selectAll() []Post {

}
