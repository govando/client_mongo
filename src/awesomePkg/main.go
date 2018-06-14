package main

import (

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
	"log"
	_"strings"
	"strings"
	//"awesomePkg/awesomeClient"

)


func main() {

	//var c
//    client()
    create_db()

	bulkOp()
    
	var host = "127.0.0.1:27017"

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	db := "test"
	coll := "emptyColl"

	exist := 0

	// Collection People
	conn := mgoSession.DB(db)


	names, err := conn.CollectionNames()
	if err != nil {
		// Handle error
		log.Printf("Failed to get coll names: %v", err)
		return
	}

	// Simply search in the names slice, e.g.
	for _, name := range names {
		if name == coll {
			//log.Printf("The collection exists!")

			exist=1
			break
		}
	}

	if exist == 0{
		// Create a Collection
		err := conn.C(coll).Create(&mgo.CollectionInfo{DisableIdIndex:false, Capped: false})
		if err != nil {
			panic(err)
		} else {
			//fmt.Println("Empty collection created")
			//fmt.Println("----------------------")
		}
	}




	//err := c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
	//	&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})



	data := strings.Repeat("a", 64)
	fmt.Println("len: ", len(data))
	var q Document
	//q.Id = 0
	q.Data = data

	time.Sleep(10)
	tini := time.Now()
	query,_ := conn.C(coll).Find(bson.M{"_id": data}).Count()
	//conn.C(coll).Find(q)
	total :=  time.Since(tini).Nanoseconds()
	fmt.Println("time connection: ", total)
	fmt.Println("----------------------")

	fmt.Print(query)
	//fmt.Println("sizeof uint32: ", unsafe.Sizeof(data), " bytes \n sizeof data: ", unsafe.Sizeof(data), " bytes")



}

func create_empty_ix_collection(){

}

type Document struct {
	Id   uint32 `bson:"_id"`
	Data string `bson:"data"`
}