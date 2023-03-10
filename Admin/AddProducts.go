package admin

import (


	"github.com/gofiber/fiber"
	"main.go/database"
	"main.go/model"
)

func AddProducts(c *fiber.Ctx){
	Products:=model.AvailableProducts{}
	if err:=c.BodyParser(&Products);err!=nil{
		c.Send(err.Error())
	}else{
		db:=database.ConnectDB()
		defer db.Close()
		db.AutoMigrate(&Products)
		db.Create(&model.AvailableProducts{
			ProductName: Products.ProductName,
			Price: Products.Price,
			ProductType: Products.ProductType,
			Quantity: Products.Quantity,

		})
		c.Send("Procts added successfully")
	}
}