package handlers

import (
	"net/http"

	"github.com/adityasunny1189/Student-Registration-API/services/db"
	"github.com/gin-gonic/gin"
)

func GetStudent(c *gin.Context) {
	usn := c.Param("usn")
	student := db.GetStudentData(usn)
	c.IndentedJSON(http.StatusOK, gin.H{"student": student})
}
