package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rohit-raj/basic-go-webapp/dbHandler"
)

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func Connect(c *gin.Context) {
	s := dbHandler.Session.Clone()

	defer s.Close()

	c.Set("dbHandler", s.DB(dbHandler.Mongo.Database))
	c.Next()
}

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()

	// TODO: Handle it in a better way
	if len(c.Errors) > 0 {
		c.HTML(http.StatusBadRequest, "400", gin.H{
			"errors": c.Errors,
		})
	}
}
