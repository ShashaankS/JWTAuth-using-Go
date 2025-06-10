package main

import (
	"fmt"
	"os"

	"github.com/ShashaankS/GoEdit/backend/controllers"
	"github.com/ShashaankS/GoEdit/backend/initializers"
	"github.com/ShashaankS/GoEdit/backend/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
} 

func main() {
	fmt.Printf("Starting the server at %s", os.Getenv("PORT"))
	router := gin.Default()

	router.POST("/signup", controllers.Signup)
	router.POST("/signin", controllers.Signin)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.Run()
}