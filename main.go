package main

import (
	"fmt"
	"library-management1/database"
	"library-management1/handlers"
	"library-management1/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	fmt.Println("Hello!")
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You are on the home page"})
	})

	router.POST("/getdata", handlers.Test)
	router.DELETE("/del", handlers.Del)
	router.POST("/auth/signup", handlers.CreateUser)
	router.POST("/auth/login", handlers.Login)
	router.GET("/user/profile", middlewares.CheckAuth(), handlers.GetUserProfile)
	router.POST("/user/create-library", middlewares.CheckAuth(), handlers.CreateLibrary)
	router.POST("/user/assign-admin", middlewares.CheckAuth(), handlers.AssignAdmin)

	router.Run("localhost:8000")
}
