package main

import "fmt"

func gt(n int) []int {
	var out = []int{}
	i := 0
	for i < n {
		var k int
		fmt.Scanln(&k)
		out = append(out, k)
		i++
	}
	return out

}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(gt(n))

}
