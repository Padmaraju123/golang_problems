package main

import "fmt"

func main() {

	var size int
	fmt.Scan(&size)

	numarray := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&numarray[i])
		fmt.Printf("%T %v\n", numarray[i], numarray[i])
	}

	arrSum := 0

	for _, a := range numarray {
		arrSum = arrSum + a
	}
	fmt.Println(arrSum)
}
