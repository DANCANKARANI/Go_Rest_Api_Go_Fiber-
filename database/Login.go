package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt"
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
	//for _,DbEmail:=range(DbEmails){
		IsValid:=CompareUserNames(c,Customers.Email)
		if IsValid{
			var HashedPassword [] model.Customers
			db.Raw("select password  from customers where email=?",Customers.Email).Scan(&HashedPassword)
			
			for _,Pass:=range HashedPassword{
				err:=ComparePasswords(c,Customers.Password,Pass.Password)
				if err!=nil{
					fmt.Println(Customers.Password," ",Pass.Password)
					c.Send("invalid password")
					break
				}else{
					fmt.Println(Customers.Password," ",Pass.Password)
					c.Send("Logged in successfully")
					Token, _:=GenerateJWT(c,Pass.Password)
					c.Send(Token)
					break			
				}	
			}
		}else{
			c.Send("The account doesn't exist")
		}
	}

//Check if the username exist in the database
func CompareUserNames(c *fiber.Ctx , email string)bool{
	db:=ConnectDB()
	return !db.Where("email=?", email).First(&model.Customers{}).RecordNotFound()
}

func ComparePasswords(c *fiber.Ctx, plaintext string, hashedPassword string) error {
	// Convert the plaintext password to a byte slice
	passwordBytes := []byte(plaintext)
	hashedPasswordBytes := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
	if err != nil {
		return err
	}
	return nil
}

func GenerateJWT(c *fiber.Ctx, Password string)(string,error){
	secretkey := "secretkey"
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	
	
	claims["authorized"] = true
	claims["user_id"] = Password
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		c.Send(err.Error())
		return "", err
	}
	c.Send(tokenString)
	return tokenString, nil
}
func ExtractClaims(c *fiber.Ctx,tokenString string)(int,error){
	token,err:=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("Your security key"),nil
	})
	if err!=nil{
		return  0,err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(string); ok {
		  // Convert the ID claim from string to int
		  idInt, err := strconv.Atoi(id)
		  if err != nil {
			return 0, err
		  }
		  return idInt, nil
		
		}
	}
	return 0, err
}