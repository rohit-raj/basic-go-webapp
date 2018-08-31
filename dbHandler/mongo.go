package dbHandler

import (
	"log"
	"gopkg.in/mgo.v2"
	"github.com/rohit-raj/basic-go-webapp/config"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

// Connect connects to mongodb
func Connect() {
	configuration := config.Config()
	uri := configuration.Mongo_Url

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		log.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	log.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}