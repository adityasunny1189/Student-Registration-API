package app

import (
	"fmt"

	"github.com/adityasunny1189/Student-Registration-API/services/db"
)

func StartApp() {
	db := db.InitializeDB()
	if db == nil {
		fmt.Println("connection failed")
		return
	}
	fmt.Println("connected to database sit")
}
