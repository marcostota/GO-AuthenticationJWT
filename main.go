package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcostota/gojwt/controllers"
	"github.com/marcostota/gojwt/initializers"
	"github.com/marcostota/gojwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Singin)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
