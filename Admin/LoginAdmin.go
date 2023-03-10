package admin

import (
	"fmt"

	"github.com/gofiber/fiber"
	"main.go/database"
	"main.go/model"
)

func LoginAdmin(c *fiber.Ctx) {
	Customers := model.Customers{}
	if err := c.BodyParser(&Customers); err != nil {
		panic(err)
	}
	db := database.ConnectDB()
	var DbEmails []string
	db.Table("customers").Select("email").Scan(&DbEmails)
	//for _,DbEmail:=range(DbEmails){
	IsValid := database.CompareUserNames(c, Customers.Email)
	if IsValid {
		var HashedPassword []model.Admin
		db.Raw("select password  from admins where email=?", Customers.Email).Scan(&HashedPassword)

		for _, Pass := range HashedPassword {
			err := database.ComparePasswords(c, Customers.Password, Pass.Password)
			if err != nil {
				fmt.Println(Customers.Password, " ", Pass.Password)
				c.Send("invalid password")
				break
			} else {
				fmt.Println(Customers.Password, " ", Pass.Password)
				c.Send("Logged in successfully")
				break
			}
		}
	} else {
		c.Send("The account doesn't exist")
	}
}
