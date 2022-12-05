package handlers

import (
	"encoding/json"
	"fmt"
	"gin-test/go/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var items []models.Items

func init() {
	items = make([]models.Items, 0)
	file, _ := ioutil.ReadFile("items.go")
	_ = json.Unmarshal([]byte(file), &items)
}

func NewItemHandler(c *gin.Context) {
	fmt.Println(c.Params)
	var json models.Items
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}
