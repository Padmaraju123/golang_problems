package main

import "fmt"

func get(m, n int) int {
	dd := m + n
	return dd
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(get(m, n))
}
