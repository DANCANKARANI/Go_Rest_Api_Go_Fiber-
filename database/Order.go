package database

import (
	"github.com/gofiber/fiber"
	"main.go/model"
)

func MakeOrder(c *fiber.Ctx) {
	db := ConnectDB()
	Orders := model.Orders{}
	defer db.Close()
	if err := c.BodyParser(&Orders); err != nil {
		panic(err)
	} else {
		db.AutoMigrate(&Orders)
		db.Create(&Orders)
		c.Send(Orders)
	}
}
//view order
func ViewOrder(c *fiber.Ctx){
	db:=ConnectDB()
	defer db.Close()
	Orders:=model.Orders{}
	if err := c.BodyParser(&Orders); err != nil {
		panic(err)
	}else{
		Results:=model.Orders{}
		db.Raw("Select * From orders where id=?",1).Scan(&Results)
		c.Send(Results)
	}

}