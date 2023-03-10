package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"main.go/Admin"
	"main.go/database"
)

func HelloWorld(c *fiber.Ctx){
	c.Send("Hello Dancan")
}
func main() {
	app:=fiber.New()
	//customer
    fmt.Println("connected")
	app.Post("/customer", database.CreateUserAccount)
	app.Post("/login",database.Login)
	app.Post("/order",database.MakeOrder)
	app.Get("/get/order",database.ViewOrder)
	//Admin
	app.Post("/admin",admin.ReisterAdmin)
	app.Post("/admin/login",admin.LoginAdmin)
	//products
	app.Post("/add/product",admin.AddProducts)
	app.Get("/get/product",admin.ViewProducts)
	app.Get("/",HelloWorld)
	app.Listen(":3000")
	
	database.ConnectDB()
}