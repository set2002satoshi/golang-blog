package main

import (
	// "net/http"

	// "github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"

	// "github.com/set2002satoshi/golang-blog/controller"
	"github.com/set2002satoshi/golang-blog/router"
	"github.com/set2002satoshi/golang-blog/db"
)

func main() {
	db.DBInit()
	
	router.SetUpRouter()
	

}

