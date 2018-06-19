package main


import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"gopkg.in/mgo.v2"
)

type Delete struct{
	a int
}
func (Delete)  Delete(ES_SATURACION int,  TAREA_LOGICA int, ES_INDICE int, EXISTE_ int) {

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
		run_benchmark_delete(ONE, EXISTE_, db, cant_elementos,ES_SATURACION)
	case 2:
		run_benchmark_delete(MULTI, EXISTE_, db,cant_elementos,ES_SATURACION)
	}

}

func run_benchmark_delete(OPCION int, EXISTE_ int, db string, cant_elementos[] uint32, ES_SATURACION int){
	var search_value_init uint32
	var search_value_end uint32

	for i := 0; i < nbr_tamannnos; i++ {
		col := "BYTES_" + fmt.Sprint(tamannos[i]) + ".DAT"

		if ES_SATURACION==1{
			search_value_init = (cant_elementos[i]/2)+1
			search_value_end = cant_elementos[i] -1
			select_operation(col, MULTI,db,search_value_init,search_value_end, EXISTE_)
		}

		search_value_init = 0
		search_value_end = (cant_elementos[i]/2)-1
		delete_operation(col, OPCION,db,search_value_init,search_value_end, EXISTE_)

		//a pesar de que los datos son borrados, se mantienen en RAM
		if ES_SATURACION==1{
			search_value_init = 0
			search_value_end = (cant_elementos[i]/2)
			select_operation(col, MULTI,db,search_value_init,search_value_end, EXISTE_)
		}
		search_value_init = (cant_elementos[i]/2)+1
		search_value_end = cant_elementos[i] -1
		delete_operation(col, OPCION,db,search_value_init,search_value_end, EXISTE_)
	}

	//----------reinsertar elementos borrados
	if EXISTE_ == EXISTE {
		var insert Insert
		if OPCION == MULTI {
			if ES_SATURACION == 1 {
				insert.insert_benchmark(saturationPoint_per_col, data_base_sat)
			} else {
				insert.insert_benchmark(sin_saturationPoint_per_col, data_base)
			}
		} else if OPCION == ONE {
			if ES_SATURACION == 1 {
				insert_One(data_base_sat, cant_elementos)
			} else {
				insert_One(data_base, cant_elementos)
			}
		}
	}
}


func delete_operation(col string, OPCION int, db string,search_value_init uint32,search_value_end uint32,EXISTE_ int) {

	var Host = "127.0.0.1:27017"
	mgoSession, _ := mgo.Dial(Host)
	defer mgoSession.Close()
	var err error

	if EXISTE_==0{
		search_value_init	= 4294967295
		search_value_end	= 4294967295
	}


	switch OPCION {
	case ONE:
		colQuerier := bson.M{"cmp1": search_value_init}
		err = mgoSession.DB(db).C(col).Remove(colQuerier)
		fmt.Println("DEL 1 en ",col ," - buscando cmp1:",search_value_init," cambiando cmp2:",new_value)
	case MULTI:
		colQuerier := bson.M{"cmp1": bson.M{"$gte": search_value_init,"$lte": search_value_end}}
		_, err = mgoSession.DB(db).C(col).RemoveAll(colQuerier)
		fmt.Println("DEL MANY  en ",col ,"  - buscando cmp1:",search_value_init,"-",search_value_end ,"cambiando cmp2:",new_value)

	}

	if err!=nil {
		fmt.Print("ERROR AL BORRAR")
	}
}


