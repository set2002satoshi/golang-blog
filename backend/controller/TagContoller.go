package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/set2002satoshi/golang-blog/db"
	"github.com/set2002satoshi/golang-blog/model"
	// "github.com/set2002satoshi/golang-blog/service"
)

func TagAll(c *gin.Context) {
	DbEngine := db.ConnectDB()
	Category := []model.Tag{}
	DbEngine.Preload("Blogs").Find(&Category)
	c.JSON(200, gin.H{
		"category": Category,
	})
}
func TagAllDelete(c *gin.Context) {
	DbEngine := db.ConnectDB()
	TagTable := []model.Tag{}
	DbEngine.Find(&TagTable)
	result := DbEngine.Delete(&TagTable)
	if result.Error != nil {
		response := map[string]interface{}{
			"message": "db error",
		}
		c.JSON(400, response)
		return
	}
	c.JSON(200, gin.H{"status": "OK", "data": TagTable})
}

func AddTag(c *gin.Context) {
	DbEngine := db.ConnectDB()

	TagTable := model.Tag{}
	err := c.Bind(&TagTable)
	if err != nil {
		response := map[string]interface{}{
			"message": err,
		}
		c.JSON(400, response)
	}
	result := DbEngine.Create(&TagTable)
	if result.Error != nil {
		response := map[string]interface{}{
			"message": err,
		}
		c.JSON(500, response)
	}
	c.JSON(200, gin.H{
		"category": TagTable,
	})
}
