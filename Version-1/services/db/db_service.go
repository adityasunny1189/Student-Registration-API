package db

import (
	"fmt"
	"os"

	"github.com/adityasunny1189/Student-Registration-API/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:Adisunny123@tcp(127.0.0.1:3306)/sit?charset=utf8mb4&parseTime=True&loc=Local"
)

var db *gorm.DB

func InitializeDB() {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db.AutoMigrate(&models.Student{})
	fmt.Println("connected to database sit")
}

// Get all students data
func GetStudentsData() []models.Student {
	var data []models.Student
	fmt.Println("fetching data")
	db.Find(&data)
	return data
}

// Get a particular student data given usn
func GetStudentData(usn string) models.Student {
	var student models.Student
	db.First(&student, "usn = ?", usn)
	return student
}

// Add student to the database
func AddStudent(student models.Student) {
	db.Create(&student)
}
