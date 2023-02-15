package main

import "fmt"

func sub(n int) []string {
	out := []string{}
	for i := 1; i < n+1; i++ {
		var ss string
		fmt.Scanln(&ss)
		if i < 3 || i > 3 {
			out = append(out, ss)
		}
	}
	return out

}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(sub(n))

}
