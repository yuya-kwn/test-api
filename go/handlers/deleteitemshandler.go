package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteItemHandler(c *gin.Context) {
	id := c.Param("name")

	itemIndex := -1
	for i := 0; i < len(items); i++ {
		if items[i].Name == id {
			itemIndex = i
		}
	}
	if itemIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "item not found"})
		return
	}
	items = append(items[:itemIndex], items[itemIndex+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Items has been deleted"})
}
