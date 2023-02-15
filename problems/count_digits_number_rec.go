package main

import "fmt"

var count int = 0

func count_digits_number_rec(num int) int {

	if num > 0 {
		count = count + 1
		count_digits_number_rec(num / 10)
	}
	return count
}

func main() {
	var num int
	fmt.Scanln(&num)
	out :=count_digits_number_rec(num)
	fmt.Println(out)
}
