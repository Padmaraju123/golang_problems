package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n = 100
	fmt.Printf("%T %v\n", n, n)

	//convert int to float32
	n1 := float32(n)
	fmt.Printf("%T %v\n", n1, n1)

	//convert int to float64
	n2 := float64(n)
	fmt.Printf("%T %v\n", n2, n2)

	//convert int to string
	n3 := strconv.Itoa(n)
	fmt.Printf("%T %v\n", n3, n3)

	var p = 25.44

	//convert float to int
	n4 := int(p)
	fmt.Printf("%T %v\n", n4, n4)

	//convert float to string
	n5 := strconv.FormatFloat(p, 'E', 2, 64) //64 is float64 and 2 is digits after decimal point
	fmt.Printf("%T %v\n", n5, n5)
	fmt.Println(n5)

	var g = "34"
	//convert string to int possible only if the string is integer
	n6, error := strconv.Atoi(g)
	fmt.Println(error)
	fmt.Printf("%T %v\n", n6, n6)

	var h = "35"
	//convert string to int
	n7, er := strconv.ParseInt(h, 10, 0)
	fmt.Println(er)
	fmt.Printf("%T %v\n", n7, n7)

	var w = "36.7"
	//convert string to int
	n8, er := strconv.ParseFloat(w, 64)
	fmt.Println(er)
	fmt.Printf("%T %v\n", n8, n8)

}
