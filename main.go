package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/ryanmello/crmgo/database"
	"github.com/ryanmello/crmgo/lead"
)

var db *gorm.DB

func setupRoutes(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

type Lead struct {

}

func initDatabase(){
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main(){
	app := fiber.New()

	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

	defer db.Close()
}