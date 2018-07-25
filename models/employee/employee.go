package employee

import (
"encoding/json"
"time"
mgo "gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"mvc/db/mongodb"
)



// Employee .
type Employee struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Status     string        `json:"status" bson:"status,omitempty"`
	CreateTime time.Time     `json:"create_time" bson:"create_time"`
}

func (e *Employee) String() string {
	b, err := json.Marshal(e)
	if err != nil {

		return err.Error()
	}

	return string(b)
}

// New .
func New(name string) (e *Employee, err error) {


	// verify Employee name
	if name == "" {

		return
	}

	// create new Employee
	e = &Employee{
		Name:       name,
		CreateTime: time.Now(),
	}

	// save Employee into db
	collection := mongodb.DB.C("Employee")
	err = collection.Insert(e)
	if err != nil {

		return
	}


	return
}

// Get .
func Get(id string) (e *Employee, err error) {


	// verify Employee id
	if !bson.IsObjectIdHex(id) {
		return
	}

	// get Employee by id
	e = &Employee{}
	collection := mongodb.DB.C("Employee")
	err = collection.FindId(bson.ObjectIdHex(id)).One(e)
	if err != nil {
		if err == mgo.ErrNotFound {
			err = nil
			e = nil

			return
		}


		return
	}


	return
}
