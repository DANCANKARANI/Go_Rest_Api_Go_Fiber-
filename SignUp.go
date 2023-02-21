package database

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main.go/model"
)

//Declaring Response User struct
//it will be used to give the response to the user after singing up
type Customer struct{
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Gender    string `json:"gender"`
	Email	  string `json:"email"`
}

//Signup endpoint
func CreateUserAccount(c *fiber.Ctx){
	Customers:=model.Customers{}
	if err:=c.BodyParser(&Customers);err!=nil{
		panic(err)
	}
	//calling the credintial verifier function
	db:=ConnectDB()
	email:=EmailVerifier(c, Customers.Email)
	if email=="Correct"{
		Names:=NamesVerifier(c,Customers.FirstName,Customers.LastName)
		if Names=="ValidName"{
			Password:=PasswordVerifier(c,Customers.Password)
			if Password=="true"{
				validPhone:=PhoneVerifier(c,Customers.Phone)
				if validPhone=="Valid"{
					var DbEmails[] Customer
					db.Table("customers").Select("email").Scan(&DbEmails)
					for _,dbEmail:=range(DbEmails){
						UserExist:=CheckUserExistence(c, Customers.Email,dbEmail.Email)
						if !UserExist{
							//c.Send(dbEmail.Email)
							//c.Send(Customers.Email)
							HashedPassword, err:=HashAndSalt(c,Customers.Password)
							if err!=nil{
								c.Send(err)
							}
							db.AutoMigrate(&Customers)
							db.Create(&model.Customers{
								Model:     gorm.Model{},
								FirstName: Customers.FirstName,
								LastName:  Customers.LastName,
								Gender:    Customers.Gender,
								Email:     Customers.Email,
								Phone: Customers.Phone,
								Password:  HashedPassword,
							})	
						//calling response function
						Response:=ResponseUser(c,Customers)
						c.JSON(Response)
						}else{
							c.Send("The user with this account already exist")
							//c.Send(dbEmail.Email)
						}
					}
				}else{
					c.Send("Invalid phone")
				}
			}else{
				c.Send("Invalid Password")
			}
		}else{
			c.Send("Invalid names")
		}
	}else{
		c.Send("Check your email")
	}
}
//response user function
func ResponseUser(c *fiber.Ctx,response model.Customers) Customer{
	Response:=Customer{
		FirstName: response.FirstName,
		LastName: response.LastName,
		Gender: response.Gender,
		Email: response.Email,
	}
	return Response
}
//Hashing the password
func HashAndSalt(c *fiber.Ctx, password string)(string,error){
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),14)
	if err!=nil{
		c.Send(string(err.Error()))
	}
	return string(bytes),nil
}

func EmailVerifier(c *fiber.Ctx, email string)string {
	IsValidEmail:=strings.Contains(email, "@")
	if IsValidEmail{
		fmt.Println("Correct")
		return "Correct"
	}
	return "error"
	
}
//verifies if names have more than 2 characters
func NamesVerifier(c *fiber.Ctx,FirstName,LastName string)string{
	IsValidName:=len(FirstName)>2 && len(LastName)>2
	if IsValidName{
		fmt.Println("Correct names")
		return "ValidName"
	}else{
		fmt.Println("Invalid names")
		return "InvalidName"
	}
}
//Verifies if tha password contains special characters
func PasswordVerifier(c *fiber.Ctx,password string)string{
	LowerCase:=regexp.MustCompile(`[a-z]`).MatchString
	UpperCase:=regexp.MustCompile(`[A-Z]`).MatchString
	SpecialChar:=regexp.MustCompile(`[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]`).MatchString
	if LowerCase(password) && UpperCase(password)&&SpecialChar(password){
		//c.Send(fiber.StatusNotAcceptable," Include lower case,uppercases and special characters")
		return "true"
	}else{
	
		return "false"
	}
}
//phone verifier function
func PhoneVerifier(c *fiber.Ctx,phone string)string{
	IsValidPhone:=regexp.MustCompile(`[0-9]`).MatchString
	if IsValidPhone(phone){
		return "Valid"
	}else{
		return "invalid"
	}
}
//checking if the user has already created an account
func CheckUserExistence(c *fiber.Ctx,email,DbEmail string)bool{
	if email == DbEmail{
		return true
	}else{
		return false
	}
}
/*
db:=ConnectDB()
				HashedPassword, err:=HashAndSalt(c,(Customers.Password))
				if err!=nil{
					c.Send(err)
				}
				
			db.AutoMigrate(&Customers)
			db.Create(&model.Customers{
				Model:     gorm.Model{},
				FirstName: Customers.FirstName,
				LastName:  Customers.LastName,
				Gender:    Customers.Gender,
				Email:     Customers.Email,
				Phone: Customers.Phone,
				Password:  HashedPassword,
		    })	
			//calling response function
			Response:=ResponseUser(c,Customers)
			c.JSON(Response)

*/