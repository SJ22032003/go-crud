package routes

import (
	handler "github.com/SJ22032003/go-crud/handler"
	gin "github.com/gin-gonic/gin"
)

func TodoRoutes(v *gin.RouterGroup) {

	todoHandler := handler.TodoHandler{}

	todoRoute := v.Group("todo")
	{
		todoRoute.GET("/", todoHandler.GetTodos)
	}

}
