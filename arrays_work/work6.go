package main

import (
	"fmt"
	"sort"
	"strconv"
)

func get_s(n int) []int {
	out := []int{}
	for i := 0; i < n; i++ {
		var j string
		fmt.Scanln(&j)
		gh, _ := strconv.Atoi(j)
		out = append(out, gh)
	}
	sort.Ints(out)
	return out
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(get_s(n))
}
