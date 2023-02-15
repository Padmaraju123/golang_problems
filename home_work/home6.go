package main

import (
	"fmt"
	"math"
)

func cubb(n float64) float64 {

	out := math.Pow(n,4)
	fmt.Printf("%T %v\n",out,out)

	//converting float64 to int
	out1:=int(out)
	fmt.Printf("%T %v\n",out1,out1)
	
	return out

}

func main() {
	var num float64
	fmt.Scanln(&num)
	fmt.Println(cubb(num))
}
