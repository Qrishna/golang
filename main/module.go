package main

import (
	"fmt"
	"reflect"
	"strings"
)

func max (array [] int) int{
	mymax := array[0]
	for _, value := range(array){
		if value > mymax{
			mymax = value
		}
	}
	return mymax
}

func min (array [] int) int{
	mymin := array[0]
	for _, value := range(array){
		if value < mymin{
			mymin = value
		}
	}
	return mymin
}

func sum (array [] int) int{
	mysum := 0
	for _, value := range(array){
		mysum += value
	}
	return mysum
}

func main() {
	fmt.Println("Hello, World!")

	integer := 10
	float := 100.12
	str := "this is a string"
	const CODE int = 20

	// array
	array := [] int {1, 3, 14, 50, 6, -12, 0, -1}


	// 2d array
	array2d := [][] int {{1, 2, 3, 4},
							{5, 6, 7, 8},
							{9, 10, 11, 12}}
	fmt.Println(integer)
	fmt.Println(float)
	fmt.Println(str)

	fmt.Printf("integer is of type %T\n", integer)
	fmt.Printf("float is of type %T\n", float)
	fmt.Printf("str is of type %T\n", str)
	fmt.Printf("This is a constant: %d\n", CODE)

	var x int
	for x = 0; x < len(array); x++{
		fmt.Printf("array[%d] = %d\n", x, array[x])
	}

	/*
	var i, j int

	rows := len(array2d)
	cols := len(array2d[0])

	fmt.Printf("rows of my array is: %d\n", rows)
	fmt.Printf("cols of my array is: %d\n", cols)

	for i = 0; i<rows; i++{
		for j = 0; j<cols; j++{
			fmt.Printf("array2d[%d][%d] = %d\n", i, j, array2d[i][j])
		}
	}
	*/

	// This is the same as above
	for i, rows :=range(array2d){
		for j, cols :=range rows{
			fmt.Printf("array2d[%d][%d] = %d\n",i, j, cols)
		}
	}


	fmt.Printf("Max of my array is: %d\n", max(array))
	fmt.Printf("Min of my array is: %d\n", min(array))
	fmt.Printf("Sum of my array is: %d\n", sum(array))


	// map
	/*
	dict := make(map[string]string)	// make the map
	dict["name"] = "krsna"
	dict["address"] = "dallas"
	*/

	// This is the same as above
	dict := map[string] string{
		"name" : "krsna",
		"value": "dallas",
	}

	fmt.Printf("dict[%s] = %s\n", "name", dict["name"])
	fmt.Printf("dict[%s] = %s\n", "name", dict["address"])
	fmt.Printf("number of items in map = %d\n", len(dict))

	// iterating over maps
	for key, value := range dict{
		fmt.Println("Key:", key, "Value:", value, "dict[key]:", dict[key])
	}


	sensors := map[string] float32{
		"rsc": 3711.2,
		"r":   2138.1,
		"gri": 1908.1,
		"adg": 912.5,
	}

	for key, value := range sensors{
		fmt.Println("Key:", key, "Value:", value, "sensors[key]:", sensors[key])
	}


	var keys []string
	for key, _ := range sensors{
		keys = append(keys, key)
	}

	var values []float32
	for _, value := range sensors{
		values = append(values, value)
	}

	fmt.Println("Keys = ", keys, keys[0], keys[1], keys[2])
	fmt.Printf("%T\n", keys)
	fmt.Println(reflect.TypeOf(keys))


	fmt.Println("Values = ", values, values[0], values[1], values[2])
	fmt.Printf("%T\n", values)
	fmt.Println(reflect.TypeOf(values))


	// String manipulation
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Split("chicken,turkey,duck ", ","))
}
