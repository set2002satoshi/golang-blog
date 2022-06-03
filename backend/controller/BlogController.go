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


// func BlogCreate(c *gin.Context) {
// 	username := "user"
// 	DbEngine := db.ConnectDB()
// 	var BlogForm model.BlogForm
// 	err := c.Bind(&BlogForm)
// 	if err != nil {
// 		response := map[string]string{
// 			"message": "not Bind",
// 		}
// 		c.JSON(404, response)
// 		return
// 	}
// 	blog := model.Blog{
// 		Title: BlogForm.Title,
// 		Subtitle: BlogForm.Subtitle,
// 		Content: BlogForm.Content,
// 	}
// 	result := DbEngine.Create(&blog)
// 	if result.Error != nil {
// 		response := map[string]string{
// 			"message": "not create text",
// 		}
// 		c.JSON(404, response)
// 		return
// 	}
// 	fmt.Println(blog.ID)
// 	ImgID, err := service.ArticleImageUploadS3(c, username, blog.ID)
// 	if err != nil {
// 		response := map[string]string{
// 			"message": "not create image",
// 		}
// 		c.JSON(404, response)
// 		return
// 	}
// 	result = DbEngine.Model(&blog).Update("blog_image", ImgID)
// 	if result.Error != nil {
// 		response := map[string]string{
// 			"message": "not add image",
// 		}
// 		c.JSON(404, response)
// 		return
// 	}
// 	response := map[string]interface{}{
// 		"message": "ok",
// 		"blog": blog,
// 	}
// 	c.JSON(200, response)
	
// }

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
	ImgID, err := service.ArticleImageUploadS3(c, username, blog.ID)
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

