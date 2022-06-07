package controller

import (
	"fmt"
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
		c.JSON(404, response)
		return
	}
	c.JSON(200, &b)	
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
		c.JSON(404, response)
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
		c.JSON(404, response)
		return
	}
	fmt.Println(blog.ID)

	ImgID, err := service.ArticleUploadImageS3(c, username, blog.ID)

	if err != nil {
		response := map[string]string{
			"message": "not create image",
		}
		c.JSON(404, response)
		return
	}
	result = DbEngine.Model(&blog).Select("blog_image").Updates(map[string]interface{}{"blog_image": ImgID,})
	if result.Error != nil {
		response := map[string]string{
			"message": "not add image",
		}
		c.JSON(404, response)
		return
	}
	response := map[string]interface{}{
		"message": "ok",
		"blog": blog,
	}

	c.JSON(200, response)	
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
		c.JSON(404, response)
		return
	}
	err := service.ArticleAllDeleteImageS3(c)
	if err != nil {
		response := map[string]interface{}{
			"message": "s3 error",
			"error": err,
		}
		c.JSON(404, response)
		return
	}
	c.JSON(200, gin.H{"status": "OK", "data": BlogTable})
}


func S3testhandler(c *gin.Context) {
	err := service.ArticleAllDeleteImageS3(c)
	if err != nil {
		response := map[string]interface{}{
			"message": "s3 error",
			"error": err,
		}
		c.JSON(404, response)
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}

