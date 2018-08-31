package lib

import (
	"net/http"
	// "strconv"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"log"
	"time"
	"github.com/rohit-raj/basic-go-webapp/model"
)

func LoginPage(c *gin.Context){
	u1 := uuid.Must(uuid.NewV4())
	log.Printf("UUIDv4: %s\n", u1)
	c.HTML(http.StatusOK, "login.html", nil)
}

func SignupPage(c *gin.Context){
	c.HTML(http.StatusOK, "signup.html", nil)
}

func LoginController(c * gin.Context){
	db := c.MustGet("dbHandler").(*mgo.Database)
	username := c.PostForm("username")
	password := c.PostForm("password")
	
	user := model.User{}
	err := db.C(model.CollectionUsers).Find(bson.M{"username": username}).One(&user)

	/* OR
		queryObj := gin.H{"username" : username}
		err := db.C(model.CollectionUsers).Find(queryObj).One(&user)
	*/

	if compareHashAndPwd([]byte(user.Password), []byte(password)) {
		// log.Println("user ==> ", user.Username)

		if err != nil {
			c.Error(error(err))
			log.Println("error in mongo query")
		}
		myobj := gin.H{"username" : user.Username, "Created On" : user.CreatedOn}
		
		c.JSON(http.StatusOK, myobj)
	} else{
		c.JSON(http.StatusOK, gin.H{"user" : "No user"})
	}
}

func SignupController(c * gin.Context){
	db := c.MustGet("dbHandler").(*mgo.Database)
	username := c.PostForm("username")
	password := c.PostForm("password")

	hash := hashAndSalt([]byte(password))
	u1 := uuid.Must(uuid.NewV4()).String()

	secs := time.Now().Unix()

	user := model.User{Id : u1, Username : username, Password : hash, CreatedOn : secs, UpdatedOn : secs}
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
		return
	}
	
	err = db.C(model.CollectionUsers).Insert(user)

	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"username" : username, "password" : hash, "created_at" : secs, "_id" : u1})
}

func hashAndSalt(pwd []byte) string{
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil{
		log.Println("err in generating password : ", err)
	}
	return string(hash)
}

func compareHashAndPwd(hashedPwd, pwd []byte) bool{
	err := bcrypt.CompareHashAndPassword(hashedPwd, pwd)
	if err != nil{
		return false
	}
	return true
}

