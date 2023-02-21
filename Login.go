package database

import (
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
	"main.go/model"
)

func Login(c *fiber.Ctx) {
	Customers:=model.Customers{}
	if err:=c.BodyParser(&Customers);err!=nil{
		panic(err)
	}
	db:=ConnectDB()
	var DbEmails[] string
	db.Table("customers").Select("email").Scan(&DbEmails)
	for _,DbEmail:=range(DbEmails){
		UserExist:=CheckUserExistence(c,Customers.Email,DbEmail)
		if UserExist{
			var HashedPassword string
			db.Find("customers").Select("password").Where("email=?",Customers.Email).Scan(&HashedPassword)
			err:=CompareHashAndPassword(c,Customers.Password,HashedPassword)
			if err!=nil{
				c.Send("invalid password")
			}else{
				c.Send("Logged in successfully")
			}
		}else{
			c.Send("The account doesn't exist")
		}
	}
}
//comparing hash and password function
func CompareHashAndPassword(c *fiber.Ctx,password ,Hash string)error{
	err:=bcrypt.CompareHashAndPassword([]byte(password),[]byte(Hash))
	if err!=nil{
		return err
	}else{
		return nil
	}
}
//Check if the username exist in the database

