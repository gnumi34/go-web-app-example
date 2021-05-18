package main

import (
	"web-app-test/database"
	"web-app-test/models"
	_ "web-app-test/routers"

	"github.com/astaxie/beego"
)

func main() {
	database.Connect()
	err := database.Conn.AutoMigrate(&models.Profile{}, &models.Hobby{})
	if err != nil {
		panic(err.Error())
	}

	beego.Run()
}
