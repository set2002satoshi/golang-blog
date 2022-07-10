package controller

import (
	"fmt"

	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"github.com/set2002satoshi/golang-blog/db"
	"github.com/set2002satoshi/golang-blog/model"
	"github.com/set2002satoshi/golang-blog/service"
)

func CustomerInfoAll(c *gin.Context) {
	DbEngine := db.ConnectDB()
	CustomerInfo := []model.CustomerInfo{}
	// DbEngine.Find(&CustomerInfo)
	DbEngine.Preload("Blogs.Tags").Preload("Blogs").Preload("Customer").Find(&CustomerInfo)
	c.JSON(200, gin.H{
		"user": CustomerInfo,
	})
}

func MyCustomerInfoAll(c *gin.Context) {
	DbEngine := db.ConnectDB()
	CustomerInfo := []model.CustomerInfo{}
	Num, err := service.CheckUser(c)
	if err != nil {
		response := map[string]string{
			"message": "Unauthorized",
		}
		c.JSON(401, response)
		return 
	}
	userID, _ := strconv.Atoi(Num)
	DbEngine.Where("id = ?", userID).Preload("Blogs.Tags").Preload("Blogs").Preload("Customer").Find(&CustomerInfo)
	c.JSON(200, gin.H{
		"user": CustomerInfo,
	})
}




func CustomerInfoCreate(c *gin.Context) {
	var CustomerForm model.CustomerInfoForm
	DbEngine := db.ConnectDB()
	if c.Request.Method == "OPTION" {
		err := "Method is OPTION"
		c.JSON(200, err)
		fmt.Println(err)
		return
	}
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
	CustomerInfoData := model.CustomerInfo{
		Email:    CustomerForm.Email,
		Password: pass,
	}
	if resp := DbEngine.Create(&CustomerInfoData); resp.Error != nil {
		fmt.Println("errだよ")
		c.JSON(203, resp.Error)
		return
	}
	Customer := model.Customer{
		 CustomerInfoID: CustomerInfoData.ID,
		 Name:  CustomerForm.Name,
	}
	if resp := DbEngine.Create(&Customer); resp.Error != nil {
		c.JSON(203, resp.Error)
		var CI model.CustomerInfo
		DbEngine.Model(&CI).Delete(CustomerInfoData.ID)
		DbEngine.Where("ID = ?", CustomerInfoData.ID).First(&CI)
		c.JSON(203, gin.H{"確認": CI})
		return 
	}
	
	c.JSON(200, gin.H{
		"CustomerInfo": CustomerInfoData,

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
		"status": "ok",
		"message":   "success login",
		"ClientKey": token,
		
	}
	c.JSON(201, response)
}

func CustomerInfoAllDelete(c *gin.Context) {
	DbEngine := db.ConnectDB()
	CustomerInfoTable := []model.CustomerInfo{}

	DbEngine.Preload("Blogs").Preload("Customer").Find(&CustomerInfoTable)

	result := DbEngine.Unscoped().Delete(&CustomerInfoTable)
	if result.Error != nil {
		response := map[string]string{
			"message": "Not Delete CustomerInfo",
		}
		c.JSON(500, response) 
		return

	}
	response := map[string]string{
		"message": "successfully Delete CustomerInfo",
	}
	c.JSON(200, response)
}

