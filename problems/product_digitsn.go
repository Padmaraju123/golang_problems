package main

import "fmt"

func product_digitsN(num int) int {
	p := 1
	for num > 10 {
		re := num % 10
		num = num / 10
		p *= re
	}
	return p
}

func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Println(product_digitsN(num))
}
