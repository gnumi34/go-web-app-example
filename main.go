package main

import (
	"errors"
	"web-app-test/database"
	"web-app-test/models"
	_ "web-app-test/routers"

	"github.com/astaxie/beego"
	"gorm.io/gorm"
)

func main() {
	database.Connect()
	err := database.Conn.AutoMigrate(&models.Profile{}, &models.Hobby{})
	if err != nil {
		panic(err.Error())
	}

	// Seed new data if it is still blank
	var first models.Profile

	err = database.Conn.Where("name = ?", "Tripamungkas").First(&first).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		profile := models.Profile{
			Name: "Tripamungkas",
			Age:  25,
		}
		database.Conn.Create(&profile)
	}

	beego.Run()
}
