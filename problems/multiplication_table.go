package main

import "fmt"

func multiplication_table(num int) {
	for i := 1; i <=10; i++ {
		val := num * i
		fmt.Printf("%d * %d = %d\n", num, i, val)
	}
}

func main() {
	var num int
	fmt.Scanln(&num)
	multiplication_table(num)
}
