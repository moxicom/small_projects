package main

import (
	_ "api_first/docs" // This is required for Swagger

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Your API Title
// @version 1.0
// @description Your API description
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Your API routes
	// e.g., r.GET("/users", getUsers)

	r.Run(":8080")
}
