package main

import (

	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"

	"os"
	_"time"

	_"fmt"
)

type Client struct {
	num_cliente uint32  //identificador	
}

func testComm_emptyCount()  {

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/commEmptyCount/data")
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyCount(size,f)
	}

}

func testComm_bulkFind()  {

	f, err := os.Create("./data/commEmptyCount/data")
	check(err)
	defer f.Close()

	for _, size := range tamannos {
		commTest_emptyBulkFind(size,f)
	}

}
