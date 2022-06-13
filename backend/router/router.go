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
	v1 := router.Group("/api")
	{
		v1.GET("/customerInfo_all", controller.CustomerInfoAll)
		v1.POST("/User", controller.CustomerInfoCreate)
		v1.POST("/Certification", controller.Login)
	}
	v2 := router.Group("/api")
	{
		v2.GET("/tag", controller.TagAll)
		v2.PUT("/tag_add", controller.AddTag)
		v2.DELETE("/tag_all", controller.TagAllDelete)
	}
	v3 := router.Group("/api/app")
	{
		v3.GET("/blog", controller.BlogOne)
		v3.GET("/blog_all", controller.BlogAll)
		v3.POST("/blog_push", controller.BlogCreate)
		v3.DELETE("/blog", controller.BlogOneDelete)
		v3.DELETE("/blog_all", controller.BlogAllDelete)
	}
	V4 := router.Group("api/Customer")
	{
		V4.GET("/customer", controller.CustomerOne)
		V4.GET("/customer_all", controller.CustomerAll)
		V4.POST("/customer_push", controller.CustomerCreate)
		V4.DELETE("/customer", controller.CustomerOneDelete)
		// v3.DELETE("/customer_all", controller.Customer)

		

	}
	router.Run(":8080")
}



