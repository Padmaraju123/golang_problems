package main

import (
	"fmt"
	"math"
)

func main(){
	var num int
	fmt.Scanln(&num)
	val := math.Pow(float64(num),3)
	fmt.Printf("%T %v\n",val,val)
}