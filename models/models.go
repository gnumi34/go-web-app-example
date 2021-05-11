package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Name    string
	Age     int
	Hobbies []Hobby
}

type Hobby struct {
	gorm.Model
	Hobby     string
	ProfileID uint
}
