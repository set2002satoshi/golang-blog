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
		v2.GET("/all", controller.BlogAll)
		v2.POST("/push", controller.BlogCreate)
	}

	router.Run(":8080")
}
