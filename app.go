package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rohit-raj/basic-go-webapp/lib"
	"github.com/rohit-raj/basic-go-webapp/dbHandler"
	"log"
	"os"
	"github.com/rohit-raj/basic-go-webapp/middleware"
	"github.com/rohit-raj/basic-go-webapp/config"
)

func init(){
	
	env := os.Getenv("env")
	r := []rune(env)

	if len(r) < 1 {
		return
	} else {
		log.Println("connecting mongo")
		dbHandler.Connect()
	}
}

func main(){
	env := os.Getenv("env")
	r := []rune(env)

	if len(r) < 1 {
		log.Println("specify an enviromnent")
		log.Println("dev for dev enviromnent : env=dev go run app.go")
		log.Println("prod for prod enviromnent : env=prod go run app.go")
		return
	}
	router := gin.Default();

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	// Middlewares
	router.Use(middleware.Connect)
	router.Use(middleware.ErrorHandler)

	router.LoadHTMLGlob("views/*.html")
	router.Static("/public", "./public")

	/**
	 * APIs
	 */

	router.GET("/", lib.DispHome)
	router.GET("/login", lib.LoginPage)
	router.POST("/login", lib.LoginController)
	router.GET("/signup", lib.SignupPage)
	router.POST("/signup", lib.SignupController)
	
	configuration := config.Config()
	log.Println("app running on : http://localhost"+configuration.Port)
	router.Run(configuration.Port)
}
