package controller

import (
	"log"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/set2002satoshi/golang-blog/db"
	"github.com/set2002satoshi/golang-blog/model"
	"github.com/set2002satoshi/golang-blog/service"
)

func CustomerOne(c *gin.Context) {
	DbEngine := db.ConnectDB()
	num := c.Query("id")
	Customer := []model.Customer{}
	result := DbEngine.Where("id = ?", num).First(&Customer)
	if result.Error != nil {
		response := map[string]interface{}{
			"message": "error",
		}
		c.JSON(400, response)
		return
	}
	c.JSON(200, &Customer)	
}

func MyCustomerAll(c *gin.Context) {
	DbEngine := db.ConnectDB()
	Num, err := service.CheckUser(c)
	if err != nil {
		response := map[string]string{
			"message": "Unauthorized",
		}
		c.JSON(401, response)
		return 
	}
	fmt.Println(Num)
	userID, _ := strconv.Atoi(Num)
	customer := []model.Customer{}
	DbEngine.Where("id = ?", userID).Find(&customer)
	c.JSON(200, gin.H{
		"message": "ok",
		"data": customer,
	})
}


func CustomerAll(c *gin.Context){
	DbEngine := db.ConnectDB()
	Customer := []model.Customer{}
	DbEngine.Find(&Customer)
	c.JSON(200, gin.H{
		"user": Customer,
	})
}


func CustomerOneDelete(c *gin.Context) {
	DbEngine := db.ConnectDB()
	Customer := []model.Customer{}
	num := c.Query("id")
	result := DbEngine.Where("id = ?", num).Delete(&Customer)
	if result.Error != nil {
		log.Fatal("削除に失敗")
		c.JSON(400, gin.H{"err": result.Error})
	}
	c.JSON(200, gin.H{"message": "ok"})
}

func CustomerCreate(c *gin.Context) {
	DbEngine := db.ConnectDB()
	Num, err := service.CheckUser(c)
	if err != nil {
		response := map[string]string{
			"message": "Unauthorized",
		}
		c.JSON(401, response)
		return 
	}
	fmt.Println(Num)
	UserID, _ := strconv.Atoi(Num)
	fmt.Println(UserID)
	var userInfo model.CustomerInfo
	DbEngine.Where("id = ?", UserID).First(&userInfo)
	fmt.Println(userInfo)

	// username := "user"
	var CustomerForm model.CustomerForm
	err = c.Bind(&CustomerForm)
	if err != nil {
		response := map[string]string{
			"message": "not Bind",
		}
		c.JSON(400, response)
		return
	}
	Customer := model.Customer{
		CustomerInfoID: userInfo.ID,
		Name: CustomerForm.Name,
		// Thumbnail: CustomerForm.Thumbnail,
		Message: CustomerForm.Message,
	}
	result := DbEngine.Create(&Customer)
	if result.Error != nil {
		response := map[string]string{
			"message": "not create text",
		}
		c.JSON(400, response)
		return
	}
	fmt.Println(Customer.ID)
	ImgID, err := service.CustomerUploadImageS3(c, Num, Customer.ID)
	if err != nil {
		response := map[string]string{
			"message": "not create image",
		}
		c.JSON(500, response)
		return
	}
	result = DbEngine.Model(&Customer).Update("thumbnail", ImgID)
	if result.Error != nil {
		response := map[string]string{
			"message": "not add image",
		}
		c.JSON(500, response)
		return
	}
	response := map[string]interface{}{
		"message": "ok",
		"Customer": Customer,
	}
	c.JSON(201, response)	
}


