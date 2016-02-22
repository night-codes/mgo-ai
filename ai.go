package ai

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session = &mgo.Session{}
var collection = &mgo.Collection{}

// Counter struct
type Counter struct {
	ObjectID bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	ID       string        `json:"id"`
	Seq      uint64        `json:"seq"`
}

// Connect to database
func Connect(c *mgo.Collection) {
	collection = c
	session = c.Database.Session
}

// Next sequence of AutoIncrement
func Next(name string) uint64 {
	connectionCheck()
	var result Counter
	if _, err := collection.Find(bson.M{"id": name}).Apply(mgo.Change{
		Update:    bson.M{"$set": bson.M{"id": name}, "$inc": bson.M{"seq": 1}},
		Upsert:    true,
		ReturnNew: true,
	}, &result); err != nil {
		fmt.Println("Autoincrement error(1):", err.Error())
	}
	return result.Seq
}

// Cancel is decrement counter value
func Cancel(name string) {
	connectionCheck()
	if err := collection.Update(bson.M{"id": name}, bson.M{"$inc": bson.M{"seq": -1}}); err != nil {
		fmt.Println("Autoincrement error(2):", err.Error())
	}
}

func connectionCheck() {
	if err := session.Ping(); err != nil {
		session.Refresh()
	}
}
