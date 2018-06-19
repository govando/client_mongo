package main

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"gopkg.in/mgo.v2"
)

var (
	new_value = 70
)
type Update struct{
	a int
}
func (Update)  Update(ES_SATURACION int,  TAREA_LOGICA int, ES_INDICE int, EXISTE_ int) {

	var db string
	var cant_elementos[] uint32

	db = data_base
	cant_elementos = sin_saturationPoint_per_col
	if ES_SATURACION == SATURACION {
		db = data_base_sat
		cant_elementos = saturationPoint_per_col
	}

	if ES_INDICE==1{
		dropIndex(db)
		createIndex(db)
	} else if ES_INDICE==0{
		dropIndex(db)
	}

	//Tareas: Update one, Update Multi=True Existe
	switch TAREA_LOGICA {
	case 1:
		run_benchmark_update(ONE, EXISTE_, db, cant_elementos,ES_SATURACION)
	case 2:
		run_benchmark_update(MULTI, EXISTE_, db,cant_elementos,ES_SATURACION)
	}

}

func run_benchmark_update(OPCION int, EXISTE_ int, db string, cant_elementos[] uint32, ES_SATURACION int){
	var search_value_init uint32
	var search_value_end uint32

	for i := 0; i < nbr_tamannnos; i++ {
		col := "BYTES_" + fmt.Sprint(tamannos[i]) + ".DAT"

		for j := uint32(0); j < nbr_operationsByCol; j++ {
			if j%2==0 {
				search_value_init = 0
				search_value_end = (cant_elementos[i]/2)-1
			} else {
				search_value_init = (cant_elementos[i]/2)+1
				search_value_end = cant_elementos[i] -1
			}

			// Si hay saturacion, se llena la ram -a través de un Select -con los elementos
			// contrarios a los que se actualizaran
			if ES_SATURACION==1{
				if j%2==1 {
					search_value_init = 0
					search_value_end = (cant_elementos[i]/2)-1
				} else {
					search_value_init = (cant_elementos[i]/2)+1
					search_value_end = cant_elementos[i] -1
				}
				select_operation(col, MULTI,db,search_value_init,search_value_end, EXISTE_)
			}
			update_operation(col, OPCION,db,search_value_init,search_value_end, EXISTE_)
		}
	}
}


func update_operation(col string, OPCION int, db string,search_value_init uint32,search_value_end uint32,EXISTE_ int) {

	var Host = "127.0.0.1:27017"
	mgoSession, _ := mgo.Dial(Host)
	defer mgoSession.Close()
	var err error

	if EXISTE_==0{
		search_value_init	= 4294967295
		search_value_end	= 4294967295
	}

	change := bson.M{"$set": bson.M{ "cmp2" : new_value}}

	switch OPCION {
	case ONE:
		colQuerier := bson.M{"cmp1": search_value_init}
		err = mgoSession.DB(db).C(col).Update(colQuerier, change)
		fmt.Println("UPD 1 en ",col ," - buscando cmp1:",search_value_init," cambiando cmp2:",new_value)
	case MULTI:
		colQuerier := bson.M{"cmp1": bson.M{"$gte": search_value_init,"$lte": search_value_end}}
		_, err = mgoSession.DB(db).C(col).UpdateAll(colQuerier, change)
		fmt.Println("UPD MANY  en ",col ,"  - buscando cmp1:",search_value_init,"-",search_value_end ,"cambiando cmp2:",new_value)

	}

	if err!=nil {
		fmt.Print("ERROR AL UPDATEAR")
	}
}


