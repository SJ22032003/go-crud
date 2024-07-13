package main

import (
	"os"

	"github.com/SJ22032003/go-crud/db"
	"github.com/SJ22032003/go-crud/routes"
	gin "github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default();

	db.ConnectToDb();

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
	port := os.Getenv("PORT")
	server.Run(":" + port)
}
