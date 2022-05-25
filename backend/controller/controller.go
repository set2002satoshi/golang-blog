package controller

import (
	// "fmt"
	// "net/http"

	"github.com/gin-gonic/gin"

	// "github.com/set2002satoshi/golang-blog/db"
	// "github.com/set2002satoshi/golang-blog/model"
)

func Foge(c *gin.Context) {
	// var userD model.User
	// var formD model.UserBlog
	// DB := db.ConnectDB()
	// c.BindJSON(&formD)
	// fmt.Println(formD)
	// err := CreateDummy(formD.Name)
	// if err != nil {
	// 	response := map[string]string{
	// 		"msg": "1create err",
	// 	}
	// 	c.JSON(http.StatusAccepted, response)
	// 	return 
	// }
	// DB.First(&userD)
	// fmt.Println(userD.ID)
	// blog := model.Blog{
	// 	Title: formD.Title,
	// }

	// result := DB.Create(&blog).Association("User").Append(&userD)
	// if result != nil {
	// 	response := map[string]string{
	// 		"msg": "2create err",
	// 	}
	// 	c.JSON(http.StatusAccepted, response)
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"user": userD,
	// 	"content": blog,
	// })
	
}


// func CreateDummy(formName string) error {
// 	DB := db.ConnectDB()
// 	userDATA := model.User{
// 		Name: formName,
// 	}
// 	result := DB.Create(&userDATA)
// 	return result.Error
	
// }

// func AllFind(c *gin.Context){
// 	// var user model.User
// 	var blog model.Blog
// 	DB := db.ConnectDB()
// 	num := 7
// 	DB.Where("id = ?", num).Find(&blog)
// 	// DB.Model(&blog).Association("User").Find(&user)
// 	DB.Preload("User").Find(&blog)

// 	c.JSON(http.StatusOK,gin.H{
// 		"blog": blog,
// 	})
// }



