package main

import (

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
)


func client()  {


	var host = "127.0.0.1:27017"

	mgoSession, _ := mgo.Dial(host)

	// Collection People
	c := mgoSession.DB("test").C("people")

	err := c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
		&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})


	if err != nil {
		panic(err)
	}

	fmt.Println("ad")

}

func create_db()  {
	fmt.Println("bb")
}

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}