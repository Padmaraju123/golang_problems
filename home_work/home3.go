package main

import (
	"fmt"
	"math"
)

var p, y, i, fci, b float64

func main() {

	fmt.Scan(&p, &y, &i)

	fci = p * math.Pow(1+i/100, y)

	fmt.Printf("%T %f\n", fci, fci)

	b = fci - p

	fmt.Printf("%T %f\n", b, b)
}
