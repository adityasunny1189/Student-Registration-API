package handlers

import (
	"net/http"

	"github.com/adityasunny1189/Student-Registration-API/services/db"
	redisdb "github.com/adityasunny1189/Student-Registration-API/services/redis-db"
	"github.com/gin-gonic/gin"
)

func GetStudent(c *gin.Context) {
	usn := c.Param("usn")
	// check for data in redis
	data, err := redisdb.GetData(usn)
	if err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"student": data})
	} else {
		student := db.GetStudentData(usn)
		// Add data to redis
		redisdb.SetData(usn, student)
		c.IndentedJSON(http.StatusOK, gin.H{"student": student})
	}
}
