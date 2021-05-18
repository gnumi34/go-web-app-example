package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Name    string
	Age     int
	Hobbies []Hobby `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Hobby struct {
	gorm.Model
	Hobby     string
	ProfileID uint
}
