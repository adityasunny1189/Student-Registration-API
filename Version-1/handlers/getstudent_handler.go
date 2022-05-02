package handlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("getting data from redis")
		c.IndentedJSON(http.StatusOK, gin.H{"student": data})
	} else {
		student := db.GetStudentData(usn)
		if student.ID == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not present"})
			return
		}
		// Add data to redis
		fmt.Println("adding data to redis")
		val, err := json.Marshal(&student)
		if err != nil {
			fmt.Println(err)
			return
		}
		redisdb.SetData(usn, val)
		c.IndentedJSON(http.StatusOK, gin.H{"student": student})
	}
}
