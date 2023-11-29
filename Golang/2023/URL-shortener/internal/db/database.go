package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "dbname=sql_shorter user=postgres password=314159 host=localhost port=5432 sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

func (d *Database) InsertURL(real_url string) (int, error) {
	var id int

	stmt, err := d.db.Prepare("INSERT INTO urls(real_url) VALUES($1) RETURNING id")
	if err != nil {
		return id, err
	}

	log.Print("Statement created")

	defer stmt.Close()

	err = stmt.QueryRow(real_url).Scan(&id)
	if err != nil {
		return id, err
	}

	log.Print("Got id")

	return int(id), nil
}

func (d *Database) GetRealURL(id int) (string, error) {
	var real_url string
	db := d.db

	err := db.QueryRow("SELECT real_url FROM urls WHERE id = $1", id).Scan(&real_url)
	if err != nil {
		return "", err
	}
	return real_url, nil
}
