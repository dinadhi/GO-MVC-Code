package mongodb

import (
	"mvc/option"
	"labix.org/v2/mgo"
)

// DB .
var DB *mgo.Database

func init() {


	session, err := mgo.Dial(option.Machine.MongoURL)
	if err != nil {

		panic(err)
	}


	DB = session.DB(option.Machine.MongoDBName)
}
