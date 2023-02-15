package main

import (
	"fmt"
)

func srr(n int) int {
	var ss = 0
	for i := 1; i < n+1; i++ {
		nb := i*i
		ss += nb
	}

	return ss

}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(srr(n))

}
