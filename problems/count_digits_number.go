package main

import "fmt"

func count_digits_number(num int)int{

	count := 0

	for num > 0 {
		num = num / 10
		fmt.Println(num)
		count++
	}
	return count
}

func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Println(count_digits_number(num))
}
