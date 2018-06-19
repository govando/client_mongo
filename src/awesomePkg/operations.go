package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"time"
)

func create_db()  {

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	db := "test"
	coll := "emptyColl"

	exist := 0

	// Collection People
	conn := mgoSession.DB(db)

	names, err := conn.CollectionNames()
	if err != nil {
		fmt.Println("Failed to get coll names: %v", err)
		return
	}

	// Simply search in the names slice, e.g.
	for _, name := range names {
		if name == coll {
			fmt.Println("The collection exists!")
			exist=1
			break
		}
	}

	if exist == 0{
		// Create a Collection
		err := conn.C(coll).Create(&mgo.CollectionInfo{
			DisableIdIndex:false, Capped: false, ValidationLevel: "off"})
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Empty collection created")
			fmt.Println("----------------------")
		}
	}

	time.Sleep(1000)

}
