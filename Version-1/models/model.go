package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	USN    string `json:"usn"`
	Branch string `json:"branch"`
	Year   int    `json:"year"`
}
