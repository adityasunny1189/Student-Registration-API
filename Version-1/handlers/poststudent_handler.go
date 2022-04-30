package handlers

import (
	"net/http"

	"github.com/adityasunny1189/Student-Registration-API/models"
	"github.com/adityasunny1189/Student-Registration-API/services/db"
	"github.com/gin-gonic/gin"
)

func AddStudent(c *gin.Context) {
	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		return
	}
	db.AddStudent(student)
	c.IndentedJSON(http.StatusCreated, student)
}
