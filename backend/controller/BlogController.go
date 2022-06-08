package controller

import (
	"fmt"
	"log"
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
	DbEngine.Find(&b)
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
	b := []model.Blog{}
	num := c.Query("id")
	result := DbEngine.Where("id = ?", num).Delete(&b)
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
	username := "user"
	DbEngine := db.ConnectDB()
	var BlogForm model.BlogForm
	err := c.Bind(&BlogForm)
	if err != nil {
		response := map[string]string{
			"message": "not Bind",
		}
		c.JSON(400, response)
		return
	}
	blog := model.Blog{
		Title: BlogForm.Title,
		Subtitle: BlogForm.Subtitle,
		Content: BlogForm.Content,
	}
	result := DbEngine.Create(&blog)
	if result.Error != nil {
		response := map[string]string{
			"message": "not create text",
		}
		c.JSON(400, response)
		return
	}
	fmt.Println(blog.ID)
	ImgID, err := service.BlogUploadImageS3(c, username, blog.ID)
	if err != nil {
		response := map[string]string{
			"message": "not create image",
		}
		c.JSON(500, response)
		return
	}
	result = DbEngine.Model(&blog).Update("blog_image", ImgID)
	if result.Error != nil {
		response := map[string]string{
			"message": "not add image",
		}
		c.JSON(500, response)
		return
	}
	response := map[string]interface{}{
		"message": "ok",
		"blog": blog,
	}
	c.JSON(200, response)	
}


// func S3testhandler(c *gin.Context) {
// 	err := service.ArticleAllDeleteImageS3(c)
// 	if err != nil {
// 		response := map[string]interface{}{
// 			"message": "s3 error",
// 			"error": err,
// 		}
// 		c.JSON(500, response)
// 		return
// 	}
// 	c.JSON(200, gin.H{"status": "OK"})
// }