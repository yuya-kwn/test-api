package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListItemsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}
