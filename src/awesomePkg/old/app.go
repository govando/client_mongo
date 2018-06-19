package main


import (
	"os"
	_ "fmt"
	"strconv"
	"fmt"
)

//Inputs:
//		--Número de la tarea a realizar (1:Insert 2:Update 3:Delete)
//		--Manejo de ram - 0:Sin Saturacion 1:con saturacion
//		--TAREA_LOGICA: 1:Uno  2:Varios
//		--Indice:   0:sin ix  1:con ix
//		--Existe:	0:no existe   1:existe
func main() {
	var config Estructura
	config.configuracion()

	NUM_TAREA, _	 := strconv.Atoi(os.Args[1])
	ES_SATURACION, _ := strconv.Atoi(os.Args[2])
    TAREA_LOGICA, _ := strconv.Atoi(os.Args[3])
    ES_INDICE, _ := strconv.Atoi(os.Args[4])
    EXISTE_, _ := strconv.Atoi(os.Args[5])

    wait4db()

    if NUM_TAREA == INSERTAR {

		var insert Insert

		fmt.Print("Sin saturacion")
		insert.insert_benchmark(sin_saturationPoint_per_col,data_base)
		insert.insert_benchmark(saturationPoint_per_col,data_base_sat)

	}
	if NUM_TAREA == EDITAR {
		var upd Update
		upd.Update(ES_SATURACION, TAREA_LOGICA, ES_INDICE, EXISTE_)

	}
	if NUM_TAREA == BORRAR {
		var del Delete
		del.Delete(ES_SATURACION, TAREA_LOGICA, ES_INDICE, EXISTE_)

	}
	//var del benchmarks.Delete_benchmark
	//var update benchmarks.Update_benchmark

	//del.Delete_benchmark()
}

/*
func exe_cmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println("command is ",cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head,parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}
*/