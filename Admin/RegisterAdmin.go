package admin

import (
	"github.com/gofiber/fiber"
	
	"main.go/database"
	"main.go/model"
)

func ReisterAdmin(c *fiber.Ctx) {
	db:= database.ConnectDB();
	Admin:=model.Admin{}
	if err:=c.BodyParser(&Admin);err!=nil{
		panic(err)
	}
	db.AutoMigrate(&Admin)
	Email:=database.EmailVerifier(c,Admin.Email)
	if Email=="Correct"{
		Password:=database.PasswordVerifier(c,Admin.Password)
		if Password=="true"{
			HashedPassword, err:=database.HashAndSalt(c,(Admin.Password))
			if err!=nil{
				c.Send(err)
			}
			db.Create(&model.Admin{
				Email: Admin.Email,
				Password: HashedPassword,
			})
			//calling Response Admin
			Response:=ResponseAdmin(c, Admin.Email)
			c.Send(Response)
		}else{
			c.Send("weak password")
		}
	}else{
		c.Send("Invalid email address")
	}	
}
type Admin struct{
	Email string `json:"email"`
}
func ResponseAdmin(c *fiber.Ctx,email string) Admin{
	return Admin{
		Email: email,
	}
}