package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"github.com/set2002satoshi/golang-blog/db"
	"github.com/set2002satoshi/golang-blog/model"
	// "github.com/set2002satoshi/golang-blog/service"
)

func CustomerInfoCreate(c *gin.Context) {
	var CustomerForm model.CustomerInfoForm
	DbEngine := db.ConnectDB()
	err := c.BindJSON(&CustomerForm)
	if err != nil {
		response := map[string]string{
			"message": "not Bind",
		}
		c.JSON(404, response)
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(CustomerForm.Password), 14)
	if err != nil {
		response := map[string]string{
			"error": "bcrypt err",
		}
		c.JSON(203, response)
		return
	}
	Customer := model.CustomerInfo{
		Email:    CustomerForm.Email,
		Password: pass,
	}
	resp := DbEngine.Create(&Customer)
	if resp.Error != nil {
		c.JSON(203, resp.Error)
		return
	}
	c.JSON(200, gin.H{
		"Customer": Customer,
	})
}

func Login(c *gin.Context) {
	err := godotenv.Load("./config.env")
	if err != nil {
		fmt.Println("not read config.env")
		return
	}

	var LoginForm model.LoginForm
	CustomerInfoTable := model.CustomerInfo{}
	DbEngine := db.ConnectDB()
	err = c.BindJSON(&LoginForm)
	if err != nil {
		response := map[string]string{
			"message": "not Bind",
		}
		c.JSON(404, response)
		return
	}

	result := DbEngine.Where("email = ?", LoginForm.Email).First(&CustomerInfoTable)
	if result.Error != nil {
		response := map[string]string{
			"message": "email not found",
		}
		c.JSON(404, response)
		return
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(LoginForm.Password), 14)
	fmt.Println(pass)
	fmt.Println(LoginForm.Password)
	fmt.Println(CustomerInfoTable.Password)


	err = bcrypt.CompareHashAndPassword(CustomerInfoTable.Password, []byte(LoginForm.Password))
	if err != nil {
		c.JSON(204, err)
		fmt.Println(err)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(CustomerInfoTable.ID)),
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SecretKey")))
	if err != nil {
		response := map[string]string{
			"message": "email not found",
		}
		c.JSON(404, response)
		return
	}

	response := map[string]string{
		"message":   "success login",
		"ClientKey": token,
	}
	c.JSON(http.StatusOK, response)
}
