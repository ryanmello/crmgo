package main

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/ryanmello/crmgo/database"
)

var db *gorm.DB

func setupRoutes(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase(){
	database.Connect()
	db = database.GetDB()
}

func main(){
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)

}