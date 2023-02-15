package main

import "fmt"

func pro(n int) int {
	var p = 1
	for i := 0; i < n; i++ {
		var f int
		fmt.Scanln(&f)
		p *= f
	}

	return p

}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(pro(n))
}
