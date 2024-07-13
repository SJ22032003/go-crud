package main

import (
	"github.com/SJ22032003/go-crud/routes"
	gin "github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default();

	v := server.Group("/");

	// TEST SERVER
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	// ROUTES
	routes.TodoRoutes(v);

	// SERVER LISTENING
	server.Run(":3000")
}
