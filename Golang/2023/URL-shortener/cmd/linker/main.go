package main

import (
	"fmt"
	"link_shorter/internal/db"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <URL>")
		return
	}

	url := "http://localhost:8080"

	database, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	var id int
	id, err = database.InsertURL(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your new URL is: %v/%v", url, id)
}
