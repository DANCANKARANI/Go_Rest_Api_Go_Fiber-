package authenticatication

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gofiber/fiber"
)
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