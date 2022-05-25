package controller

import (
	"net//http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	
	"github.com/set2002satoshi/golang-blog/db"
	"github.com/set2002satoshi/golang-blog/model"
)

func CustomerCreate(c *gin.Context) {
	var CustomerForm model.CustomerInfoForm
	DbEngine := db.ConnectDB()
	c.BindJSON(&CustomerForm)

	pass, err := bcrypt.GenerateFromPassword([]byte(CustomerForm.Password), 14)
	if err !=nil {
		response := map[string]string{
			"error": "bcrypt err",
		}
		c.JSON(203, response)
		return
	}
	Customer := model.CustomerInfo{
		Email: CustomerForm.Email,
		Password : pass,
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



