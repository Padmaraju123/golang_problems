package main

import "fmt"

func add_two_numbers(num1, num2 int) int {
	return num1 + num2
}

func main() {
	var num1, num2 int

	//scan method scan the user values with out consideration of spaces and new lines
	//scanln method scan the user values one by one in only one line it will terminate the next line

	fmt.Scanln(&num1, &num2)
	fmt.Println(add_two_numbers(num1, num2))
}
