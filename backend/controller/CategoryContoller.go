package controller

import (
	"github.com/gin-gonic/gin"


	"github.com/set2002satoshi/golang-blog/db"
	"github.com/set2002satoshi/golang-blog/model"
	// "github.com/set2002satoshi/golang-blog/service"
)

func CategoryAll(c *gin.Context){
	DbEngine := db.ConnectDB()
	Category := []model.Category{}
	DbEngine.Find(&Category)
	c.JSON(200, gin.H{
		"category": Category,
	})
}

func AddCategory(c *gin.Context) {
	DbEngine := db.ConnectDB()
	
	Category := model.Category{}
	err := c.Bind(&Category)
	if err != nil {
		response := map[string]interface{}{
			"message": err,
		}
		c.JSON(400, response)
	}
	result := DbEngine.Create(&Category)
	if result.Error != nil {
		response := map[string]interface{}{
			"message": err,
		}
		c.JSON(500, response)
	}
	c.JSON(200, gin.H{
		"category": Category,
	})
}