package main

import (
	"github.com/gin-gonic/gin"
	"newton.go/controllers"
	"newton.go/middlewares"
	"newton.go/models"
)

func main() {
	models.ConnectDataBase()
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
