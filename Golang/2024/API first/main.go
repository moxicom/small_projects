package main

import (
	_ "api_first/docs" // This is required for Swagger

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Your API Title
// @version		1.0
// @description	Your API description
// @host			localhost:8080
// @BasePath		/

type Handler struct {
}

func main() {
	r := gin.Default()

	h := Handler{}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/users", h.GetUsers)

	r.Run(":8080")
}
