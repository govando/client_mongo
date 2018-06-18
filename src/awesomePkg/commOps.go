package main

import (
	"fmt"
	"strings"
	"time"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"os"
	"bytes"

)

var (
	query		*mgo.Query
	count		int
	err			error
	total_time	int64
)

func bulkOp(){
	fmt.Println("bulkOp")
}


func commTest_emptyCount(size uint32, f *os.File)  {

	//f, err := os.OpenFile("./data/commEmptyCount/data", os.O_RDWR|os.O_APPEND, 0660);
	//check(err)
	//defer f.Close()

	var buffer bytes.Buffer
	var times[] float64
	buffer.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()


	conn := mgoSession.DB(db).C(coll)

	data := strings.Repeat("a", int(size))
	for i := 0; i < n_pruebas; i++ {
		tini := time.Now()
		count, _ = conn.Find(bson.M{"_id": data}).Count()
		total_time = time.Since(tini).Nanoseconds()

		_, err = buffer.WriteString(fmt.Sprintf("%d\t%d\t%f\n",i,size,float64(total_time)/float64(1000000)))
		check(err)

		//no almaceno los mayores a un milisegundos (outliers de 20+ms)
		if total_time > 1000000 {
			times = append(times,float64(total_time)/float64(1000000))
		}

	}

	_, err = f.WriteString(buffer.String())
	check(err)

	average(times)
	fmt.Println("----------------------")
}