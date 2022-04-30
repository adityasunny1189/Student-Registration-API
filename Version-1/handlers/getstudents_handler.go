package handlers

import (
	"net/http"

	"github.com/adityasunny1189/Student-Registration-API/services/db"
	"github.com/gin-gonic/gin"
)

func GetStudentsData(c *gin.Context) {
	data := db.GetStudentsData()
	c.IndentedJSON(http.StatusOK, data)
}
