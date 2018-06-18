package main

import (

/*	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
	"strings"
	*///"awesomePkg/awesomeClient"
)


func main() {

	//var c
//    client()
    create_db()
	testComm_emptyCount()
//



	//err := c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
	//	&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})


/*
	data := strings.Repeat("a", 64)
	fmt.Println("len: ", len(data))
	var q Document
	//q.Id = 0
	q.Data = data

	time.Sleep(10)
	tini := time.Now()
	query,_ := conn.Find(bson.M{"_id": data}).Count()
	//conn.C(coll).Find(q)
	total :=  time.Since(tini).Nanoseconds()
	fmt.Println("time connection: ", total)
	fmt.Println("----------------------")

	fmt.Print(query)
	//fmt.Println("sizeof uint32: ", unsafe.Sizeof(data), " bytes \n sizeof data: ", unsafe.Sizeof(data), " bytes")
*/


}

func create_empty_ix_collection(){

}

