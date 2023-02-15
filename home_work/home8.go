package main

import "fmt"

func fscc(n int) int {
	var g = 1
	for n > 0 {
		g *= n
		n--
	}
	return g
}

func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Println(fscc(num))

}
