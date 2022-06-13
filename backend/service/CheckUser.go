package service

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)


func CheckUser(c *gin.Context) (userID string, err error) {
	err = godotenv.Load("./config.env")
	if err != nil {
		fmt.Println("400 not read config.env")
		return "", err
	}
	// cookie, err := c.Cookie("clientKey")
	cookie  := c.Request.Header.Get("clientKey")

	// DbEngine := db.OpenDB()
	

	fmt.Println(cookie)

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SecretKey")), nil
	})

	if err != nil {
		fmt.Errorf("401")
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)


	return claims.Issuer, nil

}