package db

import (
	"fmt"

	"github.com/adityasunny1189/Student-Registration-API/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:Adisunny123@tcp(127.0.0.1:3306)/sit?charset=utf8mb4&parseTime=True&loc=Local"
)

func InitializeDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.AutoMigrate(&models.Student{})
	return db
}
