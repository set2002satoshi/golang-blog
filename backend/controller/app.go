package controller

import (

	// "log"
	"net/http"
	"strconv"

	// "github.com/aws/aws-lambda-go/lambda/messages"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"

	"github.com/set2002satoshi/golang-blog/db"
	"github.com/set2002satoshi/golang-blog/model"
	"github.com/set2002satoshi/golang-blog/service"
)



func GetBlogsAll(c *gin.Context) {
	DbEngine := db.ConnectDB()
	var u bool
	var user string
	if Num, err := service.CheckUser(c); err != nil {
		u = false
		user = "not user"
	}else {
		var userInfo model.CustomerInfo
		userID, _ := strconv.Atoi(Num)
		DbEngine.Where("id = ?", userID).Preload("Customer").First(&userInfo)
		user = userInfo.Customer.Name
		u = true
	}
	b := []model.Blog{}
	DbEngine.Preload("Tags").Find(&b)
	c.JSON(http.StatusOK, gin.H{
		"blogs": b,
		"username": user,
		"certification": u,
		"status": "ok",
	})
	return
}

func UserCheckAdditionallyGetCustomer(c *gin.Context) {
	DbEngine := db.ConnectDB()
	var userInfo model.CustomerInfo
	Num, err := service.CheckUser(c)
	if err != nil {
		
		c.JSON(http.StatusOK, gin.H{"status": false})
		return
	}
	userID, _ := strconv.Atoi(Num)
	DbEngine.Where("id = ?", userID).Preload("Customer").First(&userInfo)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"userName": userInfo.Customer.Name,
	})
}



// func GetBlogsAll(c *gin.Context) {

// 	DbEngine := db.ConnectDB()
// 	b := []model.Blog{}
// 	DbEngine.Preload("Tags").Find(&b)
// 	c.JSON(200, gin.H{
// 		"user": b,
// 	})
// }














func GetMyBlogsAdditionallyMyCustomer(c *gin.Context) {
	DbEngine := db.ConnectDB()
	Num, err := service.CheckUser(c)

	var CI model.CustomerInfo
	if err != nil {
		response := map[string]string{
			"message": "Unauthorized",
		}
		c.JSON(401, response)
		return
	}

	userID, _ := strconv.Atoi(Num)
	// result := DbEngine.Where("id = ?", userID).Preload("Blogs").Preload("Customer").Find(&CI)
	result := DbEngine.Where("id = ?", userID).Preload("Blogs").Preload("Customer").Find(&CI)
	if result.Error != nil {
		response := map[string]string{
			"message": "データの取得に失敗",
		}
		c.JSON(404, response)
		return
	}
	

	// BlogData := model.Blog{
	// 	BlogID: CI.Blogs.ID,
	// 	BlogThumbnail: CI.Blogs.BlogImage,
	// 	Title: CI.Blogs.Title,
	// }

	// RespData := model.ProfilePage{
	// 	UserName: CI.Customer.Name,
	// 	UserThumbnail: CI.Customer.Thumbnail,
	// 	Message: CI.Customer.Message,
	// 	// Profile: CI.Blogs,
	// 	// BlogID: CI.Blogs.ID,
	// 	// BlogImage: CI.Blogs.BlogImage,
	// 	// Title: CI.Blogs.Title,		
	// }
	response := map[string]interface{}{
		"message": "ok",
		"Customer": CI.Customer,
		"Blogs": CI.Blogs,
	}
	c.JSON(200, response)	
}






