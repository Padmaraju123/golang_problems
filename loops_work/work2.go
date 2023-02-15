package main

import (
	"fmt"
	"strings"
)

func pat(m,n int) {
	i := 0
	for i < m {
		fmt.Println(strings.Repeat("* ", n))
		i++
	}
}

func main() {
	var m,n int
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	pat(m,n)

}
