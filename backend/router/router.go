package router

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/golang-blog/controller"
)

func SetUpRouter() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"DELETE",
		},

		AllowHeaders: []string{
			"*",
		},

		AllowCredentials: true,

		MaxAge: 24 * time.Hour,
	}))
	v1 := router.Group("/api")
	{
		v1.GET("/customer-info_all", controller.CustomerInfoAll)
		v1.GET("/my-customer-info_all", controller.MyCustomerInfoAll)
		v1.POST("/create-user", controller.CustomerInfoCreate)
		v1.POST("/certification", controller.Login)
		v1.DELETE("/customer-info_all", controller.CustomerInfoAllDelete)
		v1.OPTIONS("/create-user", controller.CustomerInfoCreate)
	}
	v2 := router.Group("/api/category")
	{
		v2.GET("/tag", controller.TagAll)
		v2.PUT("/tag_push", controller.AddTag)
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
	v4 := router.Group("api/profile")
	{
		v4.GET("/customer", controller.CustomerOne)
		v4.GET("/my-customer", controller.MyCustomerAll)
		v4.GET("/customer_all", controller.CustomerAll)
		v4.POST("/customer_push", controller.CustomerCreate)
		v4.DELETE("/customer", controller.CustomerOneDelete)
		v4.DELETE("/customer_all", controller.CustomerAllDelete)
	}
	v5 := router.Group("api/main")
	{
		v5.GET("/home", controller.GetBlogsAll)
		v5.GET("/my-profile", controller.GetMyBlogsAdditionallyMyCustomer)
	}
	v6 := router.Group("api/")
	{
		v6.GET("/check", controller.UserCheckAdditionallyGetCustomer)
	}
	


	router.Run(":8080")
}



