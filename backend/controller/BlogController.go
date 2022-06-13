package controller

import (
	"fmt"
	"log"
	"strconv"

	// "net//http"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"



	"github.com/set2002satoshi/golang-blog/db"
	"github.com/set2002satoshi/golang-blog/model"
	"github.com/set2002satoshi/golang-blog/service"
)

func BlogAll(c *gin.Context) {
	DbEngine := db.ConnectDB()
	b := []model.Blog{}
	DbEngine.Preload("Tags").Find(&b)
	c.JSON(200, gin.H{
		"user": b,
	})
}

func BlogOne(c *gin.Context) {
	DbEngine := db.ConnectDB()
	num := c.Query("id")
	b := []model.Blog{}
	result := DbEngine.Where("id = ?", num).First(&b)
	if result.Error != nil {
		response := map[string]interface{}{
			"message": "error",
		}
		c.JSON(400, response)
		return
	}
	c.JSON(200, &b)	
}

func BlogOneDelete(c *gin.Context) {
	DbEngine := db.ConnectDB()
	var Blog model.Blog
	num := c.Query("id")
	DbEngine.Where("id = ?", num).First(&Blog)

	err := service.BlogDeleteImageS3(c, string(Blog.BlogImage))
	if err != nil {
		response := map[string]interface{}{
			"message": "s3 error",
		}
		c.JSON(400, response)
		return
	}
	result := DbEngine.Delete(&Blog)
	if result.Error != nil {
		log.Fatal("削除に失敗")
		c.JSON(400, gin.H{"err": result.Error})
	}
	c.JSON(200, gin.H{"message": "ok"})
}

func BlogAllDelete(c *gin.Context){
	DbEngine := db.ConnectDB()
	BlogTable :=  []model.Blog{}
	DbEngine.Find(&BlogTable)

	result := DbEngine.Delete(&BlogTable)
	if result.Error != nil {
		response := map[string]interface{}{
			"message": "db error",
		}
		c.JSON(400, response)
		return
	}

	err := service.BlogAllDeleteImageS3(c)

	if err != nil {
		response := map[string]interface{}{
			"message": "500s3 error",
			"error": err,
		}
		c.JSON(500, response)
		return
	}
	c.JSON(200, gin.H{"status": "OK", "data": BlogTable})
}


func BlogCreate(c *gin.Context) {
	DbEngine := db.ConnectDB()
	var BlogForm model.BlogForm
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
	fmt.Println(userID)
	if err := c.Bind(&BlogForm);err != nil {
		response := map[string]string{
			"message": "not Bind",
		}
		c.JSON(400, response)
		return
	}
	var userInfo model.CustomerInfo
	DbEngine.Where("id = ?", userID).First(&userInfo)
	fmt.Println(userInfo)
	var tag model.Tag
	if result := DbEngine.Where("id = ?", BlogForm.Tag).Find(&tag); result.Error != nil {
		response := map[string]string{
			"message": "not category",
		}
		c.JSON(400, response)
		return
	}
	blog := model.Blog{
		CustomerInfoID: userInfo.ID,
		Title: BlogForm.Title,
		Subtitle: BlogForm.Subtitle,
		Content: BlogForm.Content,
		Tags: []model.Tag{tag},
	}
	if result := DbEngine.Create(&blog); result.Error != nil {
		response := map[string]string{
			"message": "not create text",
		}
		c.JSON(400, response)
		return
	}
	fmt.Println(blog)

	ImgID, err := service.BlogUploadImageS3(c, Num, blog.ID)
	if err != nil {
		response := map[string]string{
			"message": "not create image",
		}
		c.JSON(500, response)
		return
	}
	if result := DbEngine.Model(&blog).Select("blog_image").Updates(map[string]interface{}{"blog_image": ImgID,}); result.Error != nil {
		response := map[string]string{
			"message": "not add image",
		}
		c.JSON(500, response)
		return
	}

	if result := DbEngine.Model(&userInfo).Association("Blogs").Append(&blog);result.Error != nil {
		service.BlogDeleteImageS3(c, blog.BlogImage)
		DbEngine.Delete(&blog)
		fmt.Println("userとの紐付けに失敗したので全ての処理を削除")
		return
	}
	
	

	response := map[string]interface{}{
		"message": "ok",
		"blog": blog,
		"user": userInfo,
	}
	c.JSON(200, response)
}
