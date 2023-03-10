package admin

import (
	"github.com/gofiber/fiber"
	"main.go/database"
	"main.go/model"
)

func ViewProducts(c *fiber.Ctx){
	db:=database.ConnectDB()
	defer db.Close()
	Products:=[]model.AvailableProducts{}
	db.Raw("select * from available_products").Scan(&Products)
	c.JSON(Products )

}

