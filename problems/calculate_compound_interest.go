package main

import (
	"fmt"
	"math"
)

func calculate_compound_interest(pa,si,y int)float64{
	val := 1+(float64(si)/float64(100))
	a:= math.Pow(val,float64(y))
	return float64(pa)*a - float64(pa)
}

func main() {
	var pa, si, years int
	fmt.Scan(&pa, &si,&years)
	fmt.Println(calculate_compound_interest(pa,si,years))
}
