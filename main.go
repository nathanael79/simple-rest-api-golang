package main

import (
	"simple-api/controllers"
	"simple-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	r := gin.Default()

	public := r.Group("/api")
	public.GET("/books", controllers.FindBooks)
	public.GET("/books/:id", controllers.FindBook)
	public.PATCH("/books/:id", controllers.UpdateBook)
	public.DELETE("/books/:id", controllers.DeleteBook)
	public.POST("/books", controllers.CreateBook)

	public.POST("/login", controllers.Login)

	public.POST("/register", controllers.Register)

	public.GET("/users", controllers.GetUsers)

	r.Run()
}
