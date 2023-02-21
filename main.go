package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"main.go/database"
)

func HelloWorld(c *fiber.Ctx){
	c.Send("Hello Dancan")
}
func main() {
	app:=fiber.New()
    fmt.Println("connected")
	app.Post("/customer", database.CreateUserAccount)
	app.Post("/login",database.Login)
	app.Get("/",HelloWorld)
	app.Listen(":3000")
	
	database.ConnectDB()
}