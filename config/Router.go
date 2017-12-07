package config

import (
	"fmt"

	"github.com/WildanGarviandi/hafood-backend/controller"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			fmt.Println("options")
			c.Abort()
			return
		}

		c.Next()
	}
}

func SetRouter() *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", controller.CreateTodo)
		v1.GET("/", controller.FetchAllTodo)
		v1.GET("/:id", controller.FetchSingleTodo)
		v1.PUT("/:id", controller.UpdateTodo)
		v1.DELETE("/:id", controller.DeleteTodo)
	}

	return router
}
