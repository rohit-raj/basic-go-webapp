package lib

import (
	"net/http"
	// "strconv"
	"github.com/gin-gonic/gin"
)

func DispHome (c *gin.Context){
	//for the time being nil, if we need to send dynamic data, then we can create interface to send it

	c.HTML(http.StatusOK, "homepage.html", nil)
}