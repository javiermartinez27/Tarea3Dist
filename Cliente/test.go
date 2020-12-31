package main

import (
	"fmt"
)

func main() {
	// var consistencia [][]string
	// // var a [2]string
	// // a[0] = "Hello"
	// // a[1] = "World"
	// consistencia = [["test.cl 9002 225,0,0 20.0.0.1"] ["google.cl 9003 222,0,0 19.0.0.1"] ["goo.cl 9001 2,0,0 10.0.0.5"]]
	// for i, s := range consistencia {
	// 	fmt.Println("Esto es el len de consistencia:")
	// 	fmt.Println(len(consistencia[i]))
	// 	// fmt.Printf("%+q", len(consistencia))

	// }

	println("Simple Array:")
	var arrayint = [...]int{1, 2, 3, 4} //assign
	fmt.Println(arrayint, "\n")
	println("Array of arrays:")
	var arrayofarrays [3][len(arrayint)]int
	for i := range arrayofarrays { //assign
		arrayofarrays[i] = arrayint
	}
	fmt.Println(arrayofarrays, "\n")

	var arrayofslice [len(arrayofarrays)][]int
	for i := range arrayofarrays { // assign
		arrayofslice[i] = arrayofarrays[i][:]
	}
	println("Slice of slices:")
	var sliceofslices [][]int
	sliceofslices = arrayofslice[:]
	fmt.Println(sliceofslices, "\n")

	for _, s := range sliceofslices {
		fmt.Println(len(s), "\n")
	}

}
