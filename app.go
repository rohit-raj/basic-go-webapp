package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rohit-raj/basic-go-webapp/lib"
)

func main(){
	router :=gin.Default();
	router.LoadHTMLGlob("views/*.html")
	router.Static("/public", "./public")

	/**
	 * APIs
	 */

	router.GET("/", lib.DispHome)
	router.GET("/login", lib.LoginPage)
	// router.POST("/login", lib.LoginController)
	router.GET("/signup", lib.SignupPage)
	// router.POST("/signup", lib.SignupPage)
	
	router.Run(":8080")
}