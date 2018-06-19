package main

import (
	"fmt"
	_"gopkg.in/mgo.v2"
	_"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


type Stats struct {
	Updated      		int
	Removed      		int
	Matched      		int
	operationType       string
	nanoElapsedComun    int64
	operation    		string
}


type Document struct {
	Id   uint32 `bson:"_id"`
	Cmp1 uint32 `bson:"cmp1"`
	Cmp2 uint32 `bson:"cmp2"`
	Cmp3 uint32 `bson:"cmp3"`
	Data string `bson:"data"`
}

type Estructura struct {
	a int
}

var (
	tamannos              = []uint32{100, 400, 700, 1000, 4000, 7000, 10000, 40000, 70000, 100000, 400000, 700000, 800000, 900000, 1000000, 2000000, 3000000, 4000000, 5000000, 6000000, 7000000, 8000000, 9000000, 10000000, 11000000, 12000000, 13000000, 14000000, 15000000, 16000000}
	//tamannos				= []uint32{ 8000000, 16000000, 100, 400, 700, 1000, 4000, 7000 } //13000000, 14000000, 15000000,
	bytesReserved			= (uint32)(55) 	//reservados para los tipos de datos uint32
	nbr_tamannnos			= len(tamannos)
	nbr_operationsByCol		= (uint32(100)) 				//cantidad de operaciones en 16MB

	operations_perCol		= []uint32{ }		//cantidad de datos a operar para cada coleccion //valido para
	sin_saturationPoint_per_col	= []uint32{ }		//cantidad de operaciones que no saturan
	saturationPoint_per_col	= []uint32{ }		//operacion en la que se produce saturacion de ram
	ram_mongo				= (uint32)(5000)*1000000	//indica el punto de saturación de mongo cache
	data_base				= "db_test"
	data_base_sat			= "db_test_sat"
	SATURACION				= 1

	OPERACION_IGUAL			= 0
	OPERACION_DISTINTO		= 1
	OPERACION_MAYOR			= 2
	OPERACION_MENOR			= 3
	OPERACION_MAYOR_IGUAL	= 4
	OPERACION_MENOR_IGUAL	= 5

	INSERTAR 				= 1
	EDITAR					= 2
	BORRAR					= 3

	ONE 		= 1
	MULTI		= 2
	EXISTE		= 1
	NO_EXISTE	= 0
)


func (Estructura) configuracion(){

	ram_mongo_sin_sat := uint32( float32(ram_mongo)*0.70) //mongo podria ocupar el 15% en ram
	for i := 0; i < nbr_tamannnos; i++ {
		sin_saturationPoint_per_col = append(sin_saturationPoint_per_col, ram_mongo_sin_sat/tamannos[i]);
		saturationPoint_per_col = append(saturationPoint_per_col, uint32( float32(ram_mongo)*2.5)/tamannos[i]+1);
		operations_perCol = append(operations_perCol, nbr_operationsByCol* 16000000/tamannos[i]);
	}
}

func drop_collections(tamannos []uint32, data_base_ string) {
	Host := "127.0.0.1:27017"
	mgoSession, _ := mgo.Dial(Host)
	defer mgoSession.Close()

	var col string
	for i := 0; i < len(tamannos); i++ {
		col = "BYTES_" + fmt.Sprint(tamannos[i]) + ".DAT"
		mgoSession.DB(data_base_).C(col).DropCollection()
		fmt.Println(col)
	}
}

func createIndex(db string){
	Host := "127.0.0.1:27017"
	mgoSession, _ := mgo.Dial(Host)
	defer mgoSession.Close()

	var col string
	for i := 0; i < len(tamannos); i++ {
		col = "BYTES_" + fmt.Sprint(tamannos[i]) + ".DAT"
		index := mgo.Index{
			Key :  []string{"cmp1"},
		}
		err := mgoSession.DB(db).C(col).EnsureIndex(index)
		fmt.Println(col)
		if err != nil {
			panic(err)
		}
	}
}


func dropIndex(db string){
	Host := "127.0.0.1:27017"
	mgoSession, _ := mgo.Dial(Host)
	defer mgoSession.Close()

	var col string
	for i := 0; i < len(tamannos); i++ {
		col = "BYTES_" + fmt.Sprint(tamannos[i]) + ".DAT"

		mgoSession.DB(db).C(col).DropIndex("cmp1")
		mgoSession.DB(db).C(col).DropIndex("cmp2")
		mgoSession.DB(db).C(col).DropIndex("cmp1", "cmp2")
		fmt.Println(col)
		/*if err != nil || err1 != nil {
			panic(err)
		}*/
	}
}

func wait4db()  {
	Host := "127.0.0.1:27017"
	mgoSession, err := mgo.Dial(Host)

	for err != nil {
		mgoSession, err = mgo.Dial(Host)
	}
	defer mgoSession.Close()


}


func select_operation(col string, OPCION int, db string,search_value_init uint32,search_value_end uint32,EXISTE_ int) {

	var Host = "127.0.0.1:27017"
	mgoSession, _ := mgo.Dial(Host)
	defer mgoSession.Close()
	var err error
	var results []Document


	colQuerier := bson.M{"cmp1": bson.M{"$gte": search_value_init,"$lte": search_value_end}}
	err = mgoSession.DB(db).C(col).Find(colQuerier).All(&results)


	if err!=nil {
		fmt.Print("ERROR AL SELECCIONAR")
	}
}