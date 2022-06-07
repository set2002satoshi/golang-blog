package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/golang-blog/controller"
)

func SetUpRouter() {
	router := gin.Default()
	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowCredentials: true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		},
	))
	v1 := router.Group("/api/user")
	{
		v1.POST("/user", controller.CustomerCreate)
		v1.POST("/Certification", controller.Login)
	}
	v2 := router.Group("/api/app")
	{
		v2.GET("/blog", controller.BlogOne)
		v2.GET("/blog_all", controller.BlogAll)
		v2.DELETE("/blog", controller.BlogOneDelete)
		v2.DELETE("/blog_all", controller.BlogAllDelete)
		v2.POST("/blog_push", controller.BlogCreate)
	}
	v3 := router.Group("api/test")
	{
		v3.GET("S3test", controller.S3testhandler)
	}
	router.Run(":8080")
}



