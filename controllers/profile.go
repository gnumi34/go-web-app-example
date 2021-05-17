package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"web-app-test/database"
	"web-app-test/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"gorm.io/gorm"
)

type ProfileForm struct {
	Id   int    `form:"-"`
	Name string `form:"name,text" valid:"required"`
	Age  int    `form:"age,text" valid:"required"`
}

type HobbyForm struct {
	Id    int    `form:"-"`
	Hobby string `form:"hobby,text" valid:"required"`
}

type ProfileController struct {
	beego.Controller
}

func (c *ProfileController) ListProfile() {
	// Fetch all user profiles
	var profiles []models.Profile
	database.Conn.Find(&profiles)

	// Render Form
	c.Data["Form"] = &ProfileForm{}

	// Render Template
	c.Layout = "base/base.html"
	c.TplName = "profile.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Footer"] = "base/footer.html"
	c.LayoutSections["Header"] = "base/header.html"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["Xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Results"] = profiles

	fmt.Println(profiles)

	if c.GetSession("success") == true {
		c.Data["SuccessMessage"] = "Data has been successfully submitted/deleted!"
		c.SetSession("success", false)
	}
}

func (c *ProfileController) AddProfile() {
	// Templates
	c.Layout = "base/base.html"
	c.TplName = "profile.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Footer"] = "base/footer.html"
	c.LayoutSections["Header"] = "base/header.html"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	form := ProfileForm{}
	if err := c.ParseForm(&form); err != nil {
		log.Println("[Err] Error on AddProfile : ", err.Error())
		c.Abort("500")
	}
	c.Data["Form"] = form

	// Fetch all user profiles
	var profiles []models.Profile
	database.Conn.Find(&profiles)
	c.Data["Results"] = profiles

	valid := validation.Validation{}
	valid.Required(form.Name, "Name").Message("is required")
	valid.Required(form.Age, "Age").Message("is required")

	errMessages := make(map[string]string)
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println("[FormEror] ", err.Key, " : ", err.Message)
			errMessages[err.Key] = err.Message
		}
		c.Data["ErrMessages"] = errMessages
		return
	}

	checkProfile := models.Profile{}
	if !errors.Is((database.Conn.Where("name = ?", form.Name).First(&checkProfile)).Error, gorm.ErrRecordNotFound) {
		log.Println("[Err] Error on AddProfile : ", "Duplicate Name")
		errMessages["Error on"] = " duplicate entry!"
		c.Data["ErrMessages"] = errMessages
		return
	}

	newProfile := models.Profile{
		Name: form.Name,
		Age:  form.Age,
	}
	database.Conn.Create(&newProfile)

	c.SetSession("success", true)
	c.Redirect("/profile", 302)
}

func (c *ProfileController) AddHobby() {
	idProfile := c.Ctx.Input.Param(":idProfile")
	profile := models.Profile{}

	err := database.Conn.First(&profile, idProfile).Error
	if err != nil {
		log.Println("[Err] DB error on fetching profile with id = ", idProfile)
	}

	// Render Form
	c.Data["Form"] = &HobbyForm{}

	// Render Template
	c.Layout = "base/base.html"
	c.TplName = "hobby.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Footer"] = "base/footer.html"
	c.LayoutSections["Header"] = "base/header.html"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["Xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Profile"] = profile
}

func (c *ProfileController) SubmitHobby() {
	idProfile := c.Ctx.Input.Param(":idProfile")
	profile := models.Profile{}

	err := database.Conn.First(&profile, idProfile).Error
	if err != nil {
		log.Println("[Err] DB error on fetching profile with id = ", idProfile)
	}

	// Templates
	c.Layout = "base/base.html"
	c.TplName = "hobby.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Footer"] = "base/footer.html"
	c.LayoutSections["Header"] = "base/header.html"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	form := HobbyForm{}
	if err := c.ParseForm(&form); err != nil {
		log.Println("[Err] Error on SubmitHobby : ", err.Error())
		c.Abort("500")
	}
	c.Data["Form"] = form

	valid := validation.Validation{}
	valid.Required(form.Hobby, "Hobby").Message("is required")

	errMessages := make(map[string]string)
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println("[FormEror] ", err.Key, " : ", err.Message)
			errMessages[err.Key] = err.Message
		}
		c.Data["ErrMessages"] = errMessages
		return
	}

	newHobby := models.Hobby{
		Hobby:     form.Hobby,
		ProfileID: profile.ID,
	}
	database.Conn.Create(&newHobby)

	c.SetSession("success", true)
	c.Redirect("/profile", 302)
}

func (c *ProfileController) DeleteProfile() {
	id := c.Ctx.Input.Param(":idProfile")
	err := database.Conn.Delete(&models.Profile{}, id).Error
	if err != nil {
		log.Println("[Err] Error on deleting profile with id = ", id)
		c.Abort("500")
	}

	c.SetSession("success", false)
	c.Redirect("/profile", 302)
}

func (c *ProfileController) DeleteHobby() {
	return
}
