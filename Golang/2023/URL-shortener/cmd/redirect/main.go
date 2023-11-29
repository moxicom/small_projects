package main

import (
	"link_shorter/internal/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	r := gin.Default()

	r.GET("/:userURL", func(c *gin.Context) {
		userURL := c.Param("userURL")

		i, err := strconv.Atoi(userURL)
		if err != nil {
			log.Print("Atoi error: ", err)
			return
		}
		log.Print("Id: ", i)
		userURL, err = database.GetRealURL(i)
		log.Print("Original url: ", userURL)
		if err != nil {
			log.Print("GetRealURL error: ", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, userURL)
	})

	r.Run()
}
