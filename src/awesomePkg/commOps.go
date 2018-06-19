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


func commTest_emptyBulkFind(size uint32, f *os.File) {

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	var contentArray []interface{}
	var doc Doc

	bulk := mgoSession.DB(db).C(coll).Bulk()



	data := strings.Repeat("a", (int)(sizeData))

	for j := uint32(0); j < tipo_operacion[i]; j++ {
		create_document(j)

	}






	count = 0
	bulk := mgoSession.DB(data_base_).C(col).Bulk()



	var contentArray []interface{}
	for j := uint32(0); j < tipo_operacion[i]; j++ {

		var doc Document
		doc.Id = j
		doc.Cmp1 = j
		doc.Cmp2 = j
		doc.Cmp3 = j
		doc.Data = data


		contentArray =  append(contentArray, &doc)
		count++


		//fmt.Println("count: ",count," contentarray:",contentArray," j: ",j)
		if  count==1000 || j==tipo_operacion[i]-1  || (uint32)(16000000)<(count+1)*tamannos[i] {
			fmt.Println("IN: len: ", len(contentArray),
				" tipo_operacion: ",tipo_operacion[i]," count:",count, " j: ",j)
			//mgoSession.DB(data_base_).C(col).Insert(&doc)
			bulk.Insert(contentArray...)
			_, err := bulk.Run()
			if err != nil {
				fmt.Println("ERROR! y: contentArray: ",contentArray," len: ", len(contentArray))
				panic(err)
			}


			count=0
			//fmt.Println(time.Now(), " col:", col, "( i :", i, "-", nbr_tamannnos, ") (j:", j, "-", tipo_operacion[i], ")")
			bulk = mgoSession.DB(data_base_).C(col).Bulk()
			contentArray = nil
			fmt.Println("END: len: ", len(contentArray),
				" tipo_operacion: ",tipo_operacion[i]," count:",count, " j: ",j)
		}

	}






	tini := time.Now()
	var col string
	for i := 0; i < nbr_tamannnos; i++ {

		col = "BYTES_" + fmt.Sprint(tamannos[i]) + ".DAT"
		sizeData := tamannos[i] - bytesReserved
		data := strings.Repeat("a", (int)(sizeData))

		count=0
		bulk := mgoSession.DB(data_base_).C(col).Bulk()
		var contentArray []interface{}
		for j := uint32(0); j < tipo_operacion[i]; j++ {

			var doc Document
			doc.Id = j
			doc.Cmp1 = j
			doc.Cmp2 = j
			doc.Cmp3 = j
			doc.Data = data


			contentArray =  append(contentArray, &doc)
			count++


			//fmt.Println("count: ",count," contentarray:",contentArray," j: ",j)
			if  count==1000 || j==tipo_operacion[i]-1  || (uint32)(16000000)<(count+1)*tamannos[i] {
				fmt.Println("IN: len: ", len(contentArray),
					" tipo_operacion: ",tipo_operacion[i]," count:",count, " j: ",j)
				//mgoSession.DB(data_base_).C(col).Insert(&doc)
				bulk.Insert(contentArray...)
				_, err := bulk.Run()
				if err != nil {
					fmt.Println("ERROR! y: contentArray: ",contentArray," len: ", len(contentArray))
					panic(err)
				}


				count=0
				//fmt.Println(time.Now(), " col:", col, "( i :", i, "-", nbr_tamannnos, ") (j:", j, "-", tipo_operacion[i], ")")
				bulk = mgoSession.DB(data_base_).C(col).Bulk()
				contentArray = nil
				fmt.Println("END: len: ", len(contentArray),
					" tipo_operacion: ",tipo_operacion[i]," count:",count, " j: ",j)
			}

		}
	}
}



//'n' pruebas de comunicacion para cada tamaño
func commTest_emptyCount(size uint32, f *os.File)  {


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