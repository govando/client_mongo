package main

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"strings"
)

type Insert struct{
	a int
}

func (Insert)  insert_benchmark(tipo_operacion[] uint32,data_base_ string){
	Host := "127.0.0.1:27017"
	mgoSession, _ := mgo.Dial(Host)
	defer mgoSession.Close()
	var count uint32
	fmt.Println("INICIO - BORRADO DE COLECCIONES")
	drop_collections(tamannos, data_base_)
	fmt.Println("DATABASE: ",data_base_)
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
	tfin := time.Now()
	fmt.Println("FIN - INSERTS =>", tfin.Sub(tini))

}

func insert_One(data_base_ string, cant_elementos[] uint32) {
	Host := "127.0.0.1:27017"
	mgoSession, _ := mgo.Dial(Host)
	defer mgoSession.Close()


	for i := 0; i < nbr_tamannnos; i++ {

		col := "BYTES_" + fmt.Sprint(tamannos[i]) + ".DAT"
		sizeData := tamannos[i] - bytesReserved
		data := strings.Repeat("a", (int)(sizeData))

		var doc0,doc1 Document
		doc0.Id = 0
		doc0.Cmp1 = 0
		doc0.Cmp2 = 0
		doc0.Cmp3 = 0
		doc0.Data = data
		mgoSession.DB(data_base_).C(col).Insert(&doc0)

		search_value_init := (cant_elementos[i]/2)+1

		doc1.Id = search_value_init
		doc1.Cmp1 = search_value_init
		doc1.Cmp2 = search_value_init
		doc1.Cmp3 = search_value_init
		doc1.Data = data

		mgoSession.DB(data_base_).C(col).Insert(&doc1)


	}
}
