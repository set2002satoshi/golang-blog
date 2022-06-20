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
	var CustomerInfo model.CustomerInfo
	DbEngine.Where("id = ?", UserID).Preload("Customer").First(&CustomerInfo)
	fmt.Println(CustomerInfo.Customer)

	// username := "user"
	var CustomerMsgForm model.CustomerMsgForm
	err = c.Bind(&CustomerMsgForm)
	if err != nil {
		response := map[string]string{
			"message": "not Bind",
		}
		c.JSON(400, response)
		return
	}
	fmt.Println(CustomerMsgForm)
	// Customer := model.Customer{
	// 	// Thumbnail: CustomerForm.Thumbnail,
	// 	Message: CustomerForm.Message,
	// }
	// result := DbEngine.Create(&Customer)
	// if result.Error != nil {
	// 	response := map[string]string{
	// 		"message": "not create text",
	// 	}
	// 	c.JSON(400, response)
	// 	return
	// }
	fmt.Println(CustomerInfo.Customer.ID)
	ObjectKey, err := service.CustomerUploadImageS3(c, Num, CustomerInfo.Customer.ID)
	if err != nil {
		response := map[string]string{
			"message": "not create image",
		}
		c.JSON(500, response)
		return
	}

	fmt.Println("1")

	CustomerCatching := model.CustomerForm{
		Thumbnail: ObjectKey,
		Message: CustomerMsgForm.Message,
	}

	fmt.Println(CustomerCatching)

	// result := DbEngine.Model(&Customer).Update("thumbnail", ImgID)
	Customer := []model.Customer{}
	result := DbEngine.Model(&Customer).Where("id = ?", CustomerInfo.Customer.ID).Updates(&CustomerCatching)
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


