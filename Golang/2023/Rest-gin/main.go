package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getTaskHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	fmt.Printf("Params: %#v\n", c.Params)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "BLALBALBA"+strconv.Itoa(id))
}

func createTaskHandler(c *gin.Context) {
	type RequestTask struct {
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
		Due  time.Time `json:"due"`
	}
	fmt.Println("current time: ", time.Now())
	var rt RequestTask
	if err := c.ShouldBindJSON(&rt); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// send new task's id
	c.JSON(http.StatusOK, rt)
}

func main() {
	router := gin.Default()

	router.POST("/tasks/", createTaskHandler)
	router.GET("/tasks/", func(ctx *gin.Context) {})
	router.DELETE("/tasks/", func(ctx *gin.Context) {})
	router.GET("/tasks/:id", getTaskHandler)
	router.DELETE("/tasks/:id", func(ctx *gin.Context) {})
	router.GET("/tags/:tag", func(ctx *gin.Context) {})
	router.GET("/due/:year/:month/:day", func(ctx *gin.Context) {})

	router.Run()
}
