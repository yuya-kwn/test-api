package routes

import (
	"gin-test/go/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/items", handlers.NewItemHandler)
	router.GET("/items", handlers.ListItemsHandler)
	router.PUT("/items/:id", handlers.UpdateItemHandler)
	router.DELETE("/items/:id", handlers.DeleteItemHandler)

	router.LoadHTMLGlob("templates/index.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	return router
}
