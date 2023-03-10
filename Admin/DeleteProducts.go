package admin

import (
	"github.com/gofiber/fiber"
	"main.go/database"
)


func DeleteProducts(c *fiber.Ctx){
	db:=database.ConnectDB()
	defer db.Close()
}