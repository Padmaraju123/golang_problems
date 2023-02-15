package main

import "fmt"

func slik(n, t int) []int {
	var ghh = []int{}
	for i := 0; i < n; i++ {
		var sd int
		fmt.Scanln(&sd)
		ghh = append(ghh, sd)
	}

	final := []int{}
	for j := 0; j < t; j++ {
		var ff int
		fmt.Scanln(&ff)
		final =append(final,ghh[ff])

	}
	return final

}

func main() {
	var n int
	fmt.Scanln(&n)
	var t int
	fmt.Scanln(&t)
	out := slik(n, t)
	for _, ss := range out {
		fmt.Println(ss)
	}
}
