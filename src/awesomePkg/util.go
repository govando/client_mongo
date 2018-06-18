package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func average(times []float64)  {

	var avg,value float64
	//var i int
	for _, value = range times {
		avg += value
	}
	fmt.Printf("total: %f \n",avg)
}

func sd()  {

}

