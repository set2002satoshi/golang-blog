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
		v1.POST("/user", controller.CustomerInfoCreate)
		v1.POST("/Certification", controller.Login)
	}
	v2 := router.Group("/api/app")
	{
		v2.GET("/blog", controller.BlogOne)
		v2.GET("/blog_all", controller.BlogAll)
		v2.POST("/blog_push", controller.BlogCreate)
		v2.DELETE("/blog", controller.BlogOneDelete)
		v2.DELETE("/blog_all", controller.BlogAllDelete)
	}
	v3 := router.Group("api/Customer")
	{
		v3.GET("/customer", controller.CustomerOne)
		v3.GET("/customer_all", controller.CustomerAll)
		v3.POST("/customer_push", controller.CustomerCreate)
		v3.DELETE("/customer", controller.CustomerOneDelete)
		// v3.DELETE("/customer_all", controller.Customer)

		
	}

	router.Run(":8080")
}



