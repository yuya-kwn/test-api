package handlers

import (
	"gin-test/go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateItemHandler(c *gin.Context) {
	id := c.Param("id")
	var item models.Items

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erorr": err.Error()})
		return
	}
	targetItemIndex := -1
	for i := 0; i < len(items); i++ {
		if items[i].ID == id {
			targetItemIndex = i
		}
	}
	if targetItemIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item not found"})
		return
	}
	items[targetItemIndex] = item
	c.JSON(http.StatusOK, item)
}
