package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/set2002satoshi/golang-blog/controller"
)


func SetUpRouter(){
	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowCredentials: true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		},
	))
	
	router.Group("/api")
	{
		router.PUT("/", controller.Foge)
		// router.GET("/", controller.AllFind)
	}

	router.Run(":8080")
}