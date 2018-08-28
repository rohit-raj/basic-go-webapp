package lib

import (
	"net/http"
	// "strconv"
	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", nil)
}

func SignupPage(c *gin.Context){
	c.HTML(http.StatusOK, "signup.html", nil)
}

func LoginController(c * gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.JSON(http.StatusOK, gin.H{"username" : username, "password" : password})
}