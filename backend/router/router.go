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
	
	v1 := router.Group("/api")
	{
		v1.POST("/user", controller.CustomerCreate)
	}
	v2:= router.Group("/api")
	{
		v2.POST("/Certification", controller.Login)
	}

	router.Run(":8080")
}