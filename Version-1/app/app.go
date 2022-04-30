package app

import (
	"github.com/adityasunny1189/Student-Registration-API/handlers"
	"github.com/adityasunny1189/Student-Registration-API/services/db"
	redisdb "github.com/adityasunny1189/Student-Registration-API/services/redis-db"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	db.InitializeDB()
	redisdb.InitializeRedisDB()
	r := gin.Default()
	r.GET("/", handlers.HomeRoute)
	r.GET("/api/v1/students", handlers.GetStudentsData)
	r.GET("/api/v1/students/:usn", handlers.GetStudent)
	r.POST("/api/v1/student", handlers.AddStudent)
	r.Run()
}
