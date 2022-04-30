package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
