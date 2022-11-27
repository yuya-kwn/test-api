package routes

import (
	"gin-test/go/handlers"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/items", handlers.NewItemHandler)
	router.GET("/items", handlers.ListItemsHandler)
	router.PUT("/items/:id", handlers.UpdateItemHandler)
	router.DELETE("/items/:id", handlers.DeleteItemHandler)
	return router
}
