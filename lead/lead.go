package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/ryanmello/crmgo/database"
)

type Lead struct {
	gorm.Model
	Name string `json:"name"`
	Company string `json:"company"`
	Email string `json:"email`
	Phone int `json:"phone"`
}

func GetLeads(c *fiber.Ctx){
	db := database.GetDB()
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(){

}

func NewLead(){

}

func DeleteLead(){

}