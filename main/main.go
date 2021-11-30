package main

import (
	"crud-ws/controllers"
	"crud-ws/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(context *gin.Context) {
		fmt.Fprintln(context.Writer, "<h1>Seja bem vindo!</h1>")
		fmt.Fprintln(context.Writer, "<h3><li>Rotas:</li></h3>")
		fmt.Fprintln(context.Writer, "<h4><li>GET:/books</li></h4>")
		fmt.Fprintln(context.Writer, "<h4><li>GET:/books/:id</li></h4>")
		fmt.Fprintln(context.Writer, "<h4><li>POST:/books</li></h4>")
		fmt.Fprintln(context.Writer, "<h4><li>PUT:/books/:id</li></h4>")
		fmt.Fprintln(context.Writer, "<h4><li>DELETE:/books/:id</li></h4>")
	})
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
